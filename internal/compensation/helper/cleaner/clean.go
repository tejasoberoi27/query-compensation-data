package cleaner

import (
	"encoding/csv"
	"log"
	"os"
	"query-compensation-data/internal/compensation/repository/compensation/postgres"
	"query-compensation-data/pkg/util"
	"strconv"
	"strings"
	"time"
)

var exchangeRates = map[string]float64{
	"USD": 1.0,
	"EUR": 1.08,
	"GBP": 1.25,
	"INR": 0.012,
	"CAD": 0.74,
	"AUD": 0.66,
}

// TODO: Enhancement (make this real-time through API call)
// Normalize and convert salary string to float64 in USD
func normalizeComp(s string) float64 {
	s = strings.ToUpper(strings.TrimSpace(s))

	// Currency symbols
	currencySymbols := map[string]string{
		"$":  "USD",
		"€":  "EUR",
		"£":  "GBP",
		"₹":  "INR",
		"C$": "CAD",
		"A$": "AUD",
	}

	// TODO: Enhance logic
	// Detect symbol and strip
	currency := "USD"
	for symbol, curr := range currencySymbols {
		if strings.HasPrefix(s, symbol) {
			currency = curr
			s = strings.TrimPrefix(s, symbol)
			break
		}
	}

	// Remove commas and "k" or "L" suffix
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, " ", "")

	multiplier := 1.0
	if strings.HasSuffix(s, "K") {
		multiplier = 1_000
		s = strings.TrimSuffix(s, "K")
	} else if strings.HasSuffix(s, "L") {
		multiplier = 100_000
		s = strings.TrimSuffix(s, "L")
	} else if strings.HasSuffix(s, "M") {
		multiplier = 1_000_000
		s = strings.TrimSuffix(s, "M")
	}

	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}

	usdRate := exchangeRates[currency]
	return num * multiplier * usdRate
}

// Normalize location into city and state
func normalizeLocation(loc string) (string, string) {
	loc = strings.TrimSpace(loc)
	if util.ContainsDigit(loc) || loc == "" {
		return "", ""
	}

	// Prefer comma-separated format (city, state)
	if strings.Contains(loc, ",") {
		parts := strings.SplitN(loc, ",", 2)
		return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
	}

	// Fallback: if space-separated, assume last token is state
	parts := strings.Fields(loc)
	if len(parts) > 1 {
		return strings.Join(parts[:len(parts)-1], " "), parts[len(parts)-1]
	}

	// Only one word
	return loc, ""
}

func CleanAndParseCSV(filePath string) ([]postgres.CompensationModel, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var entries []postgres.CompensationModel
	for i, row := range rows {
		if i == 0 {
			continue // skip header
		}

		if row[0] == "" {
			log.Printf("row %d : timestamp is empty, skipping", i)
			continue
		}

		timestamp, err := time.Parse("1/2/2006 15:04:05", strings.TrimSpace(row[0]))
		if err != nil {
			return nil, err
		}

		city, state := normalizeLocation(row[2])

		yearsEmployed, err := strconv.ParseFloat(cleanNumber(row[4]), 64)
		if err != nil {
			yearsEmployed = 0
		}
		yearsExp, err := strconv.ParseFloat(cleanNumber(row[5]), 64)
		if err != nil {
			yearsExp = 0
		}

		entries = append(entries, postgres.CompensationModel{
			Timestamp:          timestamp,
			Company:            strings.ToLower(strings.TrimSpace(row[1])),
			City:               city,
			State:              state,
			Title:              strings.ToLower(strings.TrimSpace(row[3])),
			YearsAtEmployer:    yearsEmployed,
			YearsExp:           yearsExp,
			AnnualBasePay:      normalizeComp(row[6]),
			SigningBonus:       normalizeComp(row[7]),
			AnnualBonus:        normalizeComp(row[8]),
			AnnualStockValue:   normalizeComp(row[9]),
			Gender:             strings.ToLower(strings.TrimSpace(row[10])),
			AdditionalComments: strings.TrimSpace(row[11]),
		})
	}

	return entries, nil
}

func cleanNumber(s string) string {
	s = strings.ReplaceAll(s, "$", "")
	s = strings.ReplaceAll(s, ",", "")
	return strings.TrimSpace(s)
}

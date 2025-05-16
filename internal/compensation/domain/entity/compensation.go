package entity

import (
	"query-compensation-data/internal/compensation/enum"
	"time"
)

type Compensation struct {
	ID                 int
	Timestamp          time.Time
	Company            string
	Title              string
	City               string
	State              string
	TotalComp          float64
	AnnualBasePay      float64
	AnnualBonus        float64
	AnnualStockValue   float64
	YearsExp           float64
	AdditionalComments string
	Gender             enum.Gender
	SigningBonus       float64
	YearsAtEmployer    float64
}

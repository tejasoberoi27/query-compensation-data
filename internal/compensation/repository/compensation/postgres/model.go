package postgres

import (
	"github.com/uptrace/bun"
	"query-compensation-data/internal/compensation/domain/entity"
	"query-compensation-data/internal/compensation/enum"
	"time"
)

type CompensationModel struct {
	bun.BaseModel `bun:"table:compensation"`

	ID                 int       `bun:",pk,autoincrement"`
	Timestamp          time.Time `bun:"timestamp,notnull"`
	Company            string    `bun:"company"`
	City               string    `bun:"city"`
	State              string    `bun:"state"`
	Title              string    `bun:"title"`
	YearsAtEmployer    float64   `bun:"years_at_employer"`
	YearsExp           float64   `bun:"years_exp"`
	AnnualBasePay      float64   `bun:"annual_base_pay"`
	SigningBonus       float64   `bun:"signing_bonus"`
	AnnualBonus        float64   `bun:"annual_bonus"`
	AnnualStockValue   float64   `bun:"annual_stock_value"`
	Gender             string    `bun:"gender"`
	AdditionalComments string    `bun:"additional_comments"`
}

func (m CompensationModel) Convert() *entity.Compensation {
	result := &entity.Compensation{
		ID:                 m.ID,
		Timestamp:          m.Timestamp,
		Company:            m.Company,
		Title:              m.Title,
		City:               m.City,
		State:              m.State,
		TotalComp:          m.AnnualBasePay + m.SigningBonus + m.AnnualBonus + m.AnnualStockValue,
		SigningBonus:       m.SigningBonus,
		AnnualBasePay:      m.AnnualBasePay,
		AnnualBonus:        m.AnnualBonus,
		AnnualStockValue:   m.AnnualStockValue,
		YearsExp:           m.YearsExp,
		AdditionalComments: m.AdditionalComments,
		Gender:             enum.NewGender(m.Gender),
		YearsAtEmployer:    m.YearsAtEmployer,
	}

	return result
}

func newModel(i *entity.Compensation) *CompensationModel {
	m := &CompensationModel{
		ID:                 i.ID,
		Timestamp:          i.Timestamp,
		Company:            i.Company,
		City:               i.City,
		State:              i.State,
		Title:              i.Title,
		YearsAtEmployer:    i.YearsAtEmployer, // Update if YearsAtEmployer is available in entity.Compensation
		YearsExp:           i.YearsExp,
		AnnualBasePay:      i.AnnualBasePay,
		SigningBonus:       i.SigningBonus, // Correctly map SigningBonus if available
		AnnualBonus:        i.AnnualBonus,
		AnnualStockValue:   i.AnnualStockValue,
		Gender:             i.Gender.String(),
		AdditionalComments: i.AdditionalComments,
	}

	return m
}

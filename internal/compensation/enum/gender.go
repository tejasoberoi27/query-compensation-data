package enum

import "strings"

type Gender string

const (
	GenderMale      Gender = "male"
	GenderFemale    Gender = "female"
	GenderUndefined Gender = "undefined"
)

func (g Gender) String() string {
	return string(g)
}

func NewGender(v string) Gender {
	switch strings.ToLower(v) {
	case GenderMale.String():
		return GenderMale
	case GenderFemale.String():
		return GenderFemale
	default:
		return GenderUndefined
	}
}

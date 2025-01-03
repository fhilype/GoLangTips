package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero"`
	CPF  string `json:"cpf" validate:"len=11, regexp=^[0-9]*$"`
	RG   string `json:"rg" validate:"len=7, regexp=^[0-9]*$"`
}

func Validate(aluno *Aluno) error {
	if err := validator.Validate(aluno); err == nil {
		return nil
	} else {
		return err
	}
}

package models

import (
	"github.com/uadmin/uadmin"
)

type Balmes struct {
	uadmin.Model
	Name       string `uadmin: "required"`
	Address    string `uadmin: "required"`
	Occupation string `uadmin: "required"`
}

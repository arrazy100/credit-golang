package main

import (
	"credit/models"
	"credit/models/base"
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
)

func main() {
	models := []interface{}{
		&base.User{},
		&models.SeedVersion{},
		&models.Sequence{},
		&models.Debtor{},
		&models.DebtorTenorLimit{},
		&models.DebtorTransaction{},
		&models.DebtorInstallment{},
		&models.DebtorInstallmentLine{},
	}

	stmts, err := gormschema.New("postgres").Load(models...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed generating schema: %v", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}

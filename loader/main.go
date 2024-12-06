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
		&models.Debtor{},
		&models.DebtorToUser{},
		&models.DebtorTenorLimit{},
		&models.DebtorTransaction{},
		&models.DebtorInstallment{},
	}

	stmts, err := gormschema.New("postgres").Load(models...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed generating schema: %v", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}

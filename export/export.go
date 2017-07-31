package export

import (
	"github.com/gocarina/gocsv"
	"github.com/jmoiron/sqlx"
	"github.com/kidsdynamic/swing-test-record/model"
)

func ExportIPQCToCSV(db *sqlx.DB) string {

	ipqc := []model.IPQCDatabase{}
	err := db.Select(&ipqc, "SELECT * FROM IPQC ORDER BY date_created DESC")

	if err != nil {
		panic(err)
	}

	csvContent, err := gocsv.MarshalString(&ipqc)

	if err != nil {
		panic(err)
	}

	return csvContent

}

func ExportFunctionToCSV(db *sqlx.DB) string {

	function := []model.FunctionDatabase{}
	err := db.Select(&function, "SELECT * FROM Function ORDER BY date_created DESC")

	if err != nil {
		panic(err)
	}

	csvContent, err := gocsv.MarshalString(&function)

	if err != nil {
		panic(err)
	}

	return csvContent

}

func ExportBarcodeToCSV(db *sqlx.DB) string {

	barcode := []model.BarcodeDatabase{}
	err := db.Select(&barcode, "SELECT * FROM Barcode ORDER BY date_created DESC")

	if err != nil {
		panic(err)
	}

	csvContent, err := gocsv.MarshalString(&barcode)

	if err != nil {
		panic(err)
	}

	return csvContent

}

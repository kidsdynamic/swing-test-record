package main

import (
	"database/sql"
	"log"

	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type IPQC struct {
	Type      int           `json:"Type"`
	LotNumber string        `json:"Lot_number"`
	Data      SwingFunction `json:"Data"`
}

type SwingFunction struct {
	SerialNumber string `json:"serial_number"`
	Voltage1     string `json:"Voltage_1"`
	Voltage2     string `json:"Voltage_2"`
	Result       string `json:"Result"`
	DateTime     string `json:"Date_time"`
}

func main() {
	InitDatabase()

	router := gin.Default()

	router.POST("/ipqc", IPOCHandler)
	router.Run(":8110")

}

func IPOCHandler(c *gin.Context) {
	db := NewDB()
	defer db.Close()

	var ipqc IPQC
	err := c.BindJSON(&ipqc)

	if ipqc.LotNumber == "" {
		log.Panicf("The Log number is required. Parameters: %v", ipqc)
		ErrorHandler(c, "The Log_number is required")
		return
	}

	_, err = db.Exec("INSERT INTO IPQC (type, lot_number, serial_number, voltage_1, voltage_2, result, date_time) VALUES (?, ?, ?, ?, ?, ?, ?)",
		ipqc.Type, ipqc.LotNumber, ipqc.Data.SerialNumber, ipqc.Data.Voltage1, ipqc.Data.Voltage2, ipqc.Data.Result, ipqc.Data.DateTime)

	if err != nil {
		log.Panic(err)
		ErrorHandler(c, fmt.Sprintf("Error on inserting data to database, please check your parameters."))
		return
	}

}

func ErrorHandler(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"message": message,
	})
}

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:koe7POut@tcp(45.55.248.58:3306)/swing_test_record?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func InitDatabase() {

	db := NewDB()
	defer db.Close()
	_, err := db.Query("SELECT 1 FROM IPQC LIMIT 1")

	log.Printf("Inital Database\n")
	if err != nil {
		_, err := db.Exec("CREATE TABLE IPQC(id INT NOT NULL AUTO_INCREMENT, type INT(11) NOT NULL, lot_number VARCHAR(200), serial_number VARCHAR(200) NOT NULL," +
			"voltage_1 VARCHAR(200), voltage_2 VARCHAR(200), result VARCHAR(200), date_time VARCHAR(200), PRIMARY KEY (id))")

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Table IPQC successfully created\n")
	} else {
		log.Printf("Table IPQC already exists\n")
	}

}

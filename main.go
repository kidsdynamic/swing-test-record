package main

import (
	"log"

	"net/http"

	"fmt"

	"os"

	"github.com/swing-test-record/model"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/urfave/cli"
)

type Database struct {
	Name     string
	User     string
	Password string
	IP       string
}

var authKey = "TestKey"

var database Database

func NewDB() *sqlx.DB {
	connectString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", database.User, database.Password, database.IP, database.Name)

	//dd, errr := sqlx.Connect("mysql", connectString)
	db, err := sqlx.Connect("mysql", connectString)
	if err != nil {
		log.Println(err)
	}

	return db
}

func main() {
	app := cli.NewApp()
	app.Name = "Swing-Push-Worker"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			EnvVar: "DATABASE_USER",
			Name:   "database_user",
			Usage:  "Database user name",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "DATABASE_PASSWORD",
			Name:   "database_password",
			Usage:  "Database password",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "DATABASE_IP",
			Name:   "database_IP",
			Usage:  "Database IP address with port number",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "DATABASE_NAME",
			Name:   "database_name",
			Usage:  "Database name",
			Value:  "swing_test_record",
		},
		cli.StringFlag{
			EnvVar: "API_AUTH_KEY",
			Name:   "auth_key",
			Usage:  "API auth key",
			Value:  "TestKey",
		},
	}

	app.Action = func(c *cli.Context) error {
		database = Database{
			Name:     c.String("database_name"),
			User:     c.String("database_user"),
			Password: c.String("database_password"),
			IP:       c.String("database_IP"),
		}
		authKey = c.String("auth_key")

		fmt.Printf("Database: %v", database)
		InitDatabase()

		router := gin.Default()
		router.Use(gin.Logger())
		router.Use(gin.Recovery())

		router.LoadHTMLGlob("view/html/**")

		api := router.Group("/api", Auth())
		//api.Use(Auth())
		api.POST("/ipqc", IPQCHandler)
		api.POST("/function", FunctionHandler)
		api.POST("/barcode", BarcodeHandler)

		router.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})

		router.GET("/ipqc", IPQCPage)
		router.GET("/barcode", BarcodePage)
		router.GET("/function", FunctionPage)

		//router.Run(":8110")
		router.RunTLS(":8110", "./.ssh/childrenlab.chained.crt", "./.ssh/childrenlab.com.key")
		return nil
	}

	app.Run(os.Args)

}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientAuthKey := c.Request.Header.Get("X-AUTH-TOKEN")

		if clientAuthKey == authKey {
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Please use correct X-AUTH-TOKEN",
			})
			c.Abort()
			return
		}
	}

}

func IPQCPage(c *gin.Context) {
	db := NewDB()
	defer db.Close()

	ipqc := []model.IPQCDatabase{}
	err := db.Select(&ipqc, "SELECT * FROM IPQC ORDER BY date_created DESC")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ipqc)

	c.HTML(http.StatusOK, "ipqc.html", gin.H{
		"data": ipqc,
	})
}

func FunctionPage(c *gin.Context) {
	db := NewDB()
	defer db.Close()

	functionData := []model.FunctionDatabase{}
	err := db.Select(&functionData, "SELECT * FROM Function ORDER BY date_created DESC")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(functionData)

	c.HTML(http.StatusOK, "function.html", gin.H{
		"data": functionData,
	})
}

func BarcodePage(c *gin.Context) {
	db := NewDB()
	defer db.Close()

	barcodeData := []model.BarcodeDatabase{}
	err := db.Select(&barcodeData, "SELECT * FROM Barcode ORDER BY date_created DESC")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(barcodeData)
	c.HTML(http.StatusOK, "barcode.html", gin.H{
		"data": barcodeData,
	})
}

func IPQCHandler(c *gin.Context) {
	db := NewDB()
	defer db.Close()

	var ipqc model.IPQC
	err := c.BindJSON(&ipqc)

	if err != nil {
		log.Println(err)
		ErrorHandler(c, fmt.Sprintf("Error on converting parameters to struct. %v", err))
		return
	}

	if ipqc.LotNumber == "" {
		log.Printf("The Log number is required. Parameters: %v", ipqc)
		ErrorHandler(c, "The Log_number is required")
		return
	}

	t := db.MustBegin()
	for _, data := range ipqc.Data {
		_ = t.MustExec("INSERT INTO IPQC (type, lot_number, serial_number, voltage_1, voltage_2, result, date_time, date_created) VALUES (?, ?, ?, ?, ?, ?, ?, NOW())",
			ipqc.Type, ipqc.LotNumber, data.SerialNumber, data.Voltage1, data.Voltage2, data.Result, data.DateTime)
	}

	t.Commit()

	if err != nil {
		log.Println(err)
		ErrorHandler(c, fmt.Sprintf("Error on inserting data to database, please check your parameters."))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "true",
	})

}

func FunctionHandler(c *gin.Context) {
	db := NewDB()
	defer db.Close()

	var function model.Function

	err := c.BindJSON(&function)

	if err != nil {
		log.Println(err)
		ErrorHandler(c, fmt.Sprintf("Error on converting parameters to struct. %v", err))
		return
	}

	if function.LotNumber == "" {
		log.Printf("The Log number is required for function API. Parameters: %v", function)
		ErrorHandler(c, "The Log_number is required")
	}

	t := db.MustBegin()
	for _, data := range function.Data {
		_, err = t.Exec("INSERT INTO Function (type, lot_number, serial_number, Date_time, BLE_result, UV_max, UV_min, UV_result,"+
			" Acc_x_max, Acc_x_min, Acc_x_result, Acc_y_max, Acc_y_min, Acc_y_result, Audio_max, Audio_result, Mac_address, RSSI, date_created) VALUES ("+
			"?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW())",
			function.Type, function.LotNumber, data.SerialNumber, data.DateTime, data.BLEResult, data.UVMax, data.UVMin, data.UVResult,
			data.AccXMax, data.AccXMin, data.AccXResult, data.AccYMax, data.AccYMin, data.AccYResult,
			data.AudioMax, data.AudioResult, data.MacAddress, data.Rssi)
		if err != nil {
			log.Println(err)
			ErrorHandler(c, fmt.Sprintf("Error on inserting data to database, please check your parameters."))
			return
		}
	}

	err = t.Commit()

	if err != nil {
		log.Println(err)
		ErrorHandler(c, fmt.Sprintf("Error on inserting data to database, please check your parameters."))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "true",
	})
}

func BarcodeHandler(c *gin.Context) {
	db := NewDB()
	defer db.Close()

	var barcode model.Barcode

	err := c.BindJSON(&barcode)

	if err != nil {
		log.Println(err)
		ErrorHandler(c, fmt.Sprintf("Error on converting parameters to struct. %v", err))
		return
	}

	t := db.MustBegin()
	for _, data := range barcode.Data {
		_ = t.MustExec("INSERT INTO Barcode (type, lot_number, barcode_number, date_time, date_created) VALUES (?, ?, ?, ?, NOW())",
			barcode.Type, barcode.LotNumber, data.BarcodeNumber, data.DateTime)
	}

	t.Commit()

	if err != nil {
		log.Println(err)
		ErrorHandler(c, fmt.Sprintf("Error on inserting data to database, please check your parameters."))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "true",
	})
}

func ErrorHandler(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": "false",
		"message": message,
	})
}

func InitDatabase() {

	db := NewDB()
	defer db.Close()
	_, err := db.Query("SELECT 1 FROM IPQC LIMIT 1")

	log.Printf("Inital Database\n")

	/*
	  Create IPQC Table
	*/
	if err != nil {
		_, err := db.Exec("CREATE TABLE IPQC(id INT NOT NULL AUTO_INCREMENT, type INT(11) NOT NULL, lot_number VARCHAR(200), serial_number VARCHAR(200) NOT NULL," +
			"voltage_1 VARCHAR(200), voltage_2 VARCHAR(200), result VARCHAR(200), date_time VARCHAR(200), date_created datetime NOT NULL, PRIMARY KEY (id))")

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Table IPQC successfully created\n")
	} else {
		log.Printf("Table IPQC already exists\n")
	}

	_, err = db.Query("SELECT 1 FROM Function LIMIT 1")

	/*
	  Create Function Table
	*/
	if err != nil {
		_, err = db.Exec("CREATE TABLE Function(id INT NOT NULL AUTO_INCREMENT, type INT(11) NOT NULL, lot_number VARCHAR(200), serial_number VARCHAR(200) NOT NULL," +
			"date_time VARCHAR(200), BLE_result VARCHAR(200), UV_max VARCHAR(200), UV_min VARCHAR(200), UV_result VARCHAR(200), Acc_x_max VARCHAR(200), Acc_x_min VARCHAR(200)," +
			"Acc_x_result VARCHAR(200), Acc_y_max VARCHAR(200), Acc_y_min VARCHAR(200), Acc_y_result VARCHAR(200), Audio_max VARCHAR(200)," +
			"Audio_result VARCHAR(200), Mac_address VARCHAR(200), RSSI VARCHAR(200), date_created datetime NOT NULL, PRIMARY KEY (id))")

		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Table Function successfully created\n")

	} else {
		log.Printf("Table Function already exists\n")
	}

	_, err = db.Query("SELECT 1 FROM Barcode LIMIT 1")

	/*
	  Create Barcode Table
	*/
	if err != nil {
		_, err = db.Exec("CREATE TABLE Barcode(id INT NOT NULL AUTO_INCREMENT, type INT(11) NOT NULL, lot_number VARCHAR(200), barcode_number VARCHAR(200) NOT NULL" +
			", date_created datetime NOT NULL, date_time VARCHAR(200), PRIMARY KEY (id))")

		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Table Barcode successfully created\n")

	} else {
		log.Printf("Table Barcode already exists\n")
	}

}

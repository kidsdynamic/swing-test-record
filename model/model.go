package model

import "time"

type IPQCDatabase struct {
	Id           int
	Type         int
	LotNumber    string    `db:"lot_number"`
	SerialNumber string    `db:"serial_number"`
	Voltage1     string    `db:"voltage_1"`
	Voltage2     string    `db:"voltage_2"`
	Result       string    `db:"result"`
	DateTime     string    `db:"date_time"`
	DateCreated  time.Time `db:"date_created"`
}

type FunctionDatabase struct {
	Id           int
	Type         int
	LotNumber    string    `db:"lot_number"`
	SerialNumber string    `db:"serial_number"`
	DateTime     string    `db:"date_time"`
	BLEResult    string    `db:"BLE_result"`
	UV           string    `db:"UV"`
	UVResult     string    `db:"UV_result"`
	AccXMax      string    `db:"Acc_x_max"`
	AccXMin      string    `db:"Acc_x_min"`
	AccXResult   string    `db:"Acc_x_result"`
	AccYMax      string    `db:"Acc_y_max"`
	AccYMin      string    `db:"Acc_y_min"`
	AccYResult   string    `db:"Acc_y_result"`
	AudioMax     string    `db:"Audio_max"`
	AudioResult  string    `db:"Audio_result"`
	MacAddress   string    `db:"Mac_address"`
	Rssi         string    `db:"RSSI"`
	DateCreated  time.Time `db:"date_created"`
}

type BarcodeDatabase struct {
	Id            int
	Type          int       `db:"type"`
	LotNumber     string    `db:"lot_number"`
	BarcodeNumber string    `db:"barcode_number"`
	DateCreated   time.Time `db:"date_created"`
}

type IPQC struct {
	Type      int      `json:"Type"`
	LotNumber string   `json:"Lot_number"`
	Data      IPQCData `json:"Data"`
}

type IPQCData struct {
	SerialNumber string `json:"Serial_number"`
	Voltage1     string `json:"Voltage_1"`
	Voltage2     string `json:"Voltage_2"`
	Result       string `json:"Result"`
	DateTime     string `json:"Date_time"`
}

type Function struct {
	Type      int          `json:"Type"`
	LotNumber string       `json:"Lot_number"`
	Data      FunctionData `json:"Data"`
}

type FunctionData struct {
	SerialNumber string `json:"Serial_number"`
	DateTime     string `json:"Date_time"`
	BLEResult    string `json:"BLE_result"`
	UV           string `json:"UV"`
	UVResult     string `json:"UV_result"`
	AccXMax      string `json:"Acc_x_max"`
	AccXMin      string `json:"Acc_x_min"`
	AccXResult   string `json:"Acc_x_result"`
	AccYMax      string `json:"Acc_y_max"`
	AccYMin      string `json:"Acc_y_min"`
	AccYResult   string `json:"Acc_y_result"`
	AudioMax     string `json:"Audio_max"`
	AudioResult  string `json:"Audio_result"`
	MacAddress   string `json:"Mac_address"`
	Rssi         string `json:"RSSI"`
}

type Barcode struct {
	Type      int         `json:"Type"`
	LotNumber string      `json:"Lot_number"`
	Data      BarcodeData `json:"Data"`
}

type BarcodeData struct {
	BarcodeNumber string `json:"Barcode_number"`
}
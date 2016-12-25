## [View Database Result](https://childrenlab.com:8110/)
# Swing 產測 API
* IPQC, Function, Barcode API

## Request Header Token
* The request header have to contains ```X-AUTH-TOKEN```
* Examples:
    * ```curl -X POST -H "Content-Type: application/json" -H "X-AUTH-TOKEN: TestKey" ```

## - POST /api/ipqc
1. ContentType: ```application/json```
2. Method: ```POST```
3. Parameters
    * Type - Number
    * Lot_number - String
    * Data - JSON
        * ```Serial_number``` - String
        * ```Voltage_1``` - String
        * ```Voltage_2``` - String
        * ```Result``` - String
        * ```Date_time``` - String
4. Example
    ```
    curl -X POST -H "Content-Type: application/json" -H "X-AUTH-TOKEN: TestKey" -d '{
         	"Type": 1,
         	"Lot_number": "19850830",
         	"Data": {
         		"Serial_number": "333333",
         		"Voltage_1": "100",
         		"Voltage_2": "200",
         		"Result": "OK",
         		"Date_time": "2016-10-11"
         	}

    }' "http://localhost:8110/api/ipqc"
    ```



## - POST /api/function
1. ContentType: ```application/json```
2. Method: ```POST```
3. Parameters
    * Type - Number
    * Lot_number - String
    * Data - JSON
        * ```Serial_number``` - String
        * ```Date_time``` - String
        * ```BLE_result``` - String
        * ```UV``` - String
        * ```UV_result``` - String
        * ```Acc_x_max``` - String
        * ```Acc_x_min``` - String
        * ```Acc_x_result``` - String
        * ```Acc_y_max``` - String
        * ```Acc_y_min``` - String
        * ```Acc_y_result``` - String
        * ```Audio_max``` - String
        * ```Audio_result``` - String
        * ```Mac_address``` - String
        * ```RSSI``` - String

4. Example
    ```
    curl -X POST -H "Content-Type: application/json" -H "X-AUTH-TOKEN: TestKey" -d '{
       	"Type": 2,
       	"Lot_number": "19850830",
       	"Data": {
       		"Serial_number": "1111111",
       		"Date_time": "2016-10-11",
       		"BLE_result": "200",
       		"UV": "23",
       		"UV_result": "3",
       		"Acc_x_max": "123",
       		"Acc_x_min": "123",
       		"Acc_x_result": "33",
       		"Acc_y_max": "33",
       		"Acc_y_min": "44",
       		"Acc_y_result": "88",
       		"Audio_max": "1",
       		"Audio_result": "33",
       		"Mac_address": "ghoahgowahfw",
       		"RSSI": "33"

       	}

       }' "http://localhost:8110/api/function"
    ```



## - POST /api/barcode
1. ContentType: ```application/json```
2. Method: ```POST```
3. Parameters
    * Type - Number
    * Lot_number - String
    * Data - JSON
        * ```Barcode_number``` - String

4. Example
    ```
    curl -X POST -H "Content-Type: application/json" -H "X-AUTH-TOKEN: TestKey" -d '{
       	"Type": 2,
       	"Lot_number": "19850830",
       	"Data": {
       		"Barcode_number": "1232412412"

       	}

       }' "http://localhost:8110/api/barcode"
    ```



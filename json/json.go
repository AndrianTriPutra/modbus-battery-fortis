package json

import (
	"encoding/json"
	"log"
)

type Cell struct {
	Cell1  string
	Cell2  string
	Cell3  string
	Cell4  string
	Cell5  string
	Cell6  string
	Cell7  string
	Cell8  string
	Cell9  string
	Cell10 string
	Cell11 string
	Cell12 string
	Cell13 string
	Cell14 string
	Cell15 string
	//Cell16 string
}

type Temperature struct {
	Temp1      string
	Temp2      string
	Temp3      string
	Temp4      string
	Temp_mossf string
	Temp_env   string
}

type Data_Pack struct {
	ID          uint8
	Current     string
	Voltage     string
	Soc         string
	Soh         string
	Rem_cap     string
	Full_cap    string
	Des_cap     string
	Cycle_count string
	Balance     string
	Cell        Cell
	Temperature Temperature
}

var Battery_Data struct {
	Dev_ID    string
	Qty       uint8
	Timestamp string
	Pack      Data_Pack
}

func Insert_Json(timestamp string, pack_id uint8, data [31]string) {

	Battery_Data.Timestamp = timestamp
	Battery_Data.Pack.ID = pack_id

	Battery_Data.Pack.Current = data[0]
	Battery_Data.Pack.Voltage = data[1]
	Battery_Data.Pack.Soc = data[2]
	Battery_Data.Pack.Soh = data[3]
	Battery_Data.Pack.Rem_cap = data[4]
	Battery_Data.Pack.Full_cap = data[5]
	Battery_Data.Pack.Des_cap = data[6]
	Battery_Data.Pack.Cycle_count = data[7]
	Battery_Data.Pack.Cell.Cell1 = data[8]
	Battery_Data.Pack.Cell.Cell2 = data[9]
	Battery_Data.Pack.Cell.Cell3 = data[10]
	Battery_Data.Pack.Cell.Cell4 = data[11]
	Battery_Data.Pack.Cell.Cell5 = data[12]
	Battery_Data.Pack.Cell.Cell6 = data[13]
	Battery_Data.Pack.Cell.Cell7 = data[14]
	Battery_Data.Pack.Cell.Cell8 = data[15]
	Battery_Data.Pack.Cell.Cell9 = data[16]
	Battery_Data.Pack.Cell.Cell10 = data[17]
	Battery_Data.Pack.Cell.Cell11 = data[18]
	Battery_Data.Pack.Cell.Cell12 = data[19]
	Battery_Data.Pack.Cell.Cell13 = data[20]
	Battery_Data.Pack.Cell.Cell14 = data[21]
	Battery_Data.Pack.Cell.Cell15 = data[22]
	//Battery_Data.Pack.Cell.Cell16 = data[23]
	Battery_Data.Pack.Temperature.Temp1 = data[24]
	Battery_Data.Pack.Temperature.Temp2 = data[25]
	Battery_Data.Pack.Temperature.Temp3 = data[26]
	Battery_Data.Pack.Temperature.Temp4 = data[27]
	Battery_Data.Pack.Temperature.Temp_mossf = data[28]
	Battery_Data.Pack.Temperature.Temp_env = data[29]
	Battery_Data.Pack.Balance = data[30]

	write, err := json.MarshalIndent(Battery_Data, " ", " ")
	if err == nil {
		data_battery := string(write)
		log.Printf("data_battery :\n%s", data_battery)
	} else {
		log.Printf("[func Insert_Json] error create json data_battery : %s" + string(err.Error()))
	}
}

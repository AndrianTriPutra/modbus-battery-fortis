package main

import (
	"fmt"
	"log"
	"time"

	. "fortis/battery"
	. "fortis/config"
	. "fortis/json"
)

func init() {
	Read_Config("./config/", "json", "config")
}

func main() {
	log.Println("============ Query Battery ===============")
	for j := Configuration.Pack_Start; j <= Configuration.Pack_End; j++ {

		log.Printf("============    slave_%v    ===============", j)
		timestamp := get_time("Asia/Jakarta")
		data_battery, warning, protection, fault, err := Battery(byte(j))
		if err != nil {
			log.Printf("============ error battery %v ===============", j)
		} else {
			Battery_Data.Dev_ID = Configuration.ID
			Battery_Data.Qty = uint8((Configuration.Pack_End - Configuration.Pack_Start) + 1)

			Insert_Json(timestamp, uint8(j), data_battery)
			fmt.Println()
			alarm(warning, protection, fault)
		}
	}
}

func alarm(warning, protection, fault [16]string) {
	log.Printf("warning     :%s", warning)
	log.Printf("protection  :%s", protection)
	log.Printf("fault       :%s", fault)
}

func get_time(timezone string) (timestamp string) {
	loc, _ := time.LoadLocation(timezone)
	timeutcplus := time.Now().In(loc)

	timestamp = timeutcplus.Format(time.RFC3339)
	return timestamp
}

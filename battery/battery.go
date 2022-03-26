package battery

import (
	"log"

	"github.com/goburrow/modbus"

	. "fortis/config"
)

func Battery(pack_id byte) (str_result [31]string, warning, protection, fault [16]string, err error) {

	// ============== init handler ==============
	handler := modbus.NewRTUClientHandler(Configuration.Port)
	handler.BaudRate = Configuration.Baudrate
	handler.DataBits = Configuration.Databits
	handler.Parity = Configuration.Parity
	handler.StopBits = Configuration.Stopbits
	handler.SlaveId = pack_id
	handler.Timeout = Configuration.Timeout
	// ============== init handler ==============

	err = handler.Connect()
	if err != nil {
		log.Printf("pack_id[%v] error cause port battery busy", pack_id)
	} else {
		client := modbus.NewClient(handler)

		data, err := client.ReadHoldingRegisters(Configuration.Register_Start, Configuration.Register_Length)
		data, err = client.ReadHoldingRegisters(Configuration.Register_Start, Configuration.Register_Length)
		if err != nil {
			log.Printf("pack_id[%v] error cause battery timeout", pack_id)
		} else {
			str_result, warning, protection, fault = parsing_pack(data)
		}
	}
	handler.Close()
	return str_result, warning, protection, fault, err
}

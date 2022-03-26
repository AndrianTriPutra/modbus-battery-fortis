package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

var Configuration struct {
	ID              string
	Pack_Start      uint16
	Pack_End        uint16
	Port            string
	Baudrate        int
	Databits        int
	Parity          string
	Stopbits        int
	Timeout         time.Duration
	Register_Start  uint16
	Register_Length uint16
}

func Read_Config(path, tipe, name string) (err error) {

	viper.AddConfigPath(path)
	viper.SetConfigType(tipe)
	viper.SetConfigName(name)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Println("error load config")
	}

	err = viper.Unmarshal(&Configuration)

	Configuration.Timeout = Configuration.Timeout * time.Millisecond

	log.Println(" ====================== read_config ====================== ")
	log.Printf("ID             :%v", Configuration.ID)
	log.Printf("Pack_Start     :%v", Configuration.Pack_Start)
	log.Printf("Pack_End       :%v", Configuration.Pack_End)
	log.Printf("Port           :%s", Configuration.Port)
	log.Printf("Baudrate       :%v", Configuration.Baudrate)
	log.Printf("Databits       :%v", Configuration.Databits)
	log.Printf("Parity         :%s", Configuration.Parity)
	log.Printf("Stopbits       :%v", Configuration.Stopbits)
	log.Printf("Timeout        :%v", Configuration.Timeout)
	log.Printf("Register_Start :%v", Configuration.Register_Start)
	log.Printf("Register_Length:%v", Configuration.Register_Length)
	fmt.Println()

	return err
}

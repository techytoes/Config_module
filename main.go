package main

import (
	"github.com/spf13/viper"
)

type Setup struct {
	Port int	`json:"port"`
	Host string	`json:"host"`
	GormLogging bool	`json:"gorm_logging"`

	//db struct {
	//	dbUsername  string	`json:"db_username"`
	//	dbPassword string	`json:"db_password"`
	//	dbAddr     string	`json:"db_addr"`
	//	dbPort     int	`json:"db_port"`
	//	dbName     string	`json:"db_name"`
	//}

	DBConnectionString string `json:"db_connection_string"`
}

type Configuration struct {
	Prod Setup	`json:"prod"`
	Debug Setup	`json:"debug"`
}

func ReadConfig(filename string, defaults map[string]interface{}) (*viper.Viper, error) {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	return v, err
}

func main() {
	
}
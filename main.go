package main

import (
	"github.com/spf13/viper"
)

type setup struct {
	port int	`json:"port"`
	host string	`json:"host"`
	gormLogging bool	`json:"gorm_logging"`

	db struct {
		dbUsername  string	`json:"db_username"`
		dbPassword string	`json:"db_password"`
		dbAddr     string	`json:"db_addr"`
		dbPort     int	`json:"db_port"`
		dbName     string	`json:"db_name"`
	}
}

type configuration struct {
	prod setup	`json:"prod"`
	debug setup	`json:"debug"`
}

func readConfig(filename string, defaults map[string]interface{}) (*viper.Viper, error) {
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
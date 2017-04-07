package config
import (
)

type config struct {
	StartUrl string
	ResquestMethod string
	Reqlen int
	Reslen int
	Urllen int
	Itemlen int
}
var Config *config= new(config)
func InitConfig(){
	Config.ResquestMethod = "GET"
	Config.Reqlen = 10000000
	Config.Reslen = 10000000
	Config.Urllen = 1000000
	Config.Itemlen = 1000000
}
package config

import (
	"github.com/spf13/viper"
	"net"
)

func GetListener() net.Listener {

	grpcListener, err := net.Listen("tcp", ":"+viper.GetString("app_port"))
	if err != nil {
		panic(err.Error())
	}

	return grpcListener
}

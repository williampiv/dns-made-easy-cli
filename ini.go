package main

import (
	"errors"

	"gopkg.in/ini.v1"
)

type apiCredentials struct {
	APIKey    string
	SecretKey string
}

func getINICredentials(path string) (apiCredentials, error) {
	var iniData apiCredentials
	cfg, err := ini.Load(path)
	if err != nil {
		return iniData, errors.New("failed to load ini file")
	}
	iniData.APIKey = cfg.Section("").Key("apikey").String()
	iniData.SecretKey = cfg.Section("").Key("secretkey").String()
	return iniData, nil
}

package goessentials

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

var (
	ErrorReadConfigFailed      = errors.New("Unable to read config")
	ErrorUnmarshalConfigFailed = errors.New("Unable to unmarshal config")
	ErrorMarshalConfigFailed   = errors.New("Unable to marshal config")
	ErrorWriteConfigFailed     = errors.New("Unable to write config")
	ErrorConfigFileUnset       = errors.New("Config file wasn't set prior to call")
)

//GetConfig fetches the config located at 'configFileName'
//and puts its parsed contents into 'configPointer' if the
//config doesn't exist, it's created with the default from
//'configPointer' and returns the error 'ErrorConfigFileUnset'
func GetConfig(configFileName string, configPointer interface{}) error {
	if FileExists(configFileName) { //Get existing configuration from configFileName
		b, err := ioutil.ReadFile(configFileName)
		if err != nil {
			return ErrorReadConfigFailed
		}

		err = json.Unmarshal(b, configPointer)
		if err != nil {
			return ErrorUnmarshalConfigFailed
		}

		return nil
	}

	//If configFileName doesn't exist, create a new config file
	b, err := json.MarshalIndent(configPointer, "", " ")
	if err != nil {
		return ErrorMarshalConfigFailed
	}

	err = ioutil.WriteFile(configFileName, b, 0644)
	if err != nil {
		return ErrorWriteConfigFailed
	}

	return ErrorConfigFileUnset
}

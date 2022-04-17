package goessentials

import (
	"log"
)

type ExampleConfig struct {
	Works bool
}

//This both serves as the default configuration
//and a global config variable. If you don't
//want your config to be globally accessable,
//you could also define this somewhere else.
var config = ExampleConfig{
	Works: true,
}

//ExampleGetConfig provides a simple example
func ExampleGetConfig() {
	err := GetConfig("exampleConfig.json", &config)
	if err != nil && err != ErrorConfigFileUnset {
		log.Fatalln(err)
	}
}

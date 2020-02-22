package config

import (
	"io/ioutil"
	"log"
	"strconv"

	"gopkg.in/yaml.v2"
)

//GetProperty returns the value a given parameter
func GetProperty(Configuration interface{}, Section string, Key string) (Value string, err error) {
	data := Configuration.(map[interface{}]interface{})
	if section, ok := data[Section].(map[interface{}]interface{}); ok {
		if value, ok := section[Key]; ok {
			switch value.(type) {
			case int:
				return strconv.Itoa(value.(int)), nil
			case string:
				return value.(string), nil
			default:
				return
			}
		}
		log.Println("Key: " + Key + " is not present")
		return
	}
	log.Println("Section: ", Section, " is not present in config")
	return
}

//ReadConfigFile reads the config
func ReadConfigFile(File string) (interface{}, error) {
	var v interface{}

	data, err := ioutil.ReadFile(File)

	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(data, &v)
	if err != nil {
		log.Fatalln(err)
	}

	return v, nil
}

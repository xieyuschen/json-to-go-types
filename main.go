package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main()  {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config interface{}
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Panic(err)
	}
	m:=config.(map[string]interface{})
	fmt.Println(m["Hello"])
}

package main

import (
	"encoding/json"
	"fmt"
	"go/generator"
	"go/models"
	"io/ioutil"
	"log"
	"os"
)
func ReadSettingsFromFile(settingFilePath string)(config models.Config){
	jsonFile, err := os.Open(settingFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Panic(err)
	}
	return config
}
func main(){
	generator.GenerateModels("config.json")
	config:=ReadSettingsFromFile("config.json")
	fmt.Println(config.Config1.C12)
}

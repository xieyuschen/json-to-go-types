package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)
type Hello struct{
	hh string
	world int
}
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
	//panic: interface conversion: interface {} is map[string]interface {}, not main.Hello
	//Cannot use assets to convert an interface to a struct
	ShowToYou(m)

}
func ShowToYou(m map[string]interface{}){
	for _,value:=range m{
		t:=GetType(value)
		if t=="interface"{
			ShowToYou(value.(map[string]interface{}))
		}else {
			fmt.Println(value," type is:",t)
		}
	}
}
func GetType(val interface{}) string{
	if _,ok:=val.(int);ok{
		return "int"
	}else if _,ok:=val.(float64);ok{
		return "float64"
	}else if _,ok:=val.(string);ok{
		return "string"
	}else if _,ok:=val.(interface{});ok{
		return "interface"
	}else {
		return "halt"
	}
}

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

	f,err:=os.Create("Config_gen.go")
	if err!=nil{
		log.Panic(err.Error())
	}
	m:=config.(map[string]interface{})
	//panic: interface conversion: interface {} is map[string]interface {}, not main.Hello
	//Cannot use assets to convert an interface to a struct
	f.WriteString("package main\ntype Config struct{")
	f.Write([]byte("\n"))
	ShowToYou(m,f)
	f.WriteString("}")

}
func ShowToYou(m map[string]interface{},file *os.File){
	for index,value:=range m{
		t:=GetType(value)
		if t=="interface"{
			ShowToYou(value.(map[string]interface{}),file)
		}else {
			file.WriteString("\t"+index+" "+t)
			file.Write([]byte("\n"))
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

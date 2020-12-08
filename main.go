package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang-collections/collections/queue"
	"io/ioutil"
	"log"
	"os"
	"strings"
)
var que *queue.Queue
func main()  {
	que=queue.New()
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
	ModelGenerator(m,f)
	f.WriteString("}")

	for que.Len()!=0{
		name:=que.Dequeue().(string)
		name=strings.Title(strings.ToLower(name))
		f.Write([]byte("\n"))
		f.WriteString("type "+name+" struct{")
		f.Write([]byte("\n"))
		ModelGenerator(que.Dequeue().(map[string]interface{}),f)
		f.WriteString("}")
	}

	fmt.Println("Generate configuration models by config.json file successfully,,models are stored in Config_gen.go")
}
func ModelGenerator(m map[string]interface{},file *os.File){
	for index,value:=range m{
		t:=GetType(value)
		if t=="interface"{
			que.Enqueue(index)
			que.Enqueue(value)
		}else {
			file.WriteString("\t" + strings.Title(strings.ToLower(index))+" "+t)
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

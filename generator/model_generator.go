package generator

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
var configQue *queue.Queue
var singleQue *queue.Queue
func GenerateModels(ConfigFilePath string,GenerateModelPath string)  {
	que =queue.New()
	configQue =queue.New()
	singleQue =queue.New()
	jsonFile, err := os.Open(ConfigFilePath)
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
	fmt.Println(GenerateModelPath)
	f,err:=os.Create(GenerateModelPath)
	if err!=nil{
		log.Panic(err.Error())
	}
	m:=config.(map[string]interface{})
	//panic: interface conversion: interface {} is map[string]interface {}, not main.Hello
	//Cannot use assets to convert an interface to a struct
	f.WriteString("package models\n")
	configTypeGenerator(m,f)

	for que.Len()!=0{
		name:= que.Dequeue().(string)
		name=strings.Title(strings.ToLower(name))
		f.Write([]byte("\n"))
		configQue.Enqueue(name)
		f.WriteString("type "+name+" struct{")
		f.Write([]byte("\n"))
		modelGenerator(que.Dequeue().(map[string]interface{}),f)
		f.WriteString("}")
	}
	f.Write([]byte("\n"))
	f.WriteString("type Config struct{")
	for singleQue.Len()!=0{
		f.Write([]byte("\n"))
		_name:= singleQue.Dequeue().(string)
		_type:= singleQue.Dequeue().(string)
		f.WriteString("\t"+_name+" "+_type)
	}
	for configQue.Len()!=0{
		f.Write([]byte("\n"))
		n:= configQue.Dequeue().(string)
		f.WriteString("\t"+n+" "+n)
	}
	f.Write([]byte("\n"))
	f.WriteString("}")

	fmt.Println("Generate configuration models by config.json file successfully,,models are stored in Config_gen.go")
}
func configTypeGenerator(m map[string]interface{},file *os.File){
	for index,value:=range m{
		t:= getType(value)
		if t=="interface"{
			que.Enqueue(index)
			que.Enqueue(value)
		}else {
			singleQue.Enqueue(index)
			singleQue.Enqueue(t)
		}
	}
}
func modelGenerator(m map[string]interface{},file *os.File){
	for index,value:=range m{
		t:= getType(value)
		if t=="interface"{
			que.Enqueue(index)
			que.Enqueue(value)
		}else {
			file.WriteString("\t" + strings.Title(strings.ToLower(index))+" "+t)
			file.Write([]byte("\n"))
		}
	}
}
func getType(val interface{}) string{
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

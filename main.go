package main

import (

	"fmt"
	"gojsonconfig/generator"

	flags "github.com/jessevdk/go-flags"

	"os"
)
//func ReadSettingsFromFile(settingFilePath string)(config models.Config){
//	jsonFile, err := os.Open(settingFilePath)
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer jsonFile.Close()
//	byteValue, _ := ioutil.ReadAll(jsonFile)
//	err = json.Unmarshal(byteValue, &config)
//	if err != nil {
//		log.Panic(err)
//	}
//	return config
//}
var opts struct {

	Generator	 bool `short:"g"`
	Version bool `short:"v" long:"version"`
	ConfigPath string `short:"p" long:"path" description:"Store config path"`
	StoragePath string `short:"s" long:"storage"`
}
func main(){
	_, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		panic(err)
	}
	if opts.Version{
		fmt.Println("v0.1")
	}else if opts.Generator{
		generator.GenerateModels("config.json","./models/Config_gen.go")
	}
}
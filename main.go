package generator

import (
	"fmt"
	flags "github.com/jessevdk/go-flags"
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
type Opts struct {

	Generator	 bool `short:"g"`
	Version bool `short:"v" long:"version"`
	ConfigPath string `short:"p" long:"path" description:"Store config path"`
	StoragePath string `short:"s" long:"storage"`
}
func main(){
	//generator.GenerateModels("config.json")
	//config:=ReadSettingsFromFile("config.json")
	//fmt.Println(config.Config1.C12)
	var opts Opts

	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = "gocloc"
	parser.Usage = "[OPTIONS] PATH[...]"


	if opts.Version{
		fmt.Println("v0.1")
	}else if opts.Generator{
		fmt.Println(opts.ConfigPath)
		var jsonPath string
		var storage string
		if opts.ConfigPath!=""{
			jsonPath= opts.ConfigPath
		}else {
			jsonPath="config.json"
		}
		if opts.StoragePath!=""{
			storage= opts.StoragePath
		}else {
			storage="./models/Config_gen.go"
		}
		fmt.Println(jsonPath,storage)
		GenerateModels(jsonPath,storage)
	}
}

# Reading Json Configuration by Golang
As we all know, go language gives us a simple library to read json file. However, when it comes to more complex json file it
is weak,so we want to give you a stronger and more robust way to get json value.

# How to use it
[中文版本](./README_ZH.md)

- Use command below to install package `gojsonconfig`
```
go mod -u "https://github.com/xieyuschen/json-to-go-types"
```
- You can use command `gojsonconfig -v` to check whether you install this package successfully or not. If you are successful,
you will see output as this at terminal.
```sh
$ json-to-go-types -v
v0.1
```
- Generate a struct based on json file for further reading
Firstly your current path needs to contain a json file named `config.json`. Secondly your current path needs to have a folder
called `models` where the generated struct model will be stored.
If your path has no folder named `models`,please create one first.Your path dictionary would like this:
```
$-->your current postion
|__config,json
|__models(This is a folder)

```
Do command below after all of thing before done.
 ```
json-to-go-types -g
```
After this you can find a `Config_gen.go` file in `models` dictionary which is a struct based on the `config.json`.
- You can read json to the model generated before like this:
```go
func main(){
	conf:=ReadSettingsFromFile("config.json")
	fmt.Println(config.DbSettings.Password)
}
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
```
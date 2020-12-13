# 如何使用该库
- 安装package
```
go mod -u "https://github.com/Pivot-Studio/gojsonconfig"
```
- 你可以使用gojsonconfig -v来检查是否成功安装，如果成功安装，则会显示如下内容：
```sh
$ gojsonconfig -v
v0.1
```
- 为我们的Json文件生成结构体
首先你需要在当前目录下有一个名为`config.json`的json文件，并且在当前目录下有一个名为`models`的文件夹，`gojsonconfig`包会将结构体生成到`models·
目录下，如果没有请手动创建一个.你的项目目录应该是这样的:
```
$-->cmd中你现在的位置
|__config,json
|__models(这是一个文件夹)
    |__（这里生成之前没有内容，即models是一个空文件夹）
```
在此完成之后，在当前目录中运行命令：
```
gojsonconfig -g
```
好了，在该步完成之后，你就可以在`models`文件夹中看到一个`Config_gen.go`文件，这个文件中存储着根据`config.json`文件结构生产的结构体。

- 使用结构体读取json文件
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
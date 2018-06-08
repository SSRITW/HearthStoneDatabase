package restgo

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"path"
)

type Datasource struct {
	ConnetionURL string `yaml:"connetionURL"`
	DriveName string `yaml:"driveName"`
}

func ( datasourse *Datasource) InitData(){
	filePath := path.Join("config","application.properties")
	data, _ := ioutil.ReadFile(filePath)
	//把yaml形式的字符串解析成struct类型
	yaml.Unmarshal(data, &datasourse)
}

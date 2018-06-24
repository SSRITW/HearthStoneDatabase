package restgo

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"path"
	"os"
)

type DBConnectionInfo struct {
	Database struct{
		DriveName string `yaml:"drive_name,omitempty"`
		ConnetionURL string `yaml:"connetion_url,omitempty"`
	} `yaml:"database"`
}


func ( datasourse *DBConnectionInfo) InitData(){
	wr, _ := os.Getwd()
	filePath := path.Join(wr, "config", "db_config.yml")
	data, _ := ioutil.ReadFile(filePath)
	//把yaml形式的字符串解析成struct类型
	yaml.Unmarshal(data, &datasourse)
}

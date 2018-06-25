package restgo

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"path"
	"os"
	"strings"
)

const(
	PROJECTNAME = "HearthStoneDatabase"
	VISITOR = "visitor"
)
var ConfigDataSource DBConnectionInfo

type DBConnectionInfo struct {
	Database struct{
		DriveName string `yaml:"drive_name,omitempty"`
		ConnetionURL string `yaml:"connetion_url,omitempty"`
		MaxIdleConns int `yaml:"max_idle_conns,omitempty"`
		MaxOpenConns int `yaml:"max_open_conns,omitempty"`
	} `yaml:"database"`
	Redis struct{
		Host string `yaml:"host,omitempty"`
		Password string `yaml:"password,omitempty"`
		AccessTokenDb int `yaml:"access_token_db,omitempty"`
		RefreshTokenDb int `yaml:"refresh_token_db,omitempty"`
		MaxIdle int `yaml:"max_idle,omitempty"`
		MaxActive int `yaml:"max_active,omitempty"`
	} `yaml:"redis"`
}


func ( datasourse *DBConnectionInfo) InitData(){
	wr, _ := os.Getwd()
	//去除项目名称后面的字符串（解决test方法无法读取配置文件的问题
	projectPath := wr[0:strings.Index(wr,PROJECTNAME)+len(PROJECTNAME)]

	filePath := path.Join(projectPath, "config", "db_config.yml")
	data, _ := ioutil.ReadFile(filePath)
	//把yaml形式的字符串解析成struct类型
	yaml.Unmarshal(data, &datasourse)
}

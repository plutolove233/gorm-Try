//author:plutolove
//date:2021/2/1 16:43
//focus:分层设计尝试

package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	Mysql MySQLConfig
}

type MySQLConfig struct{
	Host string
	Ports int
	Username string
	Password string
	Database string
	Charset string
}

var (
	config *Config
	once sync.Once
)

func (c *Config)InitConfig(){
	viper.SetConfigFile("configs/config.toml")
	if err := viper.ReadInConfig();err!=nil{
		fmt.Println("配置文件导入失败，错误原因：%v",err)
	}
	if err := viper.Unmarshal(c);err !=nil{
		fmt.Println("配置文件绑定失败，错误原因：%v",err)
	}
}

func GetConfigInf()(string,string){
	once.Do(func(){
		config = &Config{}
		config.InitConfig()
	})
	return "mysql",fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",config.Mysql.Username,
		config.Mysql.Password, config.Mysql.Host, config.Mysql.Ports, config.Mysql.Database)
}
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"gopkg.in/ini.v1"
)


type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

func LoadConf(filename string, data interface{}) (err error) {
	// 参数校验
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Name(), t.Kind())
	if (t.Kind() != reflect.Ptr) || (t.Elem().Kind() != reflect.Struct) {
		err = errors.New("type errors")
		return
	}

	// 读配置文件一行一行的去遍历
	lines_data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	lines := strings.Split(string(lines_data), "\r\n")
	fmt.Println(lines)
	for index, value := range lines {
		fmt.Println(index, value)
		// 如果是注释或空行直接跳过
		value = strings.TrimSpace(value)
		if len(value) == 0 || strings.HasPrefix(value, "#") {
			continue
		}
        
		//开始解析字符串
		if strings.HasPrefix(value, "[") {
			// 如果是section
			section := value[1:len(value)-1]
			fmt.Println(section)
		} else {
			// pass
		}
	}
	return
}

func test() {
	var mysql MysqlConfig
	err := LoadConf("./conf.ini", &mysql)
	if err != nil {
		fmt.Println("err: ", err)
	}
}

// 读取配置文件
func test2() (err error) {
    cfg, err := ini.Load("./conf.ini")
	if err != nil {
		return
	}
    
	// 读取某一个属性
	address := cfg.Section("mysql").Key("address").String()
	fmt.Println(address)

	// 映射成结构体
	mysql := new(MysqlConfig)
	cfg.Section("mysql").MapTo(mysql)
	fmt.Println(mysql, mysql.Address, mysql.Port, mysql.Username, mysql.Password)
	return
}

func main() {
	test2()
}
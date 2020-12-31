package conf

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// configs 配置文件反序列化后的map
var configs = make(map[string]interface{}, 0)

func init() {
	confByteArr, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(fmt.Sprintf("读取配置文件【conf.yml】异常：%v", err))
	}
	if err := yaml.Unmarshal(confByteArr, &configs); err != nil {
		panic(fmt.Sprintf("反序列化配置异常：%v", err))
	}
}

// New 配置项赋值
func New(key string, ptr interface{}) {
	in, exists := configs[key]
	if !exists {
		return
	}

	// 序列化配置
	byteArr, err := yaml.Marshal(in)
	if err != nil {
		panic(fmt.Sprintf("序列化配置【%s】异常：%v", key, err))
	}

	// 序列化为对应的类型并赋值
	if err := yaml.Unmarshal(byteArr, ptr); err != nil {
		panic(fmt.Sprintf("配置【%s】赋值异常：%v", key, err))
	}
}

package config

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"path/filepath"
)

const (
	ConfigDefaultFile = "config.yaml"
)

// Viper //
// 优先级: 命令行 > 环境变量 > 默认值
func Viper(path ...string) *viper.Viper {
	var configIn string

	if len(path) == 0 {
		flag.StringVar(&configIn, "c", "", "choose configIn file.")
		flag.Parse()
		if configIn == "" { // 判断命令行参数是否为空
			configIn = ConfigDefaultFile
			fmt.Printf("您正在使用config的路径为%s\n", ConfigDefaultFile)
		} else { // 命令行参数不为空 将值赋值于config
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", configIn)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		configIn = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", configIn)
	}

	v := viper.New()
	v.SetConfigFile(configIn)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error configIn file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&Config); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&Config); err != nil {
		fmt.Println(err)
	}

	// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	Config.AutoCode.Root, _ = filepath.Abs("..")
	return v
}

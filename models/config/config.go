/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         config.go
@ Create Time:  2020/4/28 11:13
@ Software:     GoLand
*/

package config

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type config struct {
	Base  base  `yaml:"Base"`
	Mysql mysql `yaml:"Mysql"`
}

func (c *config) Print() {
	buf, _ := yaml.Marshal(&c)
	fmt.Println(string(buf))
}

var Config *config

func init() {
	buf, err := ioutil.ReadFile(fmt.Sprintf("configs/%s.yaml", os.Getenv("ENV")))
	if err != nil {
		log.Fatalf("load config failed. error: %s", err.Error())
	}

	if err = yaml.Unmarshal(buf, &Config); err != nil {
		log.Fatalf("format config file failed. error: %s", err.Error())
	}

	curPath, _ := os.Getwd()
	Config.Base.Logs = filepath.Join(curPath, "logs")
	if err = os.MkdirAll(Config.Base.Logs, 0755); err != nil {
		log.Fatalf("directory creation failed. error: %s", err.Error())
	}
}

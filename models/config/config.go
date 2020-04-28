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
	Base base
}

func (c *config) Print() {
	if buf, err := yaml.Marshal(&c); err != nil {
		log.Println(err.Error())
	} else {
		fmt.Println(string(buf))
	}
}

var Config *config

func init() {
	file := fmt.Sprintf("configs/%s.yml", os.Getenv("ENV"))

	buf, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("load config is failed. error: %s", err.Error())
	}

	if err = yaml.Unmarshal(buf, &Config); err != nil {
		log.Fatalf("the config file <%s> is not yaml. error: %s", "", err.Error())
	}

	curPath, _ := os.Getwd()
	Config.Base.Logs = filepath.Join(curPath, "logs")
	if err = os.MkdirAll(Config.Base.Logs, 0755); err != nil {
		log.Fatalf("directory creation failed. error: %s", err.Error())
	}
}

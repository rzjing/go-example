/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         basic.go
@ Create Time:  2020/4/28 11:25
@ Software:     GoLand
*/

package config

type base struct {
	Name string `yaml:"name"`
	Logs string `yaml:"logs"`
}

type mysql struct {
	Uri string `yaml:"uri"`
}

type redis struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Password    string `yaml:"password"`
	DB          int    `yaml:"db"`
	MaxIdle     int    `yaml:"max_idle"`
	MaxActive   int    `yaml:"max_active"`
	IdleTimeout int    `yaml:"idle_timeout"`
}

type mongo struct {
	Uri string `yaml:"uri"`
}

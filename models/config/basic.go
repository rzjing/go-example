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

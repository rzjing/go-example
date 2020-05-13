/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         struct.go
@ Create Time:  2020/5/12 17:50
@ Software:     GoLand
*/

package example

type DoOptions struct {
	f func(do *doOptions)
}

type doOptions struct {
	key    string
	value  interface{}
	expire int
}

func NewDoOptions(options ...DoOptions) *doOptions {
	do := doOptions{}
	for _, option := range options {
		option.f(&do)
	}
	return &do
}

func DoKey(key string) DoOptions {
	return DoOptions{f: func(do *doOptions) {
		do.key = key
	}}
}

func DoValue(val interface{}) DoOptions {
	return DoOptions{f: func(do *doOptions) {
		do.value = val
	}}
}

func DoExpire(expire int) DoOptions {
	return DoOptions{f: func(do *doOptions) {
		do.expire = expire
	}}
}

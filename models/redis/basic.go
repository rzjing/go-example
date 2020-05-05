/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         basic.go
@ Create Time:  2020/5/5 12:39
@ Software:     GoLand
*/

package redis

type DoOption struct {
	f func(*doOptions)
}

type doOptions struct {
	key    string
	value  interface{}
	expire int
}

func DoKey(key string) DoOption {
	return DoOption{f: func(do *doOptions) {
		do.key = key
	}}
}

func DoValue(val interface{}) DoOption {
	return DoOption{func(do *doOptions) {
		do.value = val
	}}
}

func DoExpire(expire int) DoOption {
	return DoOption{func(do *doOptions) {
		do.expire = expire
	}}
}

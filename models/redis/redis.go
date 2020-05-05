/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         redis.go
@ Create Time:  2020/5/5 10:24
@ Software:     GoLand
*/

package redis

import (
	redisGo "github.com/gomodule/redigo/redis"
	. "go-example/models/config"
	"log"
	"strconv"
	"strings"
	"time"
)

func Do(command string, options ...DoOption) (err error) {
	conn := pool.Get()
	defer conn.Close()

	do := doOptions{}
	for _, option := range options {
		option.f(&do)
	}

	switch strings.ToUpper(command) {
	case "GET":
		_, err = conn.Do(command, do.key)
	case "SETEX":
		_, err = conn.Do(command, do.key, do.expire, do.value)
	case "DEL":
		_, err = conn.Do(command, do.key)
	}
	return
}

var pool *redisGo.Pool

func init() {
	pool = &redisGo.Pool{
		MaxIdle:     Config.Redis.MaxIdle,
		MaxActive:   Config.Redis.MaxActive,
		IdleTimeout: time.Duration(Config.Redis.IdleTimeout) * time.Second,
		Wait:        true,
		Dial: func() (conn redisGo.Conn, err error) {
			uri := Config.Redis.Host + ":" + strconv.Itoa(Config.Redis.Port)
			conn, err = redisGo.Dial("tcp", uri, redisGo.DialPassword(Config.Redis.Password), redisGo.DialDatabase(Config.Redis.DB))

			if err != nil {
				log.Fatalln(err.Error())
			}

			return conn, nil
		},
	}

	// 测试连接可用性
	conn := pool.Get()
	defer conn.Close()
}

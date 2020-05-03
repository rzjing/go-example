/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         tokenizer.go
@ Create Time:  2020/5/3 11:34
@ Software:     GoLand
*/

package tools

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomString(length int) string {
	buf := make([]byte, length)
	for i := range buf {
		buf[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(buf)
}

func MD5Hash(text string, refresh bool) string {
	hash := md5.New()
	if refresh {
		text += RandomString(4)
	}
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

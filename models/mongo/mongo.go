/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         mongo.go
@ Create Time:  2020/5/13 17:06
@ Software:     GoLand
*/

package mongo

import (
	"context"
	. "go-example/models/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetConn() (client *mongo.Client, err error) {
	client, err = mongo.NewClient(options.Client().ApplyURI(Config.Mongo.Uri))
	return
}

func init() {
	conn, err := GetConn()
	if err != nil {
		panic(err.Error())
	}

	err = conn.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err.Error())
	}
}

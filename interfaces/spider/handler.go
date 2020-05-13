/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         handler.go
@ Create Time:  2020/5/13 10:37
@ Software:     GoLand
*/

package spider

import (
	"encoding/json"
	"go-example/interfaces"
	"log"
)

type DouBan struct{}

// 获取豆瓣电影热门推荐
func (db *DouBan) GetRecommend() error {
	url := "https://movie.douban.com/j/search_subjects?type=movie&tag=%E7%83%AD%E9%97%A8&sort=recommend&page_limit=20&page_start=0"
	resp, err := interfaces.Client("get", url)
	if err != nil {
		return err
	}

	var result map[string][]interface{}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return err
	}

	if _, ok := result["subjects"]; ok {
		for _, item := range result["subjects"] {
			db.Save(&item)
		}
	}
	return nil
}

// 数据存储
func (db *DouBan) Save(data *interface{}) {
	log.Printf("data %+v is saved", data)
}

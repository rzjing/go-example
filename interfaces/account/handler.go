/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         handler.go
@ Create Time:  2020/4/30 17:28
@ Software:     GoLand
*/

package account

import (
	"errors"
	"go-example/models/entity"
	"go-example/models/mysql"
	"go-example/tools"
	"log"
	"net/http"
)

func getAccount(p *getParams) (interface{}, error) {
	if db, err := mysql.GetConn(); err != nil {
		log.Println(err.Error())
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	} else {
		var count int
		var accounts []struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Email     string `json:"email"`
			Status    int    `json:"status"`
			CreatedAt int    `json:"created_at"`
			UpdatedAt int    `json:"updated_at"`
		}

		db.Table("account").
			Select("id, name, email, status, unix_timestamp(created_at) created_at, unix_timestamp(updated_at) updated_at").
			Order(p.Sort).Offset(p.Offset).Limit(p.PageSize).
			Scan(&accounts).Count(&count)

		data := map[string]interface{}{"code": count, "list": accounts}
		return data, nil
	}
}

func newAccount(p *newParams) error {
	if db, err := mysql.GetConn(); err != nil {
		log.Println(err.Error())
		return errors.New(http.StatusText(http.StatusInternalServerError))
	} else {
		var account entity.Account

		obj := db.Where("email = ?", p.Email).First(&account)

		switch obj.RecordNotFound() {
		case true:
			account.Name = p.Name
			account.Email = p.Email
			account.Password = tools.MD5Hash(p.Password, false)
			account.Status = p.Status
			account.Remark = p.Remark

			if err = db.Create(&account).Error; err != nil {
				log.Println(err.Error())
				err = errors.New("failed to create")
			}
		default:
			err = errors.New("account already exists")
		}
		return err
	}
}

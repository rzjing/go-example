/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         handler.go
@ Create Time:  2020/4/30 17:28
@ Software:     GoLand
*/

package account

import (
	"errors"
	"fmt"
	"go-example/models/entity"
	"go-example/models/mysql"
	"go-example/tools"
	"log"
	"net/http"
	"strconv"
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

		var queryWhere string
		if p.Name != "" {
			queryWhere = fmt.Sprintf("name like '%s'", "%"+p.Name+"%")
		} else if p.Email != "" {
			queryWhere = fmt.Sprintf("email like '%s'", "%"+p.Email+"%")
		} else {
			queryWhere = "1 = 1"
		}

		db.Table("account").Count(&count).
			Select("id, name, email, status, unix_timestamp(created_at) created_at, unix_timestamp(updated_at) updated_at").
			Where(queryWhere).
			Order(p.Sort).Offset(p.Offset).Limit(p.PageSize).Scan(&accounts)

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
			account.Remark = p.Remark

			status, _ := strconv.Atoi(p.Status)
			account.Status = int32(status)

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

func putAccount(id string, p *putParams) error {
	if db, err := mysql.GetConn(); err != nil {
		log.Println(err.Error())
		return errors.New(http.StatusText(http.StatusInternalServerError))
	} else {
		var account entity.Account

		obj := db.Where("id = ?", id).First(&account)

		switch obj.RecordNotFound() {
		case true:
			err = errors.New("account does not exist")
		default:
			account.Name = p.Name

			if p.Password != "" {
				account.Password = tools.MD5Hash(p.Password, false)
			}

			status, _ := strconv.Atoi(p.Status)
			account.Status = int32(status)

			if p.Remark != "" {
				account.Remark = p.Remark
			}

			if err = db.Save(&account).Error; err != nil {
				log.Println(err.Error())
				err = errors.New("failed to change")
			}
		}
		return err
	}
}

func delAccount(id string) error {
	if db, err := mysql.GetConn(); err != nil {
		log.Println(err.Error())
		return errors.New(http.StatusText(http.StatusInternalServerError))
	} else {
		idInt32, _ := strconv.Atoi(id)
		obj := db.Delete(&entity.Account{ID: int32(idInt32)})

		if obj.RowsAffected != 1 {
			err = errors.New("account does not exist")
		}
		return err
	}
}

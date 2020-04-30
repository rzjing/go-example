/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         table.go
@ Create Time:  2020/4/30 15:53
@ Software:     GoLand
*/

package entity

import "time"

// Table: example.account
type Account struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Status    int32     `json:"status"`
	Remark    string    `json:"remark"`
	CreatedAt time.Time `json:"created_at" gorm:"-"`
	UpdatedAt time.Time `json:"updated_at" gorm:"-"`
}

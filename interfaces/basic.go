/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         basic.go
@ Create Time:  2020/5/3 10:43
@ Software:     GoLand
*/

package interfaces

type Params struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
	Offset   int    `json:"-" form:"-"`
	Sort     string `json:"sort" form:"sort"`
}

func (p *Params) Init() {
	if p.Page == 0 {
		p.Page = 1
	}

	if p.PageSize <= 0 || p.PageSize > 100 {
		p.PageSize = 10
	}

	if p.Sort == "" {
		p.Sort = "updated_at desc"
	}

	p.Offset = (p.Page - 1) * p.PageSize
}

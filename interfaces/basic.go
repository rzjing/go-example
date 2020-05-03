/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         basic.go
@ Create Time:  2020/5/3 10:43
@ Software:     GoLand
*/

package interfaces

type Params struct {
	Page     int
	PageSize int
	Offset   int
	Sort     string
}

func (p *Params) Init() {
	if p.Page == 0 {
		p.Page = 1
	}

	if p.PageSize <= 10 {
		p.PageSize = 10
	}

	if p.Sort == "" {
		p.Sort = "updated_at desc"
	}

	p.Offset = (p.Page - 1) * p.PageSize
}

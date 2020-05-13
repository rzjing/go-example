/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         struct_test.go
@ Create Time:  2020/5/12 17:58
@ Software:     GoLand
*/

package example

import (
	"testing"
)

func TestNewDoOptions(t *testing.T) {
	optionTests := [][]DoOptions{
		[]DoOptions{
			DoKey("token"),
			DoValue("3bc728b2d9feafd17c3e2550eecfd942"),
			DoExpire(60),
		},
		[]DoOptions{
			DoKey("devices"),
			DoValue([]string{"1", "2", "3", "4", "5"}),
			DoExpire(60),
		},
		[]DoOptions{
			DoKey("jobs"),
			DoValue(map[string]interface{}{"id": 1, "name": "wang", "age": 18}),
			DoExpire(120),
		},
		[]DoOptions{
			DoKey("code"),
			DoValue(200),
		},
		[]DoOptions{
			DoKey("client"),
		},
	}

	for _, options := range optionTests {
		doOptions := NewDoOptions(options...)
		t.Logf("the doOptions is %+v", doOptions)
	}
}

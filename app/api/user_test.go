package api

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("hi")

	if r, err := g.Client().Post("http://0.0.0.0:8199/user/signup",
		g.Map{
			"Passport":  "knmkmn",
			"Password":  "123123",
			"Password2": "123123",
		},
	); err != nil {
		panic(err)
	} else {
		defer r.Close()

		fmt.Println("res: ", r.ReadAllString())
		fmt.Println(r.StatusCode)
		r.RawDump()
	}

}

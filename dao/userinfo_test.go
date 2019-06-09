package dao

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUserinfo(t *testing.T) {
	user:=&Userinfo{}
	s:="{\"username\":\"duanpeng\"," +
		"\"password\":\"123456\"}"
	err:=json.Unmarshal([]byte(s),user)
	if err!=nil {
		fmt.Println("parse bytes to bean error",err)
	}
	fmt.Println(user.Password,user.Username)
}
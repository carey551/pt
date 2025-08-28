package main

import (
	"project/paymoneyapi"
	_ "project/userApi"
)

func main() {
	// userApi.UserloginY1("918281997139", "qwer1234")
	paymoneyapi.ManualRecharge(1736298, 667, 0)
	// // time.Sleep(time.Second * 3)
	// userApi.AddUserRequest()
}

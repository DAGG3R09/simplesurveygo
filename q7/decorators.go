package q7

import (
	"fmt"
	"net/http"
	"simplesurveygo/dao"
	"simplesurveygo/servicehandlers"
	"time"
)

func logTime(function func()) func() {
	return func() {
		start := time.Now()
		function()
		end := time.Now()
		diff := end.Sub(start)
		fmt.Println("Time Taken: ", diff)
	}
}

// func Authenticate(function func(r *http.Request)) func(interface{}) interface{} {

func Authenticate(function func(), r *http.Request) func(interface{}) interface{} {
	return func(interface{}) interface{} {
		tokens, _ := r.Header["Token"]
		token := tokens[0]

		resp := dao.GetSessionDetails(token)
		if (resp == dao.UserCredentials{}) {
			return (servicehandlers.SimpleBadRequest("UnAuthorized Access"))
		}
		return function
	}
}

func optionHandler(function func(), r *http.Request) interface{} {

	return func(interface{}) interface{} {
		if r.Method == "OPTIONS" {
			return servicehandlers.Response200OK("All Ok")
		}
		return function
	}
}

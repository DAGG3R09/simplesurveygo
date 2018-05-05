package q5

/*
package servicehandlers

import (
	"net/http"
)

type HttpServiceHandler interface {
	Get(*http.Request) SrvcRes
	Put(*http.Request) SrvcRes
	Post(*http.Request) SrvcRes
}

var maxCount = 10
var req = make(q5.Semaphore, maxCount)

func methodRouter(p HttpServiceHandler, w http.ResponseWriter, r *http.Request) interface{} {
	var response interface{}

	req.P(1) // Acquire Resource Q5
	if r.Method == "GET" {
		response = p.Get(r)
	} else if r.Method == "PUT" {
		response = p.Put(r)
	} else if r.Method == "POST" {
		response = p.Post(r)
	}

	req.V(1) // Release resource before response -Q5
	return response
}
*/

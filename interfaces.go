package main

import "net/http"

type API interface {
	create() (http.Response, error)
	fetch() (http.Response, error)
	delete() (http.Response, error)
}

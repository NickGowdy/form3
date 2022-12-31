package main

import "net/http"

type Form3 interface {
	create() (http.Response, error)
	fetch() (http.Response, error)
	delete() (http.Response, error)
}

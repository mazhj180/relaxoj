package main

import "mazhj.com/relaxoj/router"

func main() {
	r := router.Router()
	err := r.Run()
	if err != nil {
		return
	}
}

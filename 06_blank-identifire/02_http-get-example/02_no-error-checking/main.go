package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	// i am just putting all errors in
	// my blank identifire. Because i don't
	// want to care about my errors.
	res, _ := http.Get("http://www.mcleods.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}

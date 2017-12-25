package app

import (
	"io"
	"net/http"
	"fmt"
)

// path: /app/hello
func AppHelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s",req.URL)
	io.WriteString(w, "hello, world.")
}
package main

import (
	"log"
	"net/http"
	"fmt"
	"api"
	"app"
)

func main() {
	fmt.Printf("Hello, world!\n我启动成功啦!!!\n")
	http.HandleFunc("/api", api.ApiServer)
	http.HandleFunc("/app",app.AppHelloServer)
	log.Fatal(http.ListenAndServe(":8888",nil))
}

func developNewCommit1() {
	fmt.Printf("developNewCommit1")
}

func developNewCommit2() {
	fmt.Printf("developNewCommit2")
}

func developNewCommit3() {
	fmt.Printf("developNewCommit3")
}

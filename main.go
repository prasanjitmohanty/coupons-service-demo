package main

import (
	"coupons-service-demo/handlers"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/coupons/", handlers.CouponsRouter)
	http.HandleFunc("/coupons", handlers.CouponsRouter)
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

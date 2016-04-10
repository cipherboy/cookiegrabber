package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := r.Cookies()
		fmt.Println("=====", time.Now(), "::::", r.RemoteAddr, "=====")
		for i := range c {
			fmt.Println("\t", r.RemoteAddr, "::::", c[i].Name, "::::", c[i].Value)
		}
	})

	if len(os.Args) != 3 {
		fmt.Println("Usage: cookiedump ./public.pem ./private.key")
	}

	fmt.Println("Listening on [::]:8443...")
	err := http.ListenAndServeTLS("[::]:8443", os.Args[1], os.Args[2], nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
		return
	}
}

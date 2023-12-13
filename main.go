package main

import (
	"fmt"
	"net/http"
	c "testDockerfile/config"
)

func init() {
	_ = c.ReadConfigAndSetup("config", c.Config)
	fmt.Printf("handle: %+v\n", c.Config)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(c.Config.Handle))
}

func main() {
	fmt.Println("server start")
	http.HandleFunc("/hello", helloHandler)

	port := 8080
	fmt.Printf("Starting server on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)

	}

}

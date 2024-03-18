package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// func init() {
// 	key := make([]byte, 64)
// 	if _, err := rand.Read(key); err != nil {
// 		log.Fatal(err)
// 	}

// 	string64 := base64.StdEncoding.EncodeToString(key)
// 	fmt.Println(string64)
// }

func main() {
	config.Load()
	r := router.Generate()
	fmt.Printf("serve on port %d \n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}

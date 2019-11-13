package main

import (
	"fmt"
	"github.com/MaJloe3Jlo/perx_test_done/data"
	"github.com/MaJloe3Jlo/perx_test_done/routes"
	"github.com/MaJloe3Jlo/perx_test_done/util"
	"log"
	"net/http"
)

// Before we go need to create an keys. Or to find it in Redis.
func init() {

	connect := util.RedisConnect()
	defer connect.Close()

	// Check for existing of keys in Redis.
	kys, _ := connect.Do("GET", "key:1")

	// If key 1 exist - ignore creation of keys. Another way
	if kys == nil {
		var j int
		fmt.Print("Enter number of keys to create (e.g. 10): ")
		_, err := fmt.Scanf("%d", &j)
		if err != nil {
			log.Println("Wrong input")
		}
		for i := 1; i < j+1; i++ {
			newKey := data.Key{Id: i, Key: util.RandStringRunes(), Status: "not issued"}
			data.CreateKey(&newKey)
		}
		log.Println("Keys created in Redis.")
	} else {
		log.Println("Keys exist.")
	}
}

// Start of the app.
func main() {

	// Create router.
	r := routes.Router()

	// Starting server with log.
	log.Println("Server started on http://localhost:7100")
	log.Fatal(http.ListenAndServe(":7100", r))
}

package handler

import (
	"fmt"
	"github.com/MaJloe3Jlo/perx_test_done/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// func KeyOff is a off key if it's already issued.
func KeyOff(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	keyId, err := strconv.Atoi(vars["keyID"])
	if err != nil {
		log.Println(err)
	}

	fmt.Fprint(w, data.KeyOff(keyId))
}

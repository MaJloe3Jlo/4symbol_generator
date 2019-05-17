package handler

import (
	"fmt"
	"github.com/MaJloe3Jlo/perx_test_done/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// func KeyInfo is give an information of a key to user.
func KeyInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	keyId, err := strconv.Atoi(vars["keyID"])
	if err != nil {
		log.Println(err)
	}

	fmt.Fprint(w, data.KeyInfo(keyId))
}

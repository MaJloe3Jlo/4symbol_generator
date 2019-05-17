package handler

import (
	"fmt"
	"github.com/MaJloe3Jlo/perx_test_done/data"
	"net/http"
)

// func Key is give an key to user there is not issued.
func Key(w http.ResponseWriter, r *http.Request) {
	key := data.GetKey()
	if key.Status != "" {
		fmt.Fprintf(w, "<b>Your key is: </b> %s, <b>ID of key: </b> %d", key.Key, key.Id)
	} else {
		fmt.Fprint(w, "<b>No available keys right now</b>")
	}
}

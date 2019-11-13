package handler

import (
	"encoding/json"
	"fmt"
	"github.com/MaJloe3Jlo/perx_test_done/data"
	"io/ioutil"
	"net/http"
)

// func KeyOff is a off key if it's already issued.
func KeyOff(w http.ResponseWriter, r *http.Request) {
	var key data.Key

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &key); err != nil {
		fmt.Fprint(w, "Key ID is not sended in request")
		return
	}

	if key.Id == 0 {
		fmt.Fprint(w, "Key ID is not sended in request")
		return
	}

	fmt.Fprint(w, data.KeyOff(key.Id))
}

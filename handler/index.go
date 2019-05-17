package handler

import (
	"fmt"
	"net/http"
)

// func Index is handle the main page of app.
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1><b><center>Key generator API v.0.0.1 by m3</center></b></h1><b>GET methods:</b><ul><li>path '/' - is a main information of app.</li><li>path '/key' - is gave you 4-symbol key with status 'not issued' and ID of it.</li><li>path '/key/{id}' - an information of key by ID(you get by 'key' path).</li><li>path '/statistics' - is given information about keys. How many: 'issued', 'not issued' or 'off'.</li></ul><b>POST method:</b><ul><li>path '/key/{id}' - is gave status 'off' to key, only if it's had status 'issued'.")
}

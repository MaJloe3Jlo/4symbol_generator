package handler

import (
	"fmt"
	"net/http"
)

// func Index is handle the main page of app.
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<h1><b><center>Key generator API v.0.0.1 by m3</center></b></h1><b>GET methods:</b><ul><li>path '/' - is an app's main information.</li><li>path '/key' - provides 4-symbol key with ID and status 'not issued' to user.</li><li>path '/key/{id}' - provides key's information by ID (use method path '/key' to get ID).</li><li>path '/statistics' - provides information about keys. How many: 'issued', 'not issued' or 'off'.</li></ul><b>POST method:</b><ul><li>path '/key' - receive Content-Type:application/json body (example: {"id":1}) to change key's status to 'off', only if it has status 'issued'.</li></ul>`)
}

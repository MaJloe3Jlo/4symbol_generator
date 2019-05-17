package handler

import (
	"fmt"
	"github.com/MaJloe3Jlo/perx_test_done/data"
	"net/http"
)

// func Statistics given information about keys.
func Statistics(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<b>Statistic of keys: </b>", data.GetStatistics())
}

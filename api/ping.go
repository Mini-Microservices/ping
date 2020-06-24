package ping

import (
	"encoding/json"
	"net/http"

	"github.com/HurricaneInteractive/goping"
)

type response struct {
	Code string `json:"code"`
	URL  string `json:"url"`
}

func enableCors(w *http.ResponseWriter) {
	// allows all origins but only for "POST" methods
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST")
}

// Ping other sites
func Ping(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method != "POST" {
		// Throw error if the method isnt a POST
		http.Error(w, "unsupported method type", http.StatusBadRequest)
		return
	}

	// Get the url from the request and pings the site
	url := r.FormValue("url")
	pingRes := goping.GoPing(url)
	res := response{pingRes, url}

	// encode the struct to json
	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// send the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

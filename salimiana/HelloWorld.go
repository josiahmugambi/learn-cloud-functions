// Package salimiana contains an HTTP Cloud Function.
package salimiana

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Gotea, World!" if there isn't one.
func HelloWorld(w http.ResponseWriter, r *http.Request) {

	var d struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Gotea World!")
		log.Error("Error: ", err)
		return
	}
	if d.Message == "" {
		fmt.Fprint(w, "Gotea World!")
		log.Info("d.Message blank")
		return
	}
	fmt.Fprint(w, html.EscapeString(d.Message))
}

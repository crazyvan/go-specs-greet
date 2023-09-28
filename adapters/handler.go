package adapters

import (
	"fmt"
	"net/http"

	"github.com/crazyvan/go-specs-greet/domain/interactions"
)

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, interactions.Greet(name))
}

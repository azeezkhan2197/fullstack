package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/azeezkhan2197/fullstack/src/api/router"
)

func Run() {
	fmt.Println("Running...")
	r := router.New()
	log.Fatal(http.ListenAndServe(":8080", r))
}

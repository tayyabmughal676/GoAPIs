package api

import (
	"fmt"
	"log"
	"net/http"
	"src/api/router"
	"src/config"
)

func Run() {
	config.Load()
	fmt.Printf("\n\tListening [::]:%d\n", config.PORT)
	Listen(config.PORT)
}

func Listen(port int)  {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

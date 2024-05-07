package main

import (
	"fmt"
	"os"
	"poker-sim/internal/server"

	"github.com/labstack/gommon/log"
)

func main() {

	server := server.NewServer()

	fmt.Printf("server running on %s", os.Getenv("PORT"))
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf(fmt.Sprintf("cannot start server: %s", err))
	}
}

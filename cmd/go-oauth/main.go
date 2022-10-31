package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	config "github.com/md/go-auth/internal/configs"
	"github.com/md/go-auth/internal/routes"
)

func main() {
	// Load env
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load Config
	config.LoadConfig()

	r := mux.NewRouter()
	routes.Router(r)

	var port = os.Getenv("PORT")
	fmt.Println("server running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

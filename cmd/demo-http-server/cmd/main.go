package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"

	// CA bundle for FROM Scratch
	_ "golang.org/x/crypto/x509roots/fallback"

	"github.com/spf13/viper"
)

//go:embed static
var staticFiles embed.FS

//go:embed templates
var templateFiles embed.FS

func main() {
	// Set up Viper
	viper.SetDefault("PORT", "3000")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// Read configuration
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("No .env file found, using default or environment variables.")
		} else {
			// Config file was found but another error was produced
			log.Fatalf("Error reading config file: %s", err)
		}
	}

	port := viper.GetString("PORT")

	mux := http.NewServeMux()

	// Serve static files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFiles))))

	// Serve templates
	tmpl := template.Must(template.ParseFS(templateFiles, "templates/index.html"))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

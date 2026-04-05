package main

import (
	"fmt"
	"log"
	"net/http"
	"vibefs/src/config"

	"github.com/go-chi/chi/v5"
)

func main() {
	var cfg config.Config
	if err := config.Load(&cfg); err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	// redirect root → APP_ROOT
	if cfg.APP_PATH != "/" {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, cfg.APP_PATH, http.StatusFound)
		})
	}

	// serve the static app
	fs := http.FileServer(http.Dir("./public"))
	r.Handle(cfg.APP_PATH+"*", http.StripPrefix(cfg.APP_PATH, fs))

	fmt.Printf("Server running at http://localhost:%s%s\n", cfg.APP_PORT, cfg.APP_PATH)
	log.Fatal(http.ListenAndServe(":"+cfg.APP_PORT, r))
}

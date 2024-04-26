package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src *; script-src 'self' https://cdn.jsdelivr.net/npm/dompurify/dist/purify.min.js https://unpkg.com/htmx.org@1.9.11/dist/htmx.js 'unsafe-eval'; style-src *")
		// Log request details before handling
		log.Printf("%s - %s %s", r.RemoteAddr, r.Method, r.URL)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Create a new router instance
	router := mux.NewRouter()

	// Static assets
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))

	router.HandleFunc("/randomimage", randomimage).Methods("GET")
	router.HandleFunc("/share/{path:.*}", shareimage).Methods("GET")
	router.HandleFunc("/{path:.*.png}", renderimage).Methods("GET")

	// Catch all index.html
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates := template.Must(template.New("").ParseGlob("public/templates/*.html"))
		err := templates.ExecuteTemplate(w, "start.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}).Methods("GET")

	server := &http.Server{
		Addr:         "0.0.0.0:3000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 20,
		Handler:      loggingMiddleware(router), // Pass our instance of gorilla/mux in.
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("listening on 3000")
	}
}

func renderimage(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.New("").ParseGlob("public/templates/*.html"))
	code := mux.Vars(r)["path"]
	err := templates.ExecuteTemplate(w, "image.html", "public/assets/"+code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func randomimage(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.New("").ParseGlob("public/templates/*.html"))
	images, err := filepath.Glob("./public/assets/*.png")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	selected := images[rand.Intn(len(images))]
	w.Header().Set("HX-Push-Url", "/share/"+filepath.Base(selected))
	err = templates.ExecuteTemplate(w, "image.html", selected)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func shareimage(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.New("").ParseGlob("public/templates/*.html"))
	code := mux.Vars(r)["path"]
	err := templates.ExecuteTemplate(w, "share.html", code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

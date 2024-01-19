package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type mailer struct {
	from     string
	password string
	to       string
	smtpHost string
	smtpPort string
}

func (m *mailer) send(w http.ResponseWriter, r *http.Request) {
}

func main() {
	var m mailer
	var apiPort uint
	flag.StringVar(&m.from, "from", "", "from email address")
	flag.StringVar(&m.password, "password", "", "from email password")
	flag.StringVar(&m.to, "to", "", "to email address")
	flag.StringVar(&m.smtpHost, "smtpHost", "", "smtp host")
	flag.StringVar(&m.smtpPort, "smtpPort", "", "smtp port")
	flag.UintVar(&apiPort, "apiPort", 5000, "api port")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Post("/send", m.send)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", apiPort), r))
}

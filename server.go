package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/cinn-labs/auth"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Server struct {
	Router *mux.Router
	DB     *sql.DB
	Auth   *auth.Model
}

func New(db *sql.DB, jwtSignature string) Server {
	s := Server{mux.NewRouter(), db, auth.Generate(jwtSignature)}
	return s
}

func (s *Server) Run(address string) {
	log.Printf("SERVER STARTED AT PORT: %s", address)
	methods := handlers.AllowedMethods([]string{"DELETE", "GET", "HEAD", "POST", "PUT", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(address, handlers.LoggingHandler(
		os.Stdout, handlers.CORS(methods, headers, origins)(s.Router))))
}

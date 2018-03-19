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
	Origin []string
}

func New(db *sql.DB, jwtSignature string) Server {
	s := Server{mux.NewRouter(), db, auth.Generate(jwtSignature), []string{"*"}}
	return s
}

func (s *Server) Run(port string) {
	log.Printf("SERVER STARTED AT PORT: %s", port)
	methods := handlers.AllowedMethods([]string{"DELETE", "GET", "HEAD", "POST", "PUT", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"Origins", "Authorization", "X-Requested-With", "Content-Type", "Auth"})
	origins := handlers.AllowedOrigins(s.Origin)

	log.Fatal(http.ListenAndServe(port, handlers.LoggingHandler(
		os.Stdout, handlers.CORS(origins, headers, methods)(s.Router))))
}

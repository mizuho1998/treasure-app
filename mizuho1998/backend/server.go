package server

import (
	"fmt"

	"github.com/voyagegroup/treasure-app/sample"

	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/justinas/alice"

	"firebase.google.com/go/auth"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
	"github.com/voyagegroup/treasure-app/controller"
	db2 "github.com/voyagegroup/treasure-app/db"
	"github.com/voyagegroup/treasure-app/firebase"
	"github.com/voyagegroup/treasure-app/middleware"
)

type Server struct {
	db         *sqlx.DB
	router     *mux.Router
	authClient *auth.Client
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init(datasource string) {
	authClient, err := firebase.InitAuthClient()
	if err != nil {
		log.Fatalf("failed init auth client. %s", err)
	}
	s.authClient = authClient

	db := db2.NewDB(datasource)
	dbcon, err := db.Open()
	if err != nil {
		log.Fatalf("failed db init. %s", err)
	}
	s.db = dbcon
	s.router = s.Route()
}

func (s *Server) Run(addr string) {
	log.Printf("Listening on port %s", addr)
	err := http.ListenAndServe(
		fmt.Sprintf(":%s", addr),
		handlers.CombinedLoggingHandler(os.Stdout, s.router),
	)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Route() *mux.Router {
	// authMiddleware := middleware.NewAuth(s.authClient, s.db)
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Authorization"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
	})

	commonChain := alice.New(
		middleware.RecoverMiddleware,
		corsMiddleware.Handler,
	)

	// authChain := commonChain.Append(
	// 	authMiddleware.Handler,
	// )

	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/public").Handler(commonChain.Then(sample.NewPublicHandler()))
	r.Methods(http.MethodGet).Path("/private").Handler(commonChain.Then(sample.NewPrivateHandler(s.db)))

	keyboardController := controller.NewKeyboard(s.db)
	r.Methods(http.MethodGet).Path("/keyboard").Handler(commonChain.Then(AppHandler{keyboardController.AllKeyboard}))
	r.Methods(http.MethodPost).Path("/keyboard").Handler(commonChain.Then(AppHandler{keyboardController.Serch}))
	r.Methods(http.MethodGet).Path("/keyboard/{id}").Handler(commonChain.Then(AppHandler{keyboardController.FindKeyboard}))

	questionController := controller.NewQuestion(s.db)
	r.Methods(http.MethodGet).Path("/question/{id}").Handler(commonChain.Then(AppHandler{questionController.Find}))

	return r
}
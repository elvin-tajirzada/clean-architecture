package app

import (
	"fmt"
	"github.com/elvin-tacirzade/clean-architecture/pkg/config"
	"github.com/elvin-tacirzade/clean-architecture/pkg/controllers"
	"github.com/elvin-tacirzade/clean-architecture/pkg/db"
	"github.com/elvin-tacirzade/clean-architecture/pkg/repositories"
	"github.com/elvin-tacirzade/clean-architecture/pkg/services"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
)

type App struct {
	Router *mux.Router
	DB     *sqlx.DB
}

func (a *App) Start() {
	a.initialize()
	a.routes()
	a.run(os.Getenv("ADDR"))
}

func (a *App) initialize() {
	var err error
	err = config.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()

	a.DB, err = db.ConnectPostgres(os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASS"), os.Getenv("POSTGRES_DB_NAME"), os.Getenv("POSTGRES_SSL_MODE"))
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) routes() {
	usersRoute := initUsers(a.DB)
	a.Router.HandleFunc("/api/users", usersRoute.GetAllUsers).Methods(http.MethodGet)
	a.Router.HandleFunc("/api/users/{id:[0-9]+}", usersRoute.FindById).Methods(http.MethodGet)
	a.Router.HandleFunc("/api/users", usersRoute.InsertUser).Methods(http.MethodPost)
	a.Router.HandleFunc("/api/users/{id:[0-9]+}", usersRoute.DeleteUser).Methods(http.MethodDelete)
}

func (a *App) run(addr string) {
	fmt.Printf("Server started at %s\n", addr)
	log.Fatal(http.ListenAndServe(":"+addr, a.Router))
}

func initUsers(db *sqlx.DB) controllers.UsersController {
	usersRepository := repositories.NewUsersRepository(db)
	usersService := services.NewUsersServices(usersRepository)
	usersRoute := controllers.NewUsersController(usersService)
	return usersRoute
}

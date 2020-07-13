package router

import (
	"database/sql"
	"fmt"
	"login_page_gerin/domains/user"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	USER_MAIN_ROUTE = "/user"
)

type ConfigRouter struct {
	DB     *sql.DB
	Router *mux.Router
}

func (ar *ConfigRouter) InitRouter() {
	user.InitUserRoute(USER_MAIN_ROUTE, ar.DB, ar.Router)
	ar.Router.NotFoundHandler = http.HandlerFunc(notFound)
}

// NewAppRouter for creating new Route
func NewAppRouter(db *sql.DB, r *mux.Router) *ConfigRouter {
	return &ConfigRouter{
		DB:     db,
		Router: r,
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, `<h1>404 Status Not Found</h1>`)
}

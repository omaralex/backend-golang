package web

import (
	"backend-kata/config"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
)

type Server struct {
	httpSrv *http.Server
	router  *mux.Router
	cfg     *config.Config
}

func NewServer(cfg *config.Config) *Server {
	router := mux.NewRouter()
	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}

	return &Server{
		httpSrv: httpSrv,
		router:  router,
		cfg:     cfg,
	}
}

func (s *Server) ListenAndServe() {
	log.Info().Msgf("app running in %s with version %s", strconv.Itoa(s.cfg.Server.Port), s.cfg.Server.Version)
	err := s.httpSrv.ListenAndServe()
	if err != nil {
		log.Err(err).Msgf("server: Unexpected error")
	}
}

func (s *Server) Run() {
	go func() {
		s.ListenAndServe()
	}()
	// Wait for terminate signal to shut down server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func (s *Server) RegisterRoutes(container Container) {
	routes := initRoutes(
		container.GetAuthHandler(),
		container.GetDrugHandler(),
		container.GetVaccinationHandler(),
		container.GetAuthMiddleware(),
	)

	for _, route := range routes {
		s.router.Handle(route.Path, buildChain(route.Handler, route.Middlewares...)).Methods(route.Method)
	}
}

type middleware func(http.HandlerFunc) http.HandlerFunc

func buildChain(handlerFunc http.HandlerFunc, middlewares ...middleware) http.HandlerFunc {
	if len(middlewares) == 0 {
		return handlerFunc
	}

	return middlewares[0](buildChain(handlerFunc, middlewares[1:cap(middlewares)]...))
}

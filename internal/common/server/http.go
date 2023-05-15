package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"

	"github.com/duckhue01/wild-workouts/internal/common/auth"
	"github.com/duckhue01/wild-workouts/internal/common/logs"
)

type Conf struct {
	CreateHandler func(router chi.Router) http.Handler
	Port          int
	Parser        auth.Parser
}

func Run(c Conf) {
	apiRouter := chi.NewRouter()
	setMiddleWares(apiRouter, c.Parser)

	rootRouter := chi.NewRouter()
	// we are mounting all APIs under /api path
	rootRouter.Mount("/api", c.CreateHandler(apiRouter))

	logrus.Info("starting http server")

	err := http.ListenAndServe(fmt.Sprintf(":%d", c.Port), rootRouter)
	if err != nil {
		logrus.WithError(err).Panic("unable to start http server")
	}
}

func setMiddleWares(router *chi.Mux, p auth.Parser) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(logs.NewStructuredLogger(logrus.StandardLogger()))
	router.Use(middleware.Recoverer)

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	router.Use(middleware.NoCache)

	if p != nil {
		router.Use(auth.AuthMiddleware{P: p}.Middleware)
	}
}

// todo: add CORS middleware
// func addCorsMiddleware(router *chi.Mux) {
// 	allowedOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ";")
// 	if len(allowedOrigins) == 0 {
// 		return
// 	}

// 	corsMiddleware := cors.New(cors.Options{
// 		AllowedOrigins:   allowedOrigins,
// 		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
// 		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
// 		ExposedHeaders:   []string{"Link"},
// 		AllowCredentials: true,
// 		MaxAge:           300,
// 	})
// 	router.Use(corsMiddleware.Handler)
// }

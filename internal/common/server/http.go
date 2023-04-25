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

func Run(createHandler func(router chi.Router) http.Handler, port int, p auth.Parser) {
	run(fmt.Sprintf(":%d", port), createHandler, p)
}

func run(addr string, createHandler func(router chi.Router) http.Handler, p auth.Parser) {
	apiRouter := chi.NewRouter()
	setMiddleWares(apiRouter, p)

	rootRouter := chi.NewRouter()
	// we are mounting all APIs under /api path
	rootRouter.Mount("/api", createHandler(apiRouter))

	logrus.Info("starting HTTP server")

	err := http.ListenAndServe(addr, rootRouter)
	if err != nil {
		logrus.WithError(err).Panic("unable to start HTTP server")
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

	// 	// todo: implement mock
	// 	// if mockAuth, _ := strconv.ParseBool(os.Getenv("MOCK_AUTH")); mockAuth {
	// 	// 	router.Use(auth.HttpMockMiddleware)
	// 	// 	return
	// 	// }
	router.Use(auth.AuthMiddleware{P: p}.Middleware)
}

// // todo: add CORS middleware
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

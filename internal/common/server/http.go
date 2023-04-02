package server

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

func Run(createHandler func(router chi.Router) http.Handler) {
	run(":"+os.Getenv("PORT"), createHandler)
}

func run(addr string, createHandler func(router chi.Router) http.Handler) {
	apiRouter := chi.NewRouter()
	setMiddleWares(apiRouter)

	rootRouter := chi.NewRouter()
	// we are mounting all APIs under /api path
	rootRouter.Mount("/api", createHandler(apiRouter))

	logrus.Info("Starting HTTP server")

	err := http.ListenAndServe(addr, rootRouter)
	if err != nil {
		logrus.WithError(err).Panic("Unable to start HTTP server")
	}
}

func setMiddleWares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	// router.Use(logs.NewStructuredLogger(logrus.StandardLogger()))
	router.Use(middleware.Recoverer)

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	router.Use(middleware.NoCache)
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
// todo: add auth middleware
// func addAuthMiddleware(router *chi.Mux) {
// 	if mockAuth, _ := strconv.ParseBool(os.Getenv("MOCK_AUTH")); mockAuth {
// 		router.Use(auth.HttpMockMiddleware)
// 		return
// 	}

// 	var opts []option.ClientOption
// 	if file := os.Getenv("SERVICE_ACCOUNT_FILE"); file != "" {
// 		opts = append(opts, option.WithCredentialsFile(file))
// 	}

// 	config := &firebase.Config{ProjectID: os.Getenv("GCP_PROJECT")}
// 	firebaseApp, err := firebase.NewApp(context.Background(), config, opts...)
// 	if err != nil {
// 		logrus.Fatalf("error initializing app: %v\n", err)
// 	}

// 	authClient, err := firebaseApp.Auth(context.Background())
// 	if err != nil {
// 		logrus.WithError(err).Fatal("Unable to create firebase Auth client")
// 	}

// 	router.Use(auth.FirebaseHttpMiddleware{authClient}.Middleware)
// }

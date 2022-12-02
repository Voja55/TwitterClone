package main

import (
	"Profile/db"
	"Profile/handlers"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "[auth-api] ", log.LstdFlags)

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	port := os.Getenv("app_port")
	if len(port) == 0 {
		port = "8080"
	}

	// NoSQL: Initialize Product Repository store
	userRepo, err := db.NewUserRepoDB(timeoutContext, logger)
	if err != nil {
		logger.Fatal(err)
	}
	defer userRepo.Disconnect(timeoutContext)

	// NoSQL: Checking if the connection was established
	userRepo.Ping()

	//Initialize the handler and inject said logger
	usersHandler := handlers.NewUsersHandler(logger, userRepo)

	//Initialize the router and add a middleware for all the requests
	routerUser := mux.NewRouter()
	routerUser.Use(usersHandler.MiddlewareContentTypeSet)

	getUsersRouter := routerUser.Methods(http.MethodGet).Subrouter()
	getUsersRouter.HandleFunc("/profile", usersHandler.GetProfile)

	postUserRouter := routerUser.Methods(http.MethodPost).Subrouter()
	postUserRouter.HandleFunc("/profile", usersHandler.Register)
	postUserRouter.Use(usersHandler.MiddlewareUsersValidation)

	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"https://localhost:4200/"}))

	server := http.Server{
		Addr:         ":" + port,        // Addr optionally specifies the TCP address for the server to listen on, in the form "host:port". If empty, ":http" (port 80) is used.
		Handler:      cors(routerUser),  // handler to invoke, http.DefaultServeMux if nil
		IdleTimeout:  120 * time.Second, // IdleTimeout is the maximum amount of time to wait for the next request when keep-alives are enabled.
		ReadTimeout:  2 * time.Second,   // ReadTimeout is the maximum duration for reading the entire request, including the body. A zero or negative value means there will be no timeout.
		WriteTimeout: 2 * time.Second,   // WriteTimeout is the maximum duration before timing out writes of the response.
	}

	logger.Println("Server listening on port", port)

	go func() {
		err := server.ListenAndServeTLS("certificates/group3.crt", "certificates/group3.key")
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT)
	signal.Notify(sigCh, syscall.SIGKILL)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)
	timeoutContext, _ = context.WithTimeout(context.Background(), 30*time.Second)

	if server.Shutdown(timeoutContext) != nil {
		logger.Fatal("Cannot gracefully shutdown...")
	}
}

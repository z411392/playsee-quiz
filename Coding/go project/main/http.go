package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	quiz "playsee.co/interview/modules/quiz/presentation/http"
	auth "playsee.co/interview/utils/auth"
)

var server *http.Server
var addr string

const readTimeout = 15 * time.Second
const writeTimeout = 15 * time.Second

func accessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		log.Printf("%s %s\n", request.Method, request.RequestURI)
		next.ServeHTTP(responseWriter, request)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		apiKey := request.Header.Get("api-key")
		err := auth.ValidateApiKey(apiKey)
		if err != nil {
			http.Error(responseWriter, fmt.Sprint(err), http.StatusForbidden)
			return
		}
		next.ServeHTTP(responseWriter, request)
	})
}

func init() {
	handler := mux.NewRouter()
	handler.StrictSlash(true)
	handler.Use(accessLogMiddleware)
	handler.Use(authMiddleware)
	handler.HandleFunc("/test-1", quiz.Quiz).Methods("POST")
	addr = fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))
	server = &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
}

func startup() {
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func waitForSignal() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Playsee Backend Interview Test server is running at: %s\n", addr)
	<-stop
}

func shutdown() {
	currentContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(currentContext)
	if err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
	log.Println("Server gracefully stopped")
}

func main() {
	go startup()
	waitForSignal()
	shutdown()
}

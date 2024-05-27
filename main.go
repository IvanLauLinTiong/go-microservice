package main

// *******************************************************************
// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/go-chi/chi/v5"
// 	"github.com/go-chi/chi/v5/middleware"
// )

// func main() {
// 	router := chi.NewRouter()
// 	router.Use(middleware.Logger)

// 	router.Get("/hello", basicHandler)

// 	server := &http.Server{
// 		Addr: ":3000",
// 		// Handler: http.HandlerFunc(basicHandler),
// 		Handler: router,

// 	}

// 	err := server.ListenAndServe()
// 	if err != nil {
// 		fmt.Println("failed to lister server: ", err)
// 	}
// }

// func basicHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello world!"))
// }
// *******************************************************************

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/IvanLauLinTiong/go-microservice/application"
)

func main() {
	app := application.New(application.LoadConfig())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	// call cancel at this end of current scope
	// same as calling it at the end
	defer cancel()
	// defer func() { fmt.Println("cancel 1..."); cancel()}()


	// err := app.Start(context.TODO())
	err := app.Start(ctx)

	if err != nil {
		fmt.Println("failed to start the app:", err)
	}
}
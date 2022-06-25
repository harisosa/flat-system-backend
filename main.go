package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/harisosa/flat-system-backend/config"
	database "github.com/harisosa/flat-system-backend/connection"
	"github.com/harisosa/flat-system-backend/flat"
	"github.com/harisosa/flat-system-backend/helper"
	"github.com/harisosa/flat-system-backend/neighborhood"
	"github.com/harisosa/flat-system-backend/seeder"
	"github.com/harisosa/flat-system-backend/user"
)

var (
	addr = ":8080"
)

func main() {
	help := helper.NewHelper()
	l := log.New(os.Stdout, "POS API", log.LstdFlags)

	//help := helper.NewHelper()
	env := config.ENV

	pgsql := database.NewPgsql()
	pgsql.ConnectionPool()

	//pgfunc := dbhelper.NewPgFunction(l)

	db := database.PgConnectionPool

	if env.Debug {
		seeder := seeder.NewPostgresSeeder(db)
		seeder.DatabaseMigration()
		seeder.SeedFile()
	}

	router := mux.NewRouter()

	uRep := user.NewUserRepository()
	uCC := user.NewUserUsecase(db, uRep, help)
	user.NewUserController(router, uCC, help)

	nRep := neighborhood.NewNeighborhoodRepository()
	nCC := neighborhood.NewNeighborhoodUsecase(db, nRep)

	fRep := flat.NewFlatRepository()
	fCC := flat.NewFlatUsecase(db, fRep, nCC)
	flat.NewFlatController(router, fCC, help)

	corsMw := mux.CORSMethodMiddleware(router)
	router.Use(corsMw)

	ch := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Origin", "Authorization", "X-Requested-With", "Content-Type", "Accept"}),
		handlers.AllowedMethods([]string{"OPTIONS", "GET", "POST", "PUT", "DELETE"}),
		handlers.AllowCredentials(),
	)
	// create a new server
	s := http.Server{
		Addr:         addr,              // configure the bind address
		Handler:      ch(router),        // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port " + addr)
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()
	//trap sigterm or interuppt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	//Block until a signal is received
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

}

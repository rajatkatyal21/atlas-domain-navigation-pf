package app

import (
	"context"
	health "dns/internal/health_check"
	dataBank "dns/internal/domain_navigation"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

// Configuration required.
type Config struct {
	Port     int
	Name     string
	Version  string
	SectorId int64
}

// Server Configuration
type Server struct {
	*Config
	StatusHandler           health.Handler
	DataBankLocationHandler dataBank.Handler
}

// NewServer is used for creating the instance of the server.
func NewServer(c *Config) *Server {
	s := &Server{
		Config:        c,
		StatusHandler: &health.Controller{},
		DataBankLocationHandler: &dataBank.Controller{
			DataBankCalculator: &dataBank.DataBankService{
				SectorId: c.SectorId,
			},
		},
	}

	return s
}

// Start the app.
func (s *Server) Start() {

	//Converting the port to string
	port := strconv.Itoa(s.Port)

	//initializing the router
	h := s.InitRouter()

	// creating the server
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: h,
	}

	// start the server in new goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatalf("Error While starting the server at port %s", err.Error())
			}

		}
	}()

	log.Printf("Server started at port %s", port)

	// Setting up signal capturing
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	//waiting for the terminate signal
	sig := <-stop
	log.Printf("signal %+v is triggered", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Shutting down the server.
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown failed %s", err.Error())
	}
}

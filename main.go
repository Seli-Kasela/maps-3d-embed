package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// =====================================
// MODELS
// =====================================

type GeoLocation struct {
	IP        string  `json:"ip"`
	ASN       string  `json:"asn"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	ISP       string  `json:"isp"`
	Org       string  `json:"org"`
}

type MapConfig struct {
	Location GeoLocation `json:"location"`
	Zoom     int         `json:"zoom"`
	Tilt     int         `json:"tilt"`
	Heading  int         `json:"heading"`
}

// =====================================
// SERVER
// =====================================

type Server struct {
	router *mux.Router
	port   string
}

// =====================================
// CREATE SERVER
// =====================================

func NewServer(port string) *Server {

	return &Server{
		router: mux.NewRouter(),
		port:   port,
	}
}

// =====================================
// ROUTES
// =====================================

func (s *Server) setupRoutes() {

	s.router.HandleFunc(
		"/api/geolocation",
		s.getGeolocation,
	).Methods("GET")

	s.router.HandleFunc(
		"/api/map-config",
		s.getMapConfig,
	).Methods("GET")

	s.router.HandleFunc(
		"/",
		s.serveIndex,
	).Methods("GET")

	s.router.PathPrefix("/static/").Handler(

		http.StripPrefix(
			"/static/",
			http.FileServer(
				http.Dir("static"),
			),
		),
	)
}

// =====================================
// GEOLOCATION API
// =====================================

func (s *Server) getGeolocation(
	w http.ResponseWriter,
	r *http.Request,
) {

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	location := GeoLocation{

		IP:  getEnv(
			"DEFAULT_IP",
			"103.137.82.147",
		),

		ASN: getEnv(
			"DEFAULT_ASN",
			"AS138828",
		),

		City: getEnv(
			"DEFAULT_CITY",
			"Badung",
		),

		Country: getEnv(
			"DEFAULT_COUNTRY",
			"Indonesia",
		),

		Latitude:  -8.7933893,
		Longitude: 115.1227501,

		ISP: getEnv(
			"APP_ORG",
			"Seli Kasela Creative Studio",

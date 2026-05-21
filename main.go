package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// GeoLocation represents user's geographical location
type GeoLocation struct {
	IP        string  `json:"ip"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
	ISP       string  `json:"isp"`
}

// MapConfig represents the 3D map configuration
type MapConfig struct {
	APIKey    string      `json:"apiKey"`
	Location  GeoLocation `json:"location"`
	Zoom      int         `json:"zoom"`
	Tilt      int         `json:"tilt"`
	Heading   int         `json:"heading"`
}

// Server represents the application server
type Server struct {
	router *mux.Router
	port   string
}

// NewServer creates a new server instance
func NewServer(port string) *Server {
	return &Server{
		router: mux.NewRouter(),
		port:   port,
	}
}

// setupRoutes configures all API routes
func (s *Server) setupRoutes() {
	s.router.HandleFunc("/api/geolocation", s.getGeolocation).Methods("GET")
	s.router.HandleFunc("/api/map-config", s.getMapConfig).Methods("GET")
	s.router.HandleFunc("/", s.serveIndex).Methods("GET")
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}

// getGeolocation returns user's geolocation based on IP
func (s *Server) getGeolocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get client IP
	ip := getClientIP(r)

	// Mock geolocation data (in production, use MaxMind GeoIP2 or similar service)
	geo := GeoLocation{
		IP:        ip,
		Latitude:  37.7749,  // San Francisco example
		Longitude: -122.4194,
		City:      "San Francisco",
		Country:   "United States",
		ISP:       "Example ISP",
	}

	json.NewEncoder(w).Encode(geo)
}

// getMapConfig returns 3D map configuration
func (s *Server) getMapConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	apiKey := os.Getenv("GOOGLE_MAPS_API_KEY")
	if apiKey == "" {
		http.Error(w, "Google Maps API key not configured", http.StatusInternalServerError)
		return
	}

	config := MapConfig{
		APIKey: apiKey,
		Location: GeoLocation{
			Latitude:  37.7749,
			Longitude: -122.4194,
			City:      "San Francisco",
			Country:   "United States",
		},
		Zoom:    18,
		Tilt:    45,
		Heading: 0,
	}

	json.NewEncoder(w).Encode(config)
}

// serveIndex serves the main HTML file
func (s *Server) serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

// getClientIP extracts the client's IP address from the request
func getClientIP(r *http.Request) string {
	// Check X-Forwarded-For header (for proxies)
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		return xff
	}

	// Check X-Real-IP header
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}

	// Fall back to RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

// Start starts the HTTP server
func (s *Server) Start() error {
	s.setupRoutes()
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(s.router)

	log.Printf("Server starting on port %s", s.port)
	return http.ListenAndServe(":"+s.port, corsHandler)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := NewServer(port)
	if err := server.Start(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

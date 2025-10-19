package httpx

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// Router is a facade for chi.Router
type Router = chi.Router

// Context is a facade for chi.Context
type Context = chi.Context

// Middleware types
type (
	// Handler is a facade for http.Handler
	Handler = http.Handler

	// HandlerFunc is a facade for http.HandlerFunc
	HandlerFunc = http.HandlerFunc
)

// NewRouter creates a new HTTP router
func NewRouter() Router {
	return chi.NewRouter()
}

// NewMux creates a new HTTP router (alias for NewRouter)
func NewMux() Router {
	return chi.NewRouter()
}

// URLParam is a facade for chi.URLParam
func URLParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

// RouteContext is a facade for chi.RouteContext
func RouteContext(r *http.Request) *chi.Context {
	return chi.RouteContext(r.Context())
}

// Middleware exports
var (
	// Logger is middleware that logs HTTP requests
	Logger = middleware.Logger

	// Recoverer is middleware that recovers from panics
	Recoverer = middleware.Recoverer

	// RequestID is middleware that injects a request ID
	RequestID = middleware.RequestID

	// RealIP is middleware that sets RemoteAddr to real client IP
	RealIP = middleware.RealIP

	// Timeout is middleware that cancels context after timeout
	Timeout = middleware.Timeout

	// Throttle is middleware that limits number of in-flight requests
	Throttle = middleware.Throttle

	// Compress is middleware that compresses responses
	Compress = middleware.Compress

	// StripSlashes is middleware that strips trailing slashes
	StripSlashes = middleware.StripSlashes

	// GetReqID gets the request ID from context
	GetReqID = middleware.GetReqID

	// NoCache sets headers to prevent client caching
	NoCache = middleware.NoCache

	// Heartbeat mounts a heartbeat endpoint
	Heartbeat = middleware.Heartbeat

	// AllowContentType limits request content types
	AllowContentType = middleware.AllowContentType

	// AllowContentEncoding limits request content encodings
	AllowContentEncoding = middleware.AllowContentEncoding

	// SetHeader sets a response header
	SetHeader = middleware.SetHeader
)

// CORS creates a new CORS handler with options
func CORS(options cors.Options) func(http.Handler) http.Handler {
	return cors.Handler(options)
}

// DefaultCORS creates a CORS handler with permissive defaults
func DefaultCORS() func(http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Request-ID"},
		ExposedHeaders:   []string{"Link", "X-Request-ID"},
		AllowCredentials: true,
		MaxAge:           300, // 5 minutes
	})
}

// NewCORSHandler creates a new CORS handler (alias for CORS)
func NewCORSHandler(options cors.Options) func(http.Handler) http.Handler {
	return cors.Handler(options)
}

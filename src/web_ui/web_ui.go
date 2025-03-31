package web_ui

import (
	// "io"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	// "path/filepath"
	"strings"

	"github.com/cloakwiss/cobweb/web_ui/messaging"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const DOMAIN = ":8080"
const PUBLIC_PREFIX = "./public"

// THIS ENTIRE THING NEEDS PROPER LOGGING
func Launch() {
	println("Running on: localhost", DOMAIN)
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	// Handling Get Requests
	// ---------------------------------------------------------------------------
	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		HandleGetRequest("/index.html", writer, request)
	})

	router.Get("/styles.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/styles.css")
		w.Header().Set("Content-Type", "text/css")
	})

	router.Get("/script.js", func(writer http.ResponseWriter, request *http.Request) {
		HandleGetRequest("/script.js", writer, request)
	})

	router.Get("/favicon.ico", func(writer http.ResponseWriter, request *http.Request) {
		HandleGetRequest("/favicon.ico", writer, request)
	})

	router.Get("/error.html", func(writer http.ResponseWriter, request *http.Request) {
		HandleGetRequest("/error.html", writer, request)
	})
	// ---------------------------------------------------------------------------

	// Handling Post Requests
	// ---------------------------------------------------------------------------
	router.Post("/archive", func(writer http.ResponseWriter, request *http.Request) {
		ArchiveRequest(writer, request)
	})
	// ---------------------------------------------------------------------------

	// Handling messaging
	// ---------------------------------------------------------------------------
	router.HandleFunc("/messaging", messaging.HandleWebSocket)
	// ---------------------------------------------------------------------------

	// File server routing for being able to download archived epub
	// files should be in ./results
	// ---------------------------------------------------------------------------
	workDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	filesDir := http.Dir(workDir)
	FileServer(router, "/", filesDir)
	// ---------------------------------------------------------------------------

	// Launching the server
	launch_err := http.ListenAndServe(DOMAIN, router)
	if launch_err != nil {
		fmt.Println(err)
	}
}

func CheckValidGetRoutes(route string) bool {
	_, err := os.Stat(route)
	return !errors.Is(err, os.ErrNotExist)
}

func HandleGetRequest(path string, writer http.ResponseWriter, request *http.Request) {
	route := PUBLIC_PREFIX + path

	var body []byte
	if CheckValidGetRoutes(route) {
		bytes, err := os.ReadFile(route)

		if err != nil {
			fmt.Println("requested file is not present: " + route)
			fmt.Println(err)
			body = []byte("Extreme error in server")
		} else {
			body = bytes
		}
	} else {
		bytes, err := os.ReadFile(PUBLIC_PREFIX + "error.html")

		if err != nil {
			fmt.Println("error.html is not present")
			fmt.Println(err)
			body = []byte("Extreme error in server")
		} else {
			body = bytes
		}
	}

	writer.Write(body)
}

// File server for serving resultant epub
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

// Post Request Handling for archive
func ArchiveRequest(writer http.ResponseWriter, request *http.Request) {
	var request_obj WebOptions

	err := json.NewDecoder(request.Body).Decode(&request_obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(request_obj)

	args := WebOptToOpt(request_obj)
	fmt.Println(args.String())

	// Messaging start
	messaging.StartMessaging(writer)
	// ---------------------------------------------
	// Running the application
	result := RunApp(args)
	// ---------------------------------------------
	messaging.EndMessaging()
	// Messaging end

	response_bytes, resp_err := json.Marshal(result)
	fmt.Println(string(response_bytes))
	if resp_err != nil {
		fmt.Println(err)
		response_bytes = []byte("error in response")
	}

	_, write_err := writer.Write(response_bytes)
	if write_err != nil {
		fmt.Println(err)
	}
}

package spiderweb

import(
	"net/http"
	"os"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"ARISTOS/spiderweb/middleware"
	"ARISTOS/spiderweb/handlers"
)

func Run(){
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/signup", handlers.SignUpHandler).Methods("GET", "POST")
	r.HandleFunc("/editor", handlers.EditorHandler).Methods("GET")
	r.HandleFunc("/editor/image-upload", handlers.UploadImageHandler).Methods("GET", "POST")
	r.HandleFunc("/editor/video-upload", handlers.UploadVideoHandler).Methods("GET", "POST")

	// middleware stuffs
	http.Handle("/", middleware.PanicRecoveryHandler(ghandlers.LoggingHandler(os.Stdout, r)))

	// basic server stuffs
	http.ListenAndServe(":8080", nil)
}
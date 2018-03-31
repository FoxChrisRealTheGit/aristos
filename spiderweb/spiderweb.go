package spiderweb

import(
	"net/http"
	"os"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"ARISTOS/spiderweb/middleware"
	"ARISTOS/spiderweb/handlers"
	"ARISTOS/dblayer/common/asyncq"
)

func Run(){
	aysncq.StartTaskDispatcher(9)

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/login", handlers.LoginHandler(&env)).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")
	r.HandleFunc("/signup", handlers.SignUpHandler).Methods("GET", "POST")

	r.HandleFunc("/editor", middleware.GatedContentHandler(handlers.EditorHandler)).Methods("GET")
	r.HandleFunc("/editor/image-upload", middleware.GatedContentHandler(handlers.UploadImageHandler)).Methods("GET", "POST")
	r.HandleFunc("/editor/video-upload", middleware.GatedContentHandler(handlers.UploadVideoHandler)).Methods("GET", "POST")

	// middleware stuffs
	http.Handle("/", middleware.PanicRecoveryHandler(ghandlers.LoggingHandler(os.Stdout, r)))

	// basic server stuffs
	http.ListenAndServe(":8080", nil)
}
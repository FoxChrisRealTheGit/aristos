package spiderweb

import(
	"net/http"
	"os"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"ARISTOS/spiderweb/middleware"
	"ARISTOS/spiderweb/handlers"
	"ARISTOS/dblayer/common/asyncq"
	"ARISTOS/dblayer/common"
)

func Run(){
	asyncq.StartTaskDispatcher(9)

	env := common.Env{}

	r := mux.NewRouter()
	r.Handle("/", handlers.LoginHandler(&env)).Methods("GET", "POST")
	r.Handle("/login", handlers.LoginHandler(&env)).Methods("GET","POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET","POST")
	r.Handle("/signup", handlers.SignUpHandler(&env)).Methods("GET", "POST")

	// r.Handle("/editor", middleware.GatedContentHandler(handlers.EditorHandler)).Methods("GET")
	r.Handle("/editor/image-upload", middleware.GatedContentHandler(handlers.UploadImageHandler)).Methods("GET", "POST")
	r.Handle("/editor/video-upload", middleware.GatedContentHandler(handlers.UploadVideoHandler)).Methods("GET", "POST")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))


	loggedRouter := ghandlers.LoggingHandler(os.Stdout, r)
	// middleware stuffs
	http.Handle("/", middleware.PanicRecoveryHandler(loggedRouter))

	// basic server stuffs
	http.ListenAndServe(":8080", nil)
}
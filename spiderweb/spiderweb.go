package spiderweb

import(
	"net/http"
	"os"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"ARISTOS/middleware"
	"ARISTOS/spiderweb/handlers"
)

func Run(){
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)

	
	//middleware stuffs
	http.Handle("/", middleware.PanicRecoveryHandler(ghandlers.LoggingHandler(os.Stdout, r)))

	http.ListenAndServe(":8080", nil)
}
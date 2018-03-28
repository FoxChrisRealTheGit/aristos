package spiderweb

import(
	"net/http"
	"fmt"
)

func Run(){
	http.HandleFunc("/", sroot)
	http.Handle("/testhandle", newHandler())
	http.HandleFunc("/testquery", queryTestHandler)
	http.ListenAndServe(":8080", nil)
}

func queryTestHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	q := r.URL.Query()
	message := fmt.Sprintf("Query map: %v \n", q)

	v1, v2 := q.Get("key1"), q.Get("key2")
	if v1 == v2{
		message = message + fmt.Sprintf("V1 and V2 are equal %s \n")
	} else {
		message = message + fmt.Sprintf("V1is equal %s, V2 isequal %s \n", v1, v2)
	}
	fmt.Fprint(w, message)
}


func sroot(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Welcome to theAristos website builder")
}
package handlers

import (
	"net/http"
	"ARISTOS/dblayer/common/authenticate"
)
func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	authenticate.ExpireUserSession(w, r)
	authenticate.ExpireSecureCookie(w, r)
}

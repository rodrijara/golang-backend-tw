package midware

import (
	"net/http"
	"github.com/rodri-jara/golang-app-tw/db"
)

// CheckDB raises an error if DB not found
func CheckDatabase(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if db.CheckConnection() == 0{
			http.Error(w, "Lost connection", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
} 
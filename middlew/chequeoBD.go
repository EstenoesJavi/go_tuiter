package middlew

import (
	"net/http"

	"github.com/EstenoesJavi/go_tuiter/bd"
)

/*ChequeoBD es el middlew que me permite conocer el estado de la bd*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}

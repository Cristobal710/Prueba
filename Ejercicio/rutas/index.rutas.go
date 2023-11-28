package rutas

import "net/http"

func ManejoHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola mundo pedro"))
}

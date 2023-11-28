package rutas

import "net/http"

func ObtenerUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Users"))
}

func ObtenerUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User"))
}

func CrearUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Post"))
}

func EliminarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}

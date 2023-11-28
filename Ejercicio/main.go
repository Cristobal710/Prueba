package main

import (
	"net/http"

	"github.com/Cristobal710/Ejercicio/DataBase"
	"github.com/Cristobal710/Ejercicio/modelos"
	"github.com/Cristobal710/Ejercicio/rutas"
	"github.com/gorilla/mux"
)

func main() {
	DataBase.Conexion()

	DataBase.DB.AutoMigrate(modelos.Tarea{})
	DataBase.DB.AutoMigrate(modelos.Usuario{})

	r := mux.NewRouter()

	r.HandleFunc("/", rutas.ManejoHome)

	r.HandleFunc("/users", rutas.ObtenerUsuarios).Methods("GET")
	r.HandleFunc("/users/{id}", rutas.ObtenerUsuario).Methods("GET")
	r.HandleFunc("/users", rutas.CrearUsuario).Methods("POST")
	r.HandleFunc("/users", rutas.EliminarUsuario).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}

package modelos

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model

	Nombre   string `gorm:"not null"`
	Apellido string `gorm:"not null"`
	Email    string `gorm:"not null;unique_index"`
	Edad     int    `gorm:"not null"`
	//Tareas   []Tarea
}

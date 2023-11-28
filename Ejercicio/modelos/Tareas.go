package modelos

import "gorm.io/gorm"

type Tarea struct {
	gorm.Model

	Titulo      string `gorm:"not null;unique_index"`
	Descripcion string
	Terminado   bool `gorm:"default:false"`
	UserID      int
}

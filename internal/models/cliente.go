package models

import "gorm.io/gorm"

// TAREA (CP1): Complete los campos de Cliente según lo que muestran las pantallas.
//
// Pistas de trabajo:
//   - Guíese por el modelo Funcion para los tags gorm y json.
//   - Los tests de acceptance/ compilan contra los nombres EXACTOS de los campos.
//     Mientras falten campos, `go test ./acceptance/...` mostrará errores de
//     compilación que le indican qué está faltando.

type Cliente struct {
	gorm.Model

	Nombre   string `json:"nombre" gorm:"type:varchar(100);not null"`
	Cedula   string `json:"cedula" gorm:"type:varchar(20);unique;not null"`
	Telefono string `json:"telefono" gorm:"type:varchar(20)"`
}

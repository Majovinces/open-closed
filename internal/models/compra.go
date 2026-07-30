package models

import "gorm.io/gorm"

// TAREA (CP1): Complete los campos de Compra según lo que muestran las pantallas.
//
// Pistas de trabajo:
//   - Un Compra referencia a una Funcion y a un Cliente (claves foráneas).
//   - Recuerde el campo de estado (use las constantes de estados.go) y el total.
//   - Los tests de acceptance/ compilan contra los nombres EXACTOS de los campos.
type Compra struct {
	gorm.Model
	// TODO: agregue aquí los campos.

	ID        uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	ClienteID uint    `json:"cliente_id" gorm:"not null"`
	Cliente   Cliente `json:"cliente,omitempty" gorm:"foreignKey:ClienteID"`
	FuncionID uint    `json:"funcion_id" gorm:"not null"`
	Funcion   Funcion `json:"funcion,omitempty" gorm:"foreignKey:FuncionID"`
	Cantidad  uint    `json:"cantidad" gorm:"not null"` // <-- Cambiado de CantidadEntradas int a Cantidad uint
	Subtotal  float64 `json:"subtotal" gorm:"type:decimal(10,2);not null"`
	Descuento float64 `json:"descuento" gorm:"type:decimal(10,2);default:0"`
	Total     float64 `json:"total" gorm:"type:decimal(10,2);not null"`
	Estado    string  `json:"estado" gorm:"type:varchar(30);default:'COMPLETADA'"`
}

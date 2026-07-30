package storage

import (
	"errors"

	"gorm.io/gorm"

	"github.com/joancema/examen-cine/internal/models"
)

// TAREA (CP1): Implemente ClienteGORM contra la interfaz ClienteRepository.
//
// Reglas:
//   - NO cambie el nombre del tipo, del constructor ni las firmas de los métodos:
//     los tests de acceptance/ compilan contra ellos.
//   - Guíese por FuncionGORM (funcion_gorm.go): es el mismo patrón.
type ClienteGORM struct {
	db *gorm.DB
}

func NuevoClienteGORM(db *gorm.DB) *ClienteGORM {
	return &ClienteGORM{db: db}
}

func (r *ClienteGORM) Crear(c *models.Cliente) error {
	return r.db.Create(c).Error
}

func (r *ClienteGORM) ObtenerPorID(id uint) (models.Cliente, bool) {
	var cliente models.Cliente
	err := r.db.First(&cliente, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Cliente{}, false
		}
		return models.Cliente{}, false
	}
	return cliente, true
}
func (r *ClienteGORM) Listar() ([]models.Cliente, error) {
	var clientes []models.Cliente
	err := r.db.Find(&clientes).Error
	if err != nil {
		return nil, err
	}
	return clientes, nil
}

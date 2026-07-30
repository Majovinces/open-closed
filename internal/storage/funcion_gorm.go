// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import (
	"gorm.io/gorm"

	"github.com/joancema/examen-cine/internal/models"
)

// FuncionGORM implementa FuncionRepository sobre GORM.
// Esta implementación está completa: úsela como plantilla para ClienteGORM
// y CompraGORM, que usted debe implementar.
type FuncionGORM struct {
	db *gorm.DB
}

func NuevaFuncionGORM(db *gorm.DB) *FuncionGORM {
	return &FuncionGORM{db: db}
}

func (r *FuncionGORM) Crear(h *models.Funcion) error {
	return r.db.Create(h).Error
}

func (r *FuncionGORM) ObtenerPorID(id uint) (models.Funcion, bool) {
	var h models.Funcion
	if err := r.db.First(&h, id).Error; err != nil {
		return models.Funcion{}, false
	}
	return h, true
}

func (r *FuncionGORM) Listar() ([]models.Funcion, error) {
	var lista []models.Funcion
	err := r.db.Find(&lista).Error
	return lista, err
}

func (r *FuncionGORM) Actualizar(h *models.Funcion) error {
	return r.db.Save(h).Error
}

package storage

import (
	"gorm.io/gorm"

	"github.com/joancema/examen-cine/internal/models"
)

// TAREA (CP2): Implemente CompraGORM contra la interfaz CompraRepository.
//
// Reglas:
//   - NO cambie el nombre del tipo, del constructor ni las firmas de los métodos.
//   - Guíese por FuncionGORM: es el mismo patrón con una entidad distinta.
//   - Recuerde: aquí NO va lógica de negocio. Solo persistencia.
type CompraGORM struct {
	db *gorm.DB
}

// CORREGIDO: Debe retornar la interfaz CompraRepository (o permitir la asignación implícita)
func NuevaCompraGORM(db *gorm.DB) CompraRepository {
	return &CompraGORM{db: db}
}

func (r *CompraGORM) Crear(compra *models.Compra) error {
	return r.db.Create(compra).Error
}

// CORREGIDO: Ajustado a la firma exacta de la interfaz: (models.Compra, bool)
func (r *CompraGORM) ObtenerPorID(id uint) (models.Compra, bool) {
	var compra models.Compra
	err := r.db.Preload("Cliente").Preload("Funcion").First(&compra, id).Error
	if err != nil {
		return models.Compra{}, false
	}
	return compra, true
}

func (r *CompraGORM) Listar() ([]models.Compra, error) {
	var compras []models.Compra
	err := r.db.Preload("Cliente").Preload("Funcion").Find(&compras).Error
	if err != nil {
		return nil, err
	}
	return compras, nil
}

func (r *CompraGORM) Actualizar(compra *models.Compra) error {
	return r.db.Save(compra).Error
}

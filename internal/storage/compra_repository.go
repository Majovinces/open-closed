// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import "github.com/joancema/examen-cine/internal/models"

// CompraRepository define el contrato de persistencia de Compra.
// Su implementación GORM (en compra_gorm.go) debe satisfacer EXACTAMENTE
// estas firmas. Observe que el repositorio NO contiene lógica de negocio:
// las reglas (validaciones, cálculo del total, anulación) viven en el service.
type CompraRepository interface {
	Crear(a *models.Compra) error
	ObtenerPorID(id uint) (models.Compra, bool)
	Listar() ([]models.Compra, error)
	Actualizar(a *models.Compra) error
}

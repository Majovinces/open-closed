// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import "github.com/joancema/examen-cine/internal/models"

// FuncionRepository define el contrato de persistencia de la Entidad A.
type FuncionRepository interface {
	Crear(h *models.Funcion) error
	ObtenerPorID(id uint) (models.Funcion, bool)
	Listar() ([]models.Funcion, error)
	Actualizar(h *models.Funcion) error
}

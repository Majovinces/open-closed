// ARCHIVO BLOQUEADO — NO MODIFICAR
package services

import (
	"github.com/joancema/examen-cine/internal/models"
	"github.com/joancema/examen-cine/internal/storage"
)

// FuncionService contiene la lógica de negocio de la Entidad A.
// Está completo: úselo como ejemplo de cómo un service valida datos,
// devuelve errores de dominio y delega la persistencia al repository.
type FuncionService struct {
	repo storage.FuncionRepository
}

func NuevaFuncionService(repo storage.FuncionRepository) *FuncionService {
	return &FuncionService{repo: repo}
}

func (s *FuncionService) Crear(h *models.Funcion) error {
	if h.Nombre == "" || h.PrecioUnitario <= 0 {
		return ErrDatosInvalidos
	}
	return s.repo.Crear(h)
}

func (s *FuncionService) ObtenerPorID(id uint) (models.Funcion, error) {
	h, ok := s.repo.ObtenerPorID(id)
	if !ok {
		return models.Funcion{}, ErrNoEncontrado
	}
	return h, nil
}

func (s *FuncionService) Listar() ([]models.Funcion, error) {
	return s.repo.Listar()
}

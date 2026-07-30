package services

import (
	"errors"

	"github.com/joancema/examen-cine/internal/models"
	"github.com/joancema/examen-cine/internal/storage"
)

// TAREA (CP1): Implemente ClienteService.
//
// Reglas:
//   - NO cambie el nombre del tipo, del constructor ni las firmas de los métodos.
//   - Cliente no tiene reglas de negocio complejas: valide lo evidente según
//     las pantallas (campos obligatorios -> ErrDatosInvalidos) y delegue al
//     repository. Guíese por FuncionService.
type ClienteService struct {
	repo storage.ClienteRepository
}

func NuevoClienteService(repo storage.ClienteRepository) *ClienteService {
	return &ClienteService{repo: repo}
}

func (s *ClienteService) Crear(c *models.Cliente) error {
	// Validación básica de campos obligatorios
	if c.Nombre == "" || c.Cedula == "" {
		return errors.New("datos inválidos")
	}
	return s.repo.Crear(c)
}

func (s *ClienteService) ObtenerPorID(id uint) (models.Cliente, error) {
	cliente, ok := s.repo.ObtenerPorID(id)
	if !ok {
		return models.Cliente{}, errors.New("no encontrado")
	}
	return cliente, nil
}

func (s *ClienteService) Listar() ([]models.Cliente, error) {
	return s.repo.Listar()
}

// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import (
	"sync"

	"github.com/joancema/examen-cine/internal/models"
)

// FuncionMemoria implementa FuncionRepository en memoria.
// Se usa en los tests de reglas de negocio como fake del repositorio real.
type FuncionMemoria struct {
	mu     sync.Mutex
	datos  map[uint]models.Funcion
	nextID uint
}

func NuevaFuncionMemoria() *FuncionMemoria {
	return &FuncionMemoria{datos: make(map[uint]models.Funcion), nextID: 1}
}

func (r *FuncionMemoria) Crear(h *models.Funcion) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	h.ID = r.nextID
	r.nextID++
	r.datos[h.ID] = *h
	return nil
}

func (r *FuncionMemoria) ObtenerPorID(id uint) (models.Funcion, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	h, ok := r.datos[id]
	return h, ok
}

func (r *FuncionMemoria) Listar() ([]models.Funcion, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	lista := make([]models.Funcion, 0, len(r.datos))
	for _, h := range r.datos {
		lista = append(lista, h)
	}
	return lista, nil
}

func (r *FuncionMemoria) Actualizar(h *models.Funcion) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.datos[h.ID]; !ok {
		return ErrRegistroNoExiste
	}
	r.datos[h.ID] = *h
	return nil
}

// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import (
	"sync"

	"github.com/joancema/examen-cine/internal/models"
)

// CompraMemoria implementa CompraRepository en memoria.
// Se usa en los tests de reglas de negocio como fake del repositorio real.
type CompraMemoria struct {
	mu     sync.Mutex
	datos  map[uint]models.Compra
	nextID uint
}

func NuevaCompraMemoria() *CompraMemoria {
	return &CompraMemoria{datos: make(map[uint]models.Compra), nextID: 1}
}

func (r *CompraMemoria) Crear(a *models.Compra) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	a.ID = r.nextID
	r.nextID++
	r.datos[a.ID] = *a
	return nil
}

func (r *CompraMemoria) ObtenerPorID(id uint) (models.Compra, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	a, ok := r.datos[id]
	return a, ok
}

func (r *CompraMemoria) Listar() ([]models.Compra, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	lista := make([]models.Compra, 0, len(r.datos))
	for _, a := range r.datos {
		lista = append(lista, a)
	}
	return lista, nil
}

func (r *CompraMemoria) Actualizar(a *models.Compra) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.datos[a.ID]; !ok {
		return ErrRegistroNoExiste
	}
	r.datos[a.ID] = *a
	return nil
}

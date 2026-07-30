// ARCHIVO BLOQUEADO — NO MODIFICAR
package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NuevoRouter registra todas las rutas de la API. Este archivo es el
// contrato HTTP del examen: los tests httptest de acceptance/ atacan
// exactamente estas rutas.
func NuevoRouter(
	funciones *FuncionHandler,
	clientes *ClienteHandler,
	compras *CompraHandler,
) http.Handler {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/funciones", func(r chi.Router) {
			r.Get("/", funciones.Listar)
			r.Post("/", funciones.Crear)
		})

		r.Route("/clientes", func(r chi.Router) {
			r.Get("/", clientes.Listar)
			r.Post("/", clientes.Crear)
			r.Get("/{id}", clientes.ObtenerPorID)
		})

		r.Route("/compras", func(r chi.Router) {
			r.Get("/", compras.Listar)
			r.Post("/", compras.Crear)
			r.Get("/{id}", compras.ObtenerPorID)
			r.Post("/{id}/cancelar", compras.Cancelar)
		})
	})

	return r
}

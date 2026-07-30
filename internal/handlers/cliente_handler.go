package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/joancema/examen-cine/internal/models"
	"github.com/joancema/examen-cine/internal/services"
)

// TAREA (CP1): Implemente ClienteHandler.
// Reglas:
//   - NO cambie el nombre del tipo, del constructor ni las firmas de los métodos:
//     routes.go (bloqueado) los registra y los tests httptest los atacan.
//   - Guíese por FuncionHandler para decodificar JSON y mapear errores:
//     ErrDatosInvalidos -> 422, ErrNoEncontrado -> 404.
//   - Para leer el {id} de la ruta: chi.URLParam(r, "id") y strconv.

type ClienteHandler struct {
	servicio *services.ClienteService
}

func NuevoClienteHandler(s *services.ClienteService) *ClienteHandler {
	return &ClienteHandler{servicio: s}
}

func (h *ClienteHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var cliente models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&cliente); err != nil {
		RespondError(w, http.StatusUnprocessableEntity, "Datos inválidos")
		return
	}

	err := h.servicio.Crear(&cliente)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	RespondJSON(w, http.StatusCreated, cliente)
}

func (h *ClienteHandler) Listar(w http.ResponseWriter, r *http.Request) {
	clientes, err := h.servicio.Listar()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, clientes)
}

func (h *ClienteHandler) ObtenerPorID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	cliente, err := h.servicio.ObtenerPorID(uint(id))
	if err != nil {
		RespondError(w, http.StatusNotFound, "Cliente no encontrado")
		return
	}

	RespondJSON(w, http.StatusOK, cliente)
}

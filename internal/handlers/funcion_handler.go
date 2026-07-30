// ARCHIVO BLOQUEADO — NO MODIFICAR
package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/joancema/examen-cine/internal/models"
	"github.com/joancema/examen-cine/internal/services"
)

// FuncionHandler expone la Entidad A por HTTP.
// Está completo: observe cómo decodifica el body, llama al service y
// MAPEA los errores de dominio a status codes. Ese mapeo es exactamente
// lo que usted debe replicar en sus propios handlers.
type FuncionHandler struct {
	servicio *services.FuncionService
}

func NuevoFuncionHandler(s *services.FuncionService) *FuncionHandler {
	return &FuncionHandler{servicio: s}
}

func (h *FuncionHandler) Listar(w http.ResponseWriter, r *http.Request) {
	lista, err := h.servicio.Listar()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, lista)
}

func (h *FuncionHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var funcion models.Funcion
	if err := json.NewDecoder(r.Body).Decode(&funcion); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido")
		return
	}
	if err := h.servicio.Crear(&funcion); err != nil {
		switch {
		case errors.Is(err, services.ErrDatosInvalidos):
			RespondError(w, http.StatusUnprocessableEntity, err.Error())
		default:
			RespondError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	RespondJSON(w, http.StatusCreated, funcion)
}

package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/joancema/examen-cine/internal/models"
	"github.com/joancema/examen-cine/internal/services"
)

type CompraHandler struct {
	servicio *services.CompraService
}

func NuevaCompraHandler(s *services.CompraService) *CompraHandler {
	return &CompraHandler{servicio: s}
}

func (h *CompraHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var compra models.Compra
	if err := json.NewDecoder(r.Body).Decode(&compra); err != nil {
		RespondError(w, http.StatusBadRequest, "cuerpo de solicitud inválido")
		return
	}

	if err := h.servicio.Crear(&compra); err != nil {
		h.manejarError(w, err)
		return
	}

	RespondJSON(w, http.StatusCreated, compra)
}

func (h *CompraHandler) Listar(w http.ResponseWriter, r *http.Request) {
	compras, err := h.servicio.Listar()
	if err != nil {
		h.manejarError(w, err)
		return
	}
	RespondJSON(w, http.StatusOK, compras)
}

func (h *CompraHandler) ObtenerPorID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	compra, err := h.servicio.ObtenerPorID(uint(id))
	if err != nil {
		h.manejarError(w, err)
		return
	}

	RespondJSON(w, http.StatusOK, compra)
}

func (h *CompraHandler) Cancelar(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	if err := h.servicio.Cancelar(uint(id)); err != nil {
		h.manejarError(w, err)
		return
	}

	RespondJSON(w, http.StatusOK, map[string]string{"mensaje": "compra cancelada exitosamente"})
}

// manejarError traduce los errores de dominio a los códigos HTTP requeridos por el examen
func (h *CompraHandler) manejarError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, services.ErrDatosInvalidos), errors.Is(err, services.ErrReferenciaInvalida):
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
	case errors.Is(err, services.ErrStockInsuficiente), errors.Is(err, services.ErrEstadoInvalido):
		RespondError(w, http.StatusConflict, err.Error())
	case errors.Is(err, services.ErrNoEncontrado):
		RespondError(w, http.StatusNotFound, err.Error())
	default:
		RespondError(w, http.StatusInternalServerError, err.Error())
	}
}

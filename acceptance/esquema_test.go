// ARCHIVO BLOQUEADO — NO MODIFICAR
package acceptance

import (
	"testing"

	"github.com/joancema/examen-cine/internal/models"
)

// TestCP1_EsquemaMigra verifica que los modelos completados por el estudiante
// generan las tablas y columnas esperadas (nombres de columna en snake_case,
// como los produce GORM por convención).
func TestCP1_EsquemaMigra(t *testing.T) {
	db := nuevaDB(t)
	m := db.Migrator()

	columnasCliente := []string{"nombre", "cedula", "telefono"}
	for _, col := range columnasCliente {
		if !m.HasColumn(&models.Cliente{}, col) {
			t.Errorf("falta la columna %q en la tabla de Cliente (¿definió el campo en el modelo?)", col)
		}
	}

	columnasCompra := []string{"funcion_id", "cliente_id", "cantidad", "estado", "total"}
	for _, col := range columnasCompra {
		if !m.HasColumn(&models.Compra{}, col) {
			t.Errorf("falta la columna %q en la tabla de Compra (¿definió el campo en el modelo?)", col)
		}
	}
}

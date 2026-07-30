// ARCHIVO BLOQUEADO — NO MODIFICAR
//
// Las 5 reglas de negocio se verifican aquí usando los repositorios EN MEMORIA
// (ya implementados en el repo base) como fakes. Así, estos tests miden solo
// la lógica de su CompraService, sin depender de su implementación GORM.
package acceptance

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/joancema/examen-cine/internal/models"
	"github.com/joancema/examen-cine/internal/services"
	"github.com/joancema/examen-cine/internal/storage"
)

type entornoReglas struct {
	svc          *services.CompraService
	funciones *storage.FuncionMemoria
	clientes     *storage.ClienteMemoria
	compras   *storage.CompraMemoria
	principal      models.Funcion
	ana          models.Cliente
}

func nuevoEntornoReglas(t *testing.T) entornoReglas {
	t.Helper()
	hm := storage.NuevaFuncionMemoria()
	cm := storage.NuevoClienteMemoria()
	am := storage.NuevaCompraMemoria()

	principal := models.Funcion{Nombre: "Matiné familiar", PrecioUnitario: 8.5, Stock: 10, Activo: true}
	require.NoError(t, hm.Crear(&principal))
	ana := models.Cliente{Nombre: "Ana Zambrano", Cedula: "1310000001", Telefono: "0990000001"}
	require.NoError(t, cm.Crear(&ana))

	return entornoReglas{
		svc:          services.NuevaCompraService(am, hm, cm),
		funciones: hm,
		clientes:     cm,
		compras:   am,
		principal:      principal,
		ana:          ana,
	}
}

// R1: no se crea una compra si la funcion no existe o está inactiva,
// o si el cliente no existe.
func TestCP2_R1_ReferenciasValidas(t *testing.T) {
	e := nuevoEntornoReglas(t)

	a := models.Compra{FuncionID: 99999, ClienteID: e.ana.ID, Cantidad: 1}
	require.ErrorIs(t, e.svc.Crear(&a), services.ErrReferenciaInvalida,
		"crear con una funcion inexistente debe devolver ErrReferenciaInvalida")

	extra := models.Funcion{Nombre: "Función de medianoche", PrecioUnitario: 15, Stock: 3, Activo: false}
	require.NoError(t, e.funciones.Crear(&extra))
	a = models.Compra{FuncionID: extra.ID, ClienteID: e.ana.ID, Cantidad: 1}
	require.ErrorIs(t, e.svc.Crear(&a), services.ErrReferenciaInvalida,
		"crear con una funcion INACTIVA debe devolver ErrReferenciaInvalida")

	a = models.Compra{FuncionID: e.principal.ID, ClienteID: 99999, Cantidad: 1}
	require.ErrorIs(t, e.svc.Crear(&a), services.ErrReferenciaInvalida,
		"crear con un cliente inexistente debe devolver ErrReferenciaInvalida")
}

// R2: la cantidad no puede superar el stock disponible.
func TestCP2_R2_StockInsuficiente(t *testing.T) {
	e := nuevoEntornoReglas(t)

	a := models.Compra{FuncionID: e.principal.ID, ClienteID: e.ana.ID, Cantidad: 11}
	require.ErrorIs(t, e.svc.Crear(&a), services.ErrStockInsuficiente,
		"pedir 11 unidades con stock 10 debe devolver ErrStockInsuficiente")
}

// R3: Total = Cantidad x PrecioUnitario, con 10% de descuento desde 5 unidades.
func TestCP2_R3_CalculoTotal(t *testing.T) {
	e := nuevoEntornoReglas(t)

	sinDescuento := models.Compra{FuncionID: e.principal.ID, ClienteID: e.ana.ID, Cantidad: 3}
	require.NoError(t, e.svc.Crear(&sinDescuento),
		"crear una compra válida no debe devolver error")
	require.InDelta(t, 25.50, sinDescuento.Total, 0.001,
		"3 x 8.50 = 25.50 (sin descuento)")
	require.Equal(t, models.EstadoPendiente, sinDescuento.Estado,
		"una compra recién creada debe quedar en estado PENDIENTE")

	conDescuento := models.Compra{FuncionID: e.principal.ID, ClienteID: e.ana.ID, Cantidad: 5}
	require.NoError(t, e.svc.Crear(&conDescuento))
	require.InDelta(t, 38.25, conDescuento.Total, 0.001,
		"5 x 8.50 = 42.50, con 10% de descuento = 38.25")
}

// R4: solo se puede cancelar una compra en estado PENDIENTE.
func TestCP2_R4_CancelarSoloPendiente(t *testing.T) {
	e := nuevoEntornoReglas(t)

	usada := models.Compra{
		FuncionID: e.principal.ID,
		ClienteID:     e.ana.ID,
		Cantidad:      1,
		Estado:        models.EstadoUsada,
		Total:         8.5,
	}
	require.NoError(t, e.compras.Crear(&usada))
	require.ErrorIs(t, e.svc.Cancelar(usada.ID), services.ErrEstadoInvalido,
		"cancelar una compra USADA debe devolver ErrEstadoInvalido")

	require.ErrorIs(t, e.svc.Cancelar(99999), services.ErrNoEncontrado,
		"cancelar una compra inexistente debe devolver ErrNoEncontrado")
}

// R5: al crear se descuenta el stock; al cancelar, se repone.
func TestCP2_R5_ReposicionStock(t *testing.T) {
	e := nuevoEntornoReglas(t)

	a := models.Compra{FuncionID: e.principal.ID, ClienteID: e.ana.ID, Cantidad: 3}
	require.NoError(t, e.svc.Crear(&a))

	h, ok := e.funciones.ObtenerPorID(e.principal.ID)
	require.True(t, ok)
	require.Equal(t, uint(7), h.Stock,
		"al crear una compra de 3 unidades, el stock debe bajar de 10 a 7")

	require.NoError(t, e.svc.Cancelar(a.ID), "cancelar una compra PENDIENTE debe funcionar")

	cancelada, ok := e.compras.ObtenerPorID(a.ID)
	require.True(t, ok)
	require.Equal(t, models.EstadoCancelada, cancelada.Estado,
		"tras cancelar, la compra debe quedar en estado CANCELADA")

	h, ok = e.funciones.ObtenerPorID(e.principal.ID)
	require.True(t, ok)
	require.Equal(t, uint(10), h.Stock,
		"al cancelar, las 3 unidades deben reponerse al stock (7 -> 10)")
}

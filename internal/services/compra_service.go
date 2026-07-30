package services

import (
	"github.com/joancema/examen-cine/internal/models"
	"github.com/joancema/examen-cine/internal/storage"
)

// TAREA (CP2): Implemente CompraService con las 5 reglas de negocio.
//
// Las reglas están A LA VISTA en las pantallas (carpeta pantallas/) y los
// tests de acceptance/reglas_test.go las verifican una por una. Devuelva los
// errores de dominio de errores.go: los tests los comprueban con errors.Is.
//
// Reglas:
//   - NO cambie el nombre del tipo, del constructor ni las firmas de los métodos.
//   - Observe que el service recibe TRES repositories: necesita consultar
//     Funcion y Cliente para validar, y actualizar Funcion al cancelar.
type CompraService struct {
	compraRepo  storage.CompraRepository
	funcionRepo storage.FuncionRepository
	clienteRepo storage.ClienteRepository
}

func NuevaCompraService(
	compraRepo storage.CompraRepository,
	funcionRepo storage.FuncionRepository,
	clienteRepo storage.ClienteRepository,
) *CompraService {
	return &CompraService{
		compraRepo:  compraRepo,
		funcionRepo: funcionRepo,
		clienteRepo: clienteRepo,
	}
}

// Crear registra un nuevo compra aplicando R1, R2 y R3.
// TODO (R1): la funcion debe existir y estar activa; el cliente debe existir.
// TODO (R2): la cantidad no puede superar el stock disponible de la funcion.
// TODO (R3): calcule el total (observe en las pantallas cuándo aplica descuento).
// TODO: al crear, el stock de la funcion se descuenta (mire la pantalla 01
// antes y después de crear una compra; R5 es la operación inversa).
func (s *CompraService) Crear(a *models.Compra) error {
	if a.Cantidad == 0 || a.FuncionID == 0 || a.ClienteID == 0 {
		return ErrDatosInvalidos
	}

	funcion, ok := s.funcionRepo.ObtenerPorID(a.FuncionID)
	if !ok || !funcion.Activo {
		return ErrReferenciaInvalida
	}

	_, ok = s.clienteRepo.ObtenerPorID(a.ClienteID)
	if !ok {
		return ErrReferenciaInvalida
	}

	if a.Cantidad > funcion.Stock {
		return ErrStockInsuficiente
	}

	total := funcion.PrecioUnitario * float64(a.Cantidad)
	if a.Cantidad >= 5 {
		total = total * 0.90
	}
	a.Total = total
	a.Estado = models.EstadoPendiente

	// CORRECCIÓN: Restar el stock antes de actualizar en el repositorio
	funcion.Stock -= a.Cantidad
	if err := s.funcionRepo.Actualizar(&funcion); err != nil {
		return err
	}

	return s.compraRepo.Crear(a)
}

func (s *CompraService) ObtenerPorID(id uint) (models.Compra, error) {
	compra, ok := s.compraRepo.ObtenerPorID(id)
	if !ok {
		return models.Compra{}, ErrNoEncontrado
	}
	return compra, nil
}

func (s *CompraService) Listar() ([]models.Compra, error) {
	compras, err := s.compraRepo.Listar()
	if err != nil {
		return nil, err
	}
	return compras, nil
}

// Cancelar cancela una compra aplicando R4 y R5.
// TODO (R4): solo se puede cancelar una compra en estado PENDIENTE.
// TODO (R5): al cancelar, la cantidad se repone al stock de la funcion.
func (s *CompraService) Cancelar(id uint) error {
	compra, ok := s.compraRepo.ObtenerPorID(id)
	if !ok {
		return ErrNoEncontrado
	}

	if compra.Estado != models.EstadoPendiente {
		return ErrEstadoInvalido
	}

	compra.Estado = models.EstadoCancelada
	if err := s.compraRepo.Actualizar(&compra); err != nil {
		return err
	}

	funcion, ok := s.funcionRepo.ObtenerPorID(compra.FuncionID)
	if ok {
		funcion.Stock += compra.Cantidad
		_ = s.funcionRepo.Actualizar(&funcion)
	}

	return nil
}

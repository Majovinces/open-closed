// ARCHIVO BLOQUEADO — NO MODIFICAR
package main

import (
	"log"
	"net/http"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/joancema/examen-cine/internal/handlers"
	"github.com/joancema/examen-cine/internal/models"
	"github.com/joancema/examen-cine/internal/services"
	"github.com/joancema/examen-cine/internal/storage"
)

func main() {
	db, err := gorm.Open(sqlite.Open("cine.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("no se pudo abrir la base de datos: %v", err)
	}

	if err := db.AutoMigrate(
		&models.Funcion{},
		&models.Cliente{},
		&models.Compra{},
	); err != nil {
		log.Fatalf("error en la migración: %v", err)
	}

	sembrarFunciones(db)

	// Repositories (GORM)
	funcionRepo := storage.NuevaFuncionGORM(db)
	clienteRepo := storage.NuevoClienteGORM(db)
	compraRepo := storage.NuevaCompraGORM(db)

	// Services
	funcionSvc := services.NuevaFuncionService(funcionRepo)
	clienteSvc := services.NuevoClienteService(clienteRepo)
	compraSvc := services.NuevaCompraService(compraRepo, funcionRepo, clienteRepo)

	// Handlers + Router
	router := handlers.NuevoRouter(
		handlers.NuevoFuncionHandler(funcionSvc),
		handlers.NuevoClienteHandler(clienteSvc),
		handlers.NuevaCompraHandler(compraSvc),
	)

	log.Println("API del cine escuchando en http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

// sembrarFunciones carga el catálogo inicial solo si la tabla está vacía.
// Los clientes y compras se crean vía API.
func sembrarFunciones(db *gorm.DB) {
	var total int64
	db.Model(&models.Funcion{}).Count(&total)
	if total > 0 {
		return
	}
	iniciales := []models.Funcion{
		{Nombre: "Matiné familiar", PrecioUnitario: 8.50, Stock: 10, Activo: true},
		{Nombre: "Estreno 3D", PrecioUnitario: 6.00, Stock: 4, Activo: true},
		{Nombre: "Función clásica", PrecioUnitario: 5.00, Stock: 2, Activo: true},
		{Nombre: "Función de medianoche", PrecioUnitario: 15.00, Stock: 3, Activo: false},
	}
	for i := range iniciales {
		db.Create(&iniciales[i])
	}
}

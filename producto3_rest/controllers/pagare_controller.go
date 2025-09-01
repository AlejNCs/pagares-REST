package controllers

import (
	"database/sql"
	"producto3_rest/db"
	"producto3_rest/models"
	"producto3_rest/utils"
    "fmt"
    "net/http"
    "os"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

func CrearPagare(c *fiber.Ctx) error {
	var p models.Pagare
	if err := c.BodyParser(&p); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	query := `INSERT INTO pagares (monto, meses, status) VALUES ($1, $2, $3) RETURNING id`
	err := db.DB.QueryRow(query, p.Monto, p.Meses, p.Status).Scan(&p.ID)
	if err != nil {
		return c.Status(500).SendString("Error al insertar el pagaré")
	}

	return c.JSON(p)
}

func RegistrarPagare(c *fiber.Ctx) error {
	var p models.Pagare
	if err := c.BodyParser(&p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	p.Status = "pendiente"

	err := db.DB.QueryRow("INSERT INTO pagares (monto, meses, status) VALUES ($1, $2, $3) RETURNING id", p.Monto, p.Meses, p.Status).Scan(&p.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al insertar"})
	}

	// ✅ Generar PDF
	if err := utils.GenerarPDF(p); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Pagaré creado pero no se pudo generar el PDF"})
	}

	return c.JSON(p)
}

func ConsultarPagare(c *fiber.Ctx) error {
	id := c.Params("id")
	row := db.DB.QueryRow("SELECT id, monto, meses, status FROM pagares WHERE id=$1", id)

	var p models.Pagare
	err := row.Scan(&p.ID, &p.Monto, &p.Meses, &p.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{"error": "Pagaré no encontrado"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Error al consultar"})
	}

	return c.JSON(p)
}

func ModificarPagare(c *fiber.Ctx) error {
	id := c.Params("id")
	var p models.Pagare
	if err := c.BodyParser(&p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	_, err := db.DB.Exec("UPDATE pagares SET monto=$1, meses=$2, status=$3 WHERE id=$4", p.Monto, p.Meses, p.Status, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al actualizar"})
	}

	//ID = uint(parseUintOrZero(id)) // ✅ Asignar el ID manualmente

	// Regenerar PDF
	if err := utils.GenerarPDF(p); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Pagaré actualizado pero no se pudo generar el PDF"})
	}

	return c.JSON(fiber.Map{"mensaje": "Pagaré actualizado"})
}

func BorrarPagare(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := db.DB.Exec("DELETE FROM pagares WHERE id=$1", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al eliminar"})
	}

	return c.JSON(fiber.Map{"mensaje": "Pagaré eliminado"})
}

// Función auxiliar para convertir string a uint
func parseUintOrZero(s string) uint {
	var id uint
	fmt.Sscanf(s, "%d", &id)
	return id
}

func DescargarPDF(c *fiber.Ctx) error {
    idStr := c.Params("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).SendString("ID inválido")
    }

    filePath := fmt.Sprintf("pagare_%d.pdf", id)
    
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        return c.Status(http.StatusNotFound).SendString("El archivo PDF no existe")
    }

    return c.Download(filePath)
}



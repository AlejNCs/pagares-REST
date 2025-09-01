package utils

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"producto3_rest/models"
	//trconv"
)



func GenerarPDF(p models.Pagare) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "PAGARE")
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("ID: %d", p.ID))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Monto: $%.2f", p.Monto))
	pdf.Ln(8)

	// Conversión a texto
	montoLetra := NumeroALetras(int(p.Monto))
	pdf.Cell(40, 10, fmt.Sprintf("Monto en letra: %s pesos", montoLetra))
	pdf.Ln(8)

	pdf.Cell(40, 10, fmt.Sprintf("Meses: %d", p.Meses))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Status: %s", p.Status))

	fileName := fmt.Sprintf("pagare_%d.pdf", p.ID)
	return pdf.OutputFileAndClose(fileName)
}
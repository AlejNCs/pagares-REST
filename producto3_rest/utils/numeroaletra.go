package utils


var unidades = []string{
	"", "uno", "dos", "tres", "cuatro", "cinco",
	"seis", "siete", "ocho", "nueve", "diez",
	"once", "doce", "trece", "catorce", "quince",
	"dieciséis", "diecisiete", "dieciocho", "diecinueve", "veinte",
}

var decenas = []string{
	"", "", "veinte", "treinta", "cuarenta", "cincuenta",
	"sesenta", "setenta", "ochenta", "noventa",
}

var centenas = []string{
	"", "cien", "doscientos", "trescientos", "cuatrocientos",
	"quinientos", "seiscientos", "setecientos", "ochocientos", "novecientos",
}

func NumeroALetras(n int) string {
	if n == 0 {
		return "cero"
	}
	if n < 21 {
		return unidades[n]
	}
	if n < 100 {
		d := n / 10
		u := n % 10
		if u == 0 {
			return decenas[d]
		}
		return decenas[d] + " y " + unidades[u]
	}
	if n < 1000 {
		c := n / 100
		r := n % 100
		if c == 1 && r == 0 {
			return "cien"
		}
		return centenas[c] + " " + NumeroALetras(r)
	}
	if n < 1000000 {
		miles := n / 1000
		resto := n % 1000
		var milesTexto string
		if miles == 1 {
			milesTexto = "mil"
		} else {
			milesTexto = NumeroALetras(miles) + " mil"
		}
		if resto == 0 {
			return milesTexto
		}
		return milesTexto + " " + NumeroALetras(resto)
	}

	return "número demasiado grande"
}

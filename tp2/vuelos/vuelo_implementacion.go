package vuelos

import (
	"strconv"
	"strings"
)

func convertirAInt(prioridad, demora string) (int, int) {
	nPrioridad, _ := strconv.Atoi(prioridad)
	nDemora, _ := strconv.Atoi(demora)
	return nPrioridad, nDemora
}

func CrearVuelo(infoDeVuelo string) Vuelo {
	informacion := strings.Split(infoDeVuelo, ",")

	/* ------------------------------ PROBAR CUANDO LA PRUEBA DE VOLUMEN DE OK ---------------------------------------- */
	prioridad, demora := convertirAInt(informacion[PRIORIDAD], informacion[DEMORA])
	informacion[PRIORIDAD], informacion[DEMORA] = strconv.Itoa(prioridad), strconv.Itoa(demora)

	campos := CamposComparables{Codigo: Codigo(informacion[CODIGO]), Fecha: informacion[FECHA]}

	boleto := Vuelo{
		InfoComparable:      campos,
		Origen:              informacion[ORIGEN],
		Destino:             informacion[DESTINO],
		Prioridad:           prioridad,
		DemoraDeDespegue:    demora,
		InformacionCompleta: strings.Join(informacion, " "),
	}
	return boleto
}

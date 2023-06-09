package vuelos

import (
	"strconv"
	"strings"
)

func convertirAInt(prioridad, demora, tiempo, cancelado string) (int, int, int, int) {
	nPrioridad, _ := strconv.Atoi(prioridad)
	nDemora, _ := strconv.Atoi(demora)
	nTiempo, _ := strconv.Atoi(tiempo)
	nCancelado, _ := strconv.Atoi(cancelado)
	return nPrioridad, nDemora, nTiempo, nCancelado
}

func CrearVuelo(infoDeVuelo string) Vuelo {
	informacion := strings.Split(infoDeVuelo, ",")

	prioridad, demora, tiempo, cancelado := convertirAInt(informacion[PRIORIDAD], informacion[DEMORA], informacion[TIEMPO], informacion[CANCELADO])
	informacion[PRIORIDAD], informacion[DEMORA] = strconv.Itoa(prioridad), strconv.Itoa(demora)
	informacion[TIEMPO], informacion[CANCELADO] = strconv.Itoa(tiempo), strconv.Itoa(cancelado)

	campos := CamposComparables{Codigo: Codigo(informacion[CODIGO]), Prioridad: prioridad, Fecha: informacion[FECHA]}

	boleto := Vuelo{
		InfoComparable:      campos,
		Origen:              informacion[ORIGEN],
		Destino:             informacion[DESTINO],
		DemoraDeDespegue:    demora,
		TiempoDeVuelo:       tiempo,
		Cancelacion:         cancelado,
		InformacionCompleta: strings.Join(informacion, " "),
	}
	return boleto
}

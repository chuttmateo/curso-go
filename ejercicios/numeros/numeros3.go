package numeros

//Calculadora de propinas: Escribe un programa que calcule el monto de la propina en función
//del total de la cuenta y un porcentaje de propina dado.
func CalcularPropinas(monto, porcentaje float32) float32 {
	return monto * (porcentaje / 100)
}

//Conversor de unidades: Crea un programa que convierta una cantidad en una unidad de medida a otra.
//Por ejemplo, puedes convertir kilómetros a millas, grados Celsius a Fahrenheit, etc.
func ConversorDeUnidades(cantidad float32, unidadSalida string) float32 {
	if unidadSalida == "Fahrenheit" {
		return (cantidad * 9 / 5) + 32
	}
	if unidadSalida == "Millas" {
		return cantidad / 1.609
	}
	return 0
}

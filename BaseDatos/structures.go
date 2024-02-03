package BaseDatos

type GastoMensual struct {
	Id      int
	Idmes     int
	Fecha     string
	Importe   float64
	Concepto  string
	Tipocosto string
}

type Meses struct {
	Id int
	Mes string
}

type TotalPorMes struct {
	Id     int
	Idmes    int
	Totalmes float64
}

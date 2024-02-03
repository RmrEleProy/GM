package BaseDatos

const (
	Stmtablameses = `CREATE TABLE IF NOT EXISTS meses (
	id INTEGER PRIMARY KEY,
	mes TEXT NOT NULL);`

	StmTablaGastoMensual = `CREATE TABLE IF NOT EXISTS gasto_mensual (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		idmes INTEGER,
		fecha TEXT,
		importe REAL,
		concepto TEXT,
		tipocosto TEXT,
		FOREIGN KEY(idmes) REFERENCES meses(id)	);`

	StmTablaTotalPorMeses = `CREATE TABLE IF NOT EXISTS total_por_mes (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	idmes INTEGER,
	totalmes REAL,
	FOREIGN KEY(idmes) REFERENCES meses(id)	);`

	StmSelectMeses = "SELECT * FROM meses;"

	SelectMesByID = "SELECT mes FROM meses WHERE id=?"

	SelectGastoMensual = "SELECT * FROM gasto_mensual ORDER BY id DESC;"

	SelectGastoMensualById = "SELECT * FROM gasto_mensual WHERE id=?"

	SelectGastoMensualByIdmes = "SELECT * FROM gasto_mensual WHERE idmes=?"

	SelectSumImporGastoMensual = "SELECT SUM(IMPORTE) from gasto_mensual WHERE idmes=?"

	SelectTotalMes = "SELECT * FROM total_por_mes ORDER BY id DESC;"

	SelectTotalMesByIdmes = "SELECT * FROM total_por_mes WHERE idmes=?"

	UpdateTotalMesByIdmes = "UPDATE total_por_mes SET totalmes=? where idmes=?"

	InsertarEnGm = "INSERT INTO gasto_mensual(Idmes, Fecha, Importe, Concepto, Tipocosto) VALUES(?,?,?,?,?)"

	UpdateGM = "UPDATE gasto_mensual SET idmes=?, fecha=?, importe=?, concepto=?, tipocosto=? WHERE id=?"

	DirecionBaseDatosGAU = "./BaseDatos/AUGastosMensuales.db"
)

var EjecutadaFuncionDeInicio bool

//  

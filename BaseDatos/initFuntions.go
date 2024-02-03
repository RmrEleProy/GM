package BaseDatos

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// meses, gasto_mensual, total_por_mes
func Principal() {
	boolmeses, err := VerificaTablas("meses")
	if err != nil {
		log.Fatal(err)
	}
	boolgmensual, err := VerificaTablas("gasto_mensual")
	if err != nil {
		log.Fatal(err)
	}
	boolttmes, err := VerificaTablas("total_por_mes")
	if err != nil {
		log.Fatal(err)
	}

	if boolmeses == "FALSE" {
		CrearTablas(Stmtablameses)
		AlmacenarMeses()
	}
	if boolttmes == "FALSE" {
		CrearTablas(StmTablaGastoMensual)
	}
	if boolgmensual == "FALSE" {
		CrearTablas(StmTablaTotalPorMeses)
		ColocarValInitTTMes()
	}
}

// Funcion para almacenar el nombre de los meses
func AlmacenarMeses() {
	db := DBconn()
	defer db.Close()

	// Verificar si la tabla está vacía
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM meses").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	// Si la tabla está vacía, insertar los datos
	if count == 0 {
		meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}
		for i, mes := range meses {
			_, err = db.Exec("INSERT INTO meses (id, mes) VALUES (?, ?)", i+1, mes)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("Los datos han sido insertados en la tabla mes.")
	} else {
		fmt.Println("Los datos ya existen en la tabla mes.")
	}
}

// Esta funcion le coloca todos los meses el inicio a 0,
// en la base de datos total_por_mes
// para que se vaya actualizando a medida que se alimente la base de datos
func ColocarValInitTTMes() {
	db := DBconn()
	defer db.Close()

	// Verificar si la tabla está vacía
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM total_por_mes").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	// Si la tabla está vacía, insertar los datos
	if count == 0 {
		for i := 0; i < 12; i++ {
			_, err = db.Exec("INSERT INTO total_por_mes (idmes, totalmes) VALUES (?, ?)", i+1, 0)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("Los datos han sido insertados en la tabla total_por_mes.")
	} else {
		fmt.Println("Los datos ya existen en la tabla total_por_mes.")
	}
}

/*
mira si la tabla se ha creado: genera un string ya sea "FALSE" si no existe la tabla ó "TRUE" si existe la tabla,
las tablas para este caso son:
meses, gasto_mensual, total_por_mes
*/
func VerificaTablas(tabla string) (string, error) {
	db := DBconn()
	defer db.Close()
	// tabla = "meses"
	sqlStmt := `SELECT name FROM sqlite_master WHERE type='table' AND name=?;`
	err := db.QueryRow(sqlStmt, tabla).Scan(&tabla)
	// verifica si la tabla existe
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("La tabla no existe.")
			return "FALSE", nil
		} else {
			fmt.Println("Ocurrió un error:", err)
			return "FASLE", err
		}
	} else {
		// continua con el programa
		// fmt.Println("La tabla " + tabla+ " existe.")
		return "TRUE", nil
	}
}

/*
Funcion para la creacion de las tablas sqlStm es es statment
para la generacion de la tabla, los statment esta  el el script de las VariablesyConstantes
este es un ejemplo

	 // sqlStmt :=
		// 	`CREATE TABLE IF NOT EXISTS meses (
		// 		id INTEGER PRIMARY KEY,
		// 		mes TEXT,
		// 		);
		// 	`
*/
func CrearTablas(sqlStm string) {
	db := DBconn()
	defer db.Close()
	_, err := db.Exec(sqlStm)

	if err != nil {
		log.Printf("%q: %s\n", err, sqlStm)
		return
	}
}

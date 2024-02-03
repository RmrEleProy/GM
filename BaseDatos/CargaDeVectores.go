package BaseDatos

import (
	"database/sql"
	"log"
)

// carga en un arreglo los meses almacenados en la base de datos con el mismo nombre
func CargarVectorMeses() []Meses {
	db := DBconn()
	defer db.Close()

	sqldb, err := db.Query(StmSelectMeses)
	if err != nil {
		log.Fatal(err)
	}

	mess := Meses{}
	messvector := []Meses{}

	for sqldb.Next() {
		var id int
		var mes string
		err = sqldb.Scan(&id, &mes)
		if err != nil {
			log.Fatal(err)
		}
		mess.Id = id
		mess.Mes = mes

		messvector = append(messvector, mess)
	}
	return messvector
}

// carga en un arreglo todo el contenido de la tabla Gastos Mensuales
func CargarVectorGastoMensualpormes(idmes int) []GastoMensual {
	db := DBconn()
	defer db.Close()

	sqldb, err := db.Query(SelectGastoMensualByIdmes, idmes)
	if err != nil {
		log.Fatal(err)
	}

	gm := GastoMensual{}
	gmvector := []GastoMensual{}

	for sqldb.Next() {
		var id, idmes int
		var importe float64
		var fecha, concepto, tipocosto string
		err = sqldb.Scan(&id, &idmes, &fecha, &importe, &concepto, &tipocosto)
		if err != nil {
			log.Fatal(err)
		}
		gm.Id = id
		gm.Idmes = idmes
		gm.Fecha = fecha
		gm.Importe = importe
		gm.Concepto = concepto
		gm.Tipocosto = tipocosto

		gmvector = append(gmvector, gm)
	}

	return gmvector
}

func CargarVectorGastoMensualTodos() []GastoMensual {
	db := DBconn()
	defer db.Close()

	sqldb, err := db.Query(SelectGastoMensual)
	if err != nil {
		log.Fatal(err)
	}

	gm := GastoMensual{}
	gmvector := []GastoMensual{}

	for sqldb.Next() {
		var id, idmes int
		var importe float64
		var fecha, concepto, tipocosto string
		err = sqldb.Scan(&id, &idmes, &fecha, &importe, &concepto, &tipocosto)
		if err != nil {
			log.Fatal(err)
		}
		gm.Id = id
		gm.Idmes = idmes
		gm.Fecha = fecha
		gm.Importe = importe
		gm.Concepto = concepto
		gm.Tipocosto = tipocosto

		gmvector = append(gmvector, gm)
	}
	return gmvector
}

func CargarVectorTtMes(idmes int) []TotalPorMes {
	db := DBconn()
	defer db.Close()

	var sumatoria sql.NullFloat64
	
	err := db.QueryRow(SelectSumImporGastoMensual, idmes).Scan(&sumatoria)
	if err != nil {
		log.Fatal(err)
	}

	if sumatoria.Valid {
		insForm, err := db.Prepare(UpdateTotalMesByIdmes)
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(sumatoria.Float64, idmes)
	}


	sqldb, err := db.Query(SelectTotalMesByIdmes, idmes)
	if err != nil {
		log.Fatal(err)
	}
	ttmes := TotalPorMes{}
	ttmesvector := []TotalPorMes{}

	for sqldb.Next() {
		var id, idmes int
		var totalesmeses float64
		err = sqldb.Scan(&id, &idmes, &totalesmeses)
		if err != nil {
			log.Fatal(err)
		}
		ttmes.Id = id
		ttmes.Idmes = idmes
		ttmes.Totalmes = totalesmeses

		ttmesvector = append(ttmesvector, ttmes)
	}

	return ttmesvector
}

package BaseDatos

import (
	"database/sql"
	"encoding/json"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"

	"html/template"
	"net/http"
)

func DBconn() (db *sql.DB) {
	db, err := sql.Open("sqlite3", DirecionBaseDatosGAU)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func templates() *template.Template {
	return template.Must(template.ParseGlob("Templates/*.html"))
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if !EjecutadaFuncionDeInicio {
		Principal()
		EjecutadaFuncionDeInicio = true
	}
	type vectores struct {
		Vmeses   []Meses
		VGM      []GastoMensual
		Totalmes []TotalPorMes
	}
	vecmes := CargarVectorMeses()
	vecGM := CargarVectorGastoMensualTodos()
	ttm := CargarVectorTtMes(0)

	mivector := vectores{Vmeses: vecmes, VGM: vecGM, Totalmes: ttm}
	templates().ExecuteTemplate(w, "index.html", mivector)
}

// muestra Cuando se selecciona un solo mes del Menu lateral
func MuestraSoloElMesSeleccionado(w http.ResponseWriter, r *http.Request) {
	type vectores struct {
		Vmeses   []Meses
		VGM      []GastoMensual
		Totalmes []TotalPorMes
	}
	vecmes := CargarVectorMeses()
	idm, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		// maneja el error aquí
		log.Fatal(err)
	}
	vecGM := CargarVectorGastoMensualpormes(idm)

	ttmes := CargarVectorTtMes(idm)

	mivector := vectores{Vmeses: vecmes, VGM: vecGM, Totalmes: ttmes}
	templates().ExecuteTemplate(w, "index.html", mivector)

}

// devuelve el JSON de la tabla al accedel a la direccion /showJSON
func DevuelveJson(w http.ResponseWriter, r *http.Request) {
	type vectores struct {
		Vmeses   []Meses
		VGM      []GastoMensual
		Totalmes []TotalPorMes
	}
	vecmes := CargarVectorMeses()
	vecGM := CargarVectorGastoMensualTodos()
	ttm := CargarVectorTtMes(0)

	mivector := vectores{Vmeses: vecmes, VGM: vecGM, Totalmes: ttm}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mivector)
}

// Insertar un nuevo registro
func Insert(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		templates().ExecuteTemplate(w, "newexpence.html", nil)
	case http.MethodPost:
		db := DBconn()
		// log.Println(r.Method)
		if r.Method == http.MethodPost {
			idmes := r.FormValue("idmes")
			fecha := r.FormValue("fecha")
			importe := r.FormValue("importe")
			concepto := r.FormValue("concepto")
			tipocosto := r.FormValue("tipocosto")
			insForm, err := db.Prepare(InsertarEnGm)
			if err != nil {
				panic(err.Error())
			}
			insForm.Exec(idmes, fecha, importe, concepto, tipocosto)
			// log.Println("Registro insertado correctamente")
		}

		defer db.Close()
		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

// Eliminar un registro
func Delete(w http.ResponseWriter, r *http.Request) {
	bd := DBconn()
	defer bd.Close()
	emp := r.URL.Query().Get("id")
	delForm, err := bd.Prepare("DELETE FROM gasto_mensual WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := DBconn()
	defer db.Close()
	switch r.Method {
	case http.MethodGet:
		Id := r.URL.Query().Get("id")
		selDB, err := db.Query(SelectGastoMensualById, Id)
		if err != nil {
			panic(err.Error())
		}
		type gasme struct {
			GastoMensual
			Mes string
		}
		gm := gasme{}
		for selDB.Next() {
			var id, idmes int
			var importe float64
			var fecha, concepto, tipocosto string
			err = selDB.Scan(&id, &idmes, &fecha, &importe, &concepto, &tipocosto)
			if err != nil {
				panic(err.Error())
			}
			gm.Id = id
			gm.Idmes = idmes
			gm.Fecha = fecha
			gm.Importe = importe
			gm.Concepto = concepto
			gm.Tipocosto = tipocosto
		}

		idm := r.URL.Query().Get("idm")
		selDB, err = db.Query(SelectMesByID, idm)
		if err != nil {
			panic(err.Error())
		}
		for selDB.Next() {
			var mes string
			err = selDB.Scan(&mes)
			if err != nil {
				panic(err.Error())
			}
			gm.Mes = mes
		}

		templates().ExecuteTemplate(w, "edit.html", gm)

	case http.MethodPost:
		id := r.FormValue("id")
		idmes := r.FormValue("Idmes")
		fecha := r.FormValue("fecha")
		importe := r.FormValue("importe")
		concepto := r.FormValue("concepto")
		tipocosto := r.FormValue("tipocosto")
		insForm, err := db.Prepare(UpdateGM)
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(idmes, fecha, importe, concepto, tipocosto, id)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

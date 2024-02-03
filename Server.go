package main

import (
	"GastosMensuales/BaseDatos"
	"log"

	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", BaseDatos.Index)
	mux.HandleFunc("/index", BaseDatos.Index)
	mux.HandleFunc("/mes", BaseDatos.MuestraSoloElMesSeleccionado)
	mux.HandleFunc("/newexpence", BaseDatos.Insert)
	mux.HandleFunc("/Insert", BaseDatos.Insert)
	mux.HandleFunc("/edit", BaseDatos.Edit)
	mux.HandleFunc("/Update", BaseDatos.Edit)
	mux.HandleFunc("/Delete", BaseDatos.Delete)

	mux.HandleFunc("/showJSON", BaseDatos.DevuelveJson)

	mux.Handle("/Static/", http.StripPrefix("/Static/", http.FileServer(http.Dir("./Static"))))

	log.Fatal((http.ListenAndServe(":80", mux)))
}

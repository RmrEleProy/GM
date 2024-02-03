function ValidarFormulario() {
    var campos = ["fecha", "importe", "concepto"]; // Reemplaza esto con los nombres de tus campos
    for (var i = 0; i < campos.length; i++) {
        var valor = document.forms["miFormulario"][campos[i]].value;
        if (valor == "") {
            alert("¡Todos los campos deben ser completados!");
            return false;
        }
    }
}
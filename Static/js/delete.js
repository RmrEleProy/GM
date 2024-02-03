function Delete() {
    var r = confirm("Estás seguro de que quieres eliminar el elemento seleccionado? Esta acción no se puede deshacer.");
    if (r == true) {
        return true
    } else {
      return false
    }
  }
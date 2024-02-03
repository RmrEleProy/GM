fetch('/showJSON').then(response => response.json()).then(data => {

// Ordena los datos por fecha
data.VGM.sort((a, b) => new Date(a.Fecha) - new Date(b.Fecha));

var fechas = data.VGM.map(function(item) {
    return item.Fecha;
});

var importes = data.VGM.map(function(item) {
    return item.Importe;
});

var ctx = document.getElementById('graficatotal').getContext('2d');
var chart = new Chart(ctx, {
    type: 'line',
    data: {
        labels: fechas,
        datasets: [{
            label: 'Importe',
            data: importes,
            fill: false,
            borderColor: 'rgb(100, 92, 92)',
            tension: 0.1
        }]
    },
    options: {
        scales: {
            x: {
                type: 'time',
                time: {
                    unit: 'month'
                }
            }
        }
    }
});

});


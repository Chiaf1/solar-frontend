// Mappa contentenente le istanze dei vari chart creati
const chartInstances = {}

function createLineChart(canvasId, jsonId, label) {
    
    const canvas = document.getElementById(canvasId);
    const rawData = document.getElementById(jsonId);
    
    if (!canvas || !rawData) {
        console.warn(`Grafico ${canvasId} non inizializzato`);
        return;
    }

    const data = JSON.parse(rawData.textContent);

    // Distruzione chart vecchio se già esistente
    if (chartInstances[canvasId]) {
        chartInstances[canvasId].destroy();
    }

    chartInstances[canvasId] = new Chart(canvas, {
        type: "line",
        data: {
            labels: data.labels,
            datasets: [
                {
                    label: label,
                    data: data.values,
                    borderColor: "#f4b400",
                    backgroundColor: "rgba(244, 180, 0, 0.2)",
                    tension: 0.3
                }
            ]
        },
        options: {
            responsive: true,
            maintainAspectRatio: false,
            animation: false
        }
    });

} 

// Helper function to initialize charts
function initCharts() {
    createLineChart("chart-today", "chart-today-data", "Produzione oggi");
    createLineChart("chart-yesterday", "chart-yesterday-data", "Produzione ieri");
}

// Helper function to update charts data
function updateChart( chartId, newData) {
    const chart = chartInstances[chartId];
    if (!chart) return;

    chart.data.labels = newData.labels;
    chart.data.datasets[0].data = newData.values;
    chart.update('none');
}

// Inizializzazione al primo caricamento
document.addEventListener("DOMContentLoaded", initCharts);

// Ascolto per l'evento creato dalla chiamata htmx per aggiornare solo i dati del grafico today
document.body.addEventListener("updateChartToday", evt => {
    updateChart("chart-today", evt.detail);
})

document.body.addEventListener("updateChartYesterday", evt => {
    updateChart("chart-yesterday", evt.detail);
})
// Mappa contentenente le istanze dei vari chart creati
const chartInstances = {}

function createLineChart(canvasId, jsonId, labelPrefix) {
    
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
                    label: `Produzione`,
                    data: data.production,
                    borderColor: "#34a853",
                    backgroundColor: "rgba(52, 168, 83, 0.1)",
                    tension: 0.3,
                    fill: true
                },
                {
                    label: `Consumo`,
                    data: data.consumption,
                    borderColor: "#f4b400",
                    backgroundColor: "rgba(244, 180, 0, 1)",
                    tension: 0.3,
                    fill: false
                }
            ]
        },
        options: {
            responsive: true,
            maintainAspectRatio: false,
            animation: false,

            elements: {
                line: {
                    borderWidth: 1.5
                },
                point: {
                    radius: 0,
                    hoverRadius: 4,
                    hitRadius: 8
                }
            },

            scales: {
                y: {
                    beginAtZero: true,
                    ticks: {
                        font: {
                            size: 18
                        }
                    },
                    title: {
                        display: true,
                        text: "KW",
                        font: {
                            size: 14,
                            weight: "500"
                        }
                    },
                    grid: {
                        color: "#eee"
                    }
                },
            },

            plugins: {
                title: {
                    display : true,
                    text: labelPrefix,
                    align: "start",
                    font: {
                        size: 16,
                        weight: "600"
                    },
                    padding: {
                        top: 4,
                        bottom: 10
                    }
                },
                legend: {
                    position: "top",
                    align: "end",
                    labels: {
                        boxWidth: 12,
                        boxHeight: 12
                    }
                }
            },
        }
    });

} 

// Helper function to initialize charts
function initCharts() {
    createLineChart("chart-today", "chart-today-data", "Produzione oggi");
    createLineChart("chart-yesterday", "chart-yesterday-data", "Produzione ieri");
    createLineChart("chart-minus-2", "chart-minus-2-data", "Produzione -2");
    createLineChart("chart-minus-3", "chart-minus-3-data", "Produzione -3");
    createLineChart("chart-minus-4", "chart-minus-4-data", "Produzione -4");
    createLineChart("chart-minus-5", "chart-minus-5-data", "Produzione -5");
    createLineChart("chart-minus-6", "chart-minus-6-data", "Produzione -6");
}

// Helper function to update charts data
function updateChart( chartId, newData) {
    const chart = chartInstances[chartId];
    if (!chart) return;

    chart.data.labels = newData.labels;
    chart.data.datasets[0].data = newData.production;
    chart.data.datasets[1].data = newData.consumption;
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

document.body.addEventListener("updateChartHistory", evt => {
    const updates = evt.detail;
    Object.entries(updates).forEach(([chartId, data]) => {
        updateChart(chartId, data);
    });
})
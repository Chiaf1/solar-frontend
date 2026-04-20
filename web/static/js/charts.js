
function createLineChart(canvasId, jsonId, label) {
    
    const canvas = document.getElementById(canvasId);
    const rawData = document.getElementById(jsonId);
    
    if (!canvas || !rawData) {
        console.warn(`Grafico ${canvasId} non inizializzato`);
        return;
    }

    const data = JSON.parse(rawData.textContent);

    new Chart(canvas, {
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
            scales: {
                y: {
                    beginAtZero: true
                }
            }
        }
    });

}

createLineChart(
    "chart-today",
    "chart-today-data",
    "Produzione oggi",
);

createLineChart(
    "chart-yesterday",
    "chart-yesterday-data",
    "Produzione ieri",
);
const canvasToday = document.getElementById("chart-today");

if (!canvasToday) {
    console.warn("Canvas chart-today non trovato");
} else {
    const labels = [
        "00:00", "02:00", "04:00", "06:00",
        "08:00", "10:00", "12:00", "14:00",
        "16:00", "18:00", "20:00", "22:00",
    ];

    const productionData = [
        0, 0, 0, 0.5,
        1.2, 2.8, 4.1, 3.9,
        2.6, 1.1, 0.3, 0,
    ];

    new Chart(canvasToday, {
        type: "line",
        data: {
            labels: labels,
            datasets: [
                {
                    label: "Produzione oggi",
                    data: productionData,
                    borderColor: "#f4b400",
                    backgroundColor: "rgba(244, 180, 0, 0.2)",
                    tension: 0.3
                }
            ]
        },
        options: {
            responsive: true,
            mantainAspectRatio: false,
            scales: {
                y: {
                    beginAtZerto: true
                }
            }
        }
    });
}
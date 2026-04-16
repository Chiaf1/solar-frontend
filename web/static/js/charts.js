const canvasToday = document.getElementById("chart-today");
const rawData = document.getElementById("chart-today-data");

if (canvasToday && rawData) {
    const chartData = JSON.parse(rawData.textContent);

    new Chart(canvasToday, {
        type: "line",
        data: {
            labels: chartData.labels,
            datasets: [
                {
                    label: "Produzione oggi",
                    data: chartData.values,
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
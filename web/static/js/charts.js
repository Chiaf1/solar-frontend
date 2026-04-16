const ctx = document.getElementById("chart-today");

if (ctx) {
  new Chart(ctx, {
    type: "line",
    data: {
      labels: ["00:00", "02:00", "04:00", "06:00", "08:00", "10:00"],
      datasets: [
        {
          label: "Produzione",
          data: [0, 0, 1.2, 3.5, 5.8, 6.1],
          borderColor: "rgb(75, 192, 192)",
          tension: 0.3
        },
        {
          label: "Consumo",
          data: [0.8, 0.9, 1.1, 1.6, 2.0, 2.4],
          borderColor: "rgb(255, 99, 132)",
          tension: 0.3
        }
      ]
    },
    options: {
      responsive: true,
      maintainAspectRatio: false
    }
  });
}
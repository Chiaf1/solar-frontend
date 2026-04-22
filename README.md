# Solar Energy Dashboard

A web-based dashboard for visualizing **solar energy production and consumption** data in near real-time.

The project is built with **Go**, **Gin**, **HTMX**, and **Chart.js**, and is designed to be:
- lightweight
- easy to maintain
- suitable for large screens (TVs, industrial monitors)
- robust against missing or partial data

---

## ✨ Features

- **Today chart** with near real-time updates (production & consumption)
- **Historical charts** (Yesterday and −2 to −6 days)
- **Live KPIs** (current production & consumption)
- **Localized date and time** (Italian)
- **Smart refresh strategy** (data updates only when needed)
- **Graceful handling of missing data**
- **Responsive layout** (desktop-first, mobile-friendly as a bonus)

---

## 🧱 Tech Stack

### Backend
- **Go**
- **Gin** – HTTP routing
- **In-memory caching** (TTL-based)
- External REST API as data source

### Frontend
- **HTMX** – partial updates without heavy JavaScript
- **Chart.js** – data visualization
- **Go HTML templates**
- **CSS Grid & Flexbox** – responsive layout

---

## 🧠 Architecture Overview

The application follows a **layered architecture**:

```

Handlers  →  Services  →  Repositories  →  External API

```

### Handlers
- HTTP endpoints (pages, partials, API refresh)
- No business logic
- Responsible only for input/output

### Services
- Business logic
- In-memory caching
- KPI computation
- Time and date formatting
- Centralized data reuse (charts + KPIs)

### Repositories
- Communication with the external energy API
- HTTP requests and JSON decoding
- Mapping API responses to domain models

---

## 📊 Charts & Time Handling

### Data model
Each chart uses:
- a shared **time axis (00:00 → 24:00)**
- two series:
  - Production (kW)
  - Consumption (kW)

Missing data points are represented as `null`, allowing Chart.js to render gaps instead of misleading zeros.

### Time normalization
- API data is received in **UTC**
- Converted to **Europe/Rome**
- Aggregated into fixed buckets (2m, 10m, etc.)
- All charts are aligned on the same daily timeline

This guarantees:
- visual consistency
- correct comparison between days
- robustness against backend aggregation offsets

---

## 🔄 Refresh Strategy (HTMX)

Different data updates at different rates, based on how often it actually changes:

| Component        | Refresh | Backend Cache |
|------------------|---------|---------------|
| Today chart      | every 10s | 30s TTL |
| KPIs             | every 10s | same cache |
| Header (date/time)| every 60s | none |
| History charts   | on load (+ optional 6h) | long TTL |

This avoids unnecessary API calls while keeping the UI responsive.

---

## 📈 KPIs

KPIs are derived from the **same cached data** used by the Today chart.

- The “current” value is the **last non-null data point**
- No additional API calls are made
- If data is missing, KPIs gracefully show placeholders

---

## 🖥️ Layout & Design Principles

- Desktop / TV first (no scrolling on normal screens)
- Two main columns:
  - **Left:** large charts (Today / Yesterday)
  - **Right:** header, KPIs, history charts
- CSS Grid used for macro layout
- Flexbox used inside components
- Neutral colors, clear hierarchy

Mobile support is included but not the primary target.

---

## 🧪 Missing or Partial Data

The system is designed to handle:
- API outages
- empty days
- partial time ranges

Behavior:
- Charts are still rendered
- Missing values appear as gaps
- No crashes or layout breakage

---

## 🚀 Future Improvements

Possible next steps:
- Persistent cache (Redis) for multi-instance deployments
- Dark mode / theme support
- User-selectable time ranges
- Export data (CSV / image)
- Authentication & multi-plant support

---

## 📝 Notes

This project intentionally avoids heavy frontend frameworks.
The goal is to keep the system:
- understandable
- debuggable
- close to the data

HTMX + Go templates are used to minimize complexity while still delivering a modern UX.


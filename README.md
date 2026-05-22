![Go](https://img.shields.io/badge/Go-1.22-00ADD8?style=for-the-badge&logo=go)

![Google Maps](https://img.shields.io/badge/Google_Maps-3D-4285F4?style=for-the-badge&logo=googlemaps)

![OpenStreetMap](https://img.shields.io/badge/OpenStreetMap-GeoPlatform-7EBC6F?style=for-the-badge&logo=openstreetmap)

![License](https://img.shields.io/badge/License-Apache_2.0-orange?style=for-the-badge)
# 🌐 Live Demo

Coming Soon

https://selikasela.cyou
# 🌍 GeoVision3D

### Real-Time Geo Intelligence & 3D Mapping Platform

GeoVision3D is a modern geo-intelligence platform built with Go, Google Maps 3D, OpenStreetMap integration, and real-time geolocation services.

Designed for tourism, spatial visualization, location intelligence, and mapping systems, the platform combines high-performance backend architecture with immersive 3D map rendering.

---

# 🚀 Features

- 🌍 Real-time IP geolocation
- 🛰️ ASN & ISP lookup intelligence
- 🗺️ Google Maps 3D integration
- 🌐 OpenStreetMap fallback engine
- ⚡ High-performance Go backend
- 📡 REST API architecture
- 📍 Live location visualization
- 🌌 Dark luxury UI interface
- 🔐 Secure environment variable configuration
- ☁️ Production deployment ready

---

# 🏗️ Architecture

```text
Visitor
   ↓
Go Backend API
   ↓
IP Geolocation Service
   ↓
ASN Intelligence Layer
   ↓
Map Configuration Engine
   ↓
Google Maps 3D / OpenStreetMap
   ↓
Frontend Visualization
```

---

# 📂 Project Structure

```text
geo-vision-3d/

├── main.go
├── go.mod
├── README.md
│
├── static/
│   ├── index.html
│   ├── app.js
│   ├── styles.css
│   └── assets/
│
├── api/
│   ├── geolocation.go
│   ├── maps.go
│   └── asn.go
│
├── middleware/
│   ├── cors.go
│   └── logger.go
│
├── utils/
│   └── env.go
│
└── .env
```

---

# ⚡ API Example

```json
{
  "ip": "103.137.82.147",
  "asn": "AS138828",
  "city": "Badung",
  "country": "Indonesia",
  "latitude": -8.7933893,
  "longitude": 115.1227501,
  "org": "Seli Kasela Creative Studio"
}
```

---

# 🧠 Technologies Used

- Go
- Gorilla Mux
- Google Maps API
- OpenStreetMap
- Leaflet.js
- HTML5
- CSS3
- JavaScript

---

# 🔐 Security

- Store API keys in `.env`
- Restrict Google Maps API access
- Never expose secrets publicly

---

# 👤 Author

## Seli Kasela

Founder • Digital Business Architect • Geo & Tourism Technology Builder

📍 Badung, Bali, Indonesia

---

# 📜 License

Apache License 2.0

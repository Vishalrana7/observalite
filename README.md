🚀 Observalite — Lightweight Open-Source Observability Collector

A minimal, fast, Go-based observability collector that receives events and stores them in PostgreSQL.

📌 Overview

Observalite is a lightweight observability collector designed to receive structured application events and store them efficiently in PostgreSQL for further analysis, dashboards, or alerting.

This version contains:

✅ Go-based HTTP collector
✅ /v1/events endpoint to ingest events
✅ PostgreSQL integration using pgx
✅ Dockerized Postgres
✅ Migrations
✅ Local development setup
✅ Event JSON format

This is Day 1 of development.

📁 Project Structure
observalite/
│
├── collector/
│   ├── cmd/server/main.go
│   ├── handlers/event.go
│   ├── db/db.go
│   ├── models/event.go
│
├── migrations/
│   └── 001_create_events.sql
│
├── docker-compose.yml  (optional day 2)
│
├── go.mod
├── go.sum
└── README.md

🔧 Requirements

Go 1.21+

Docker + Docker Compose

PostgreSQL (local or container)

PowerShell / Bash for commands

🐘 Database Setup
1. Run PostgreSQL via Docker
docker run -d --name observalite-db -e POSTGRES_USER=user -e POSTGRES_PASSWORD=pass -e POSTGRES_DB=observalite -p 5432:5432 postgres:15

2. Apply migration

PowerShell:

Get-Content migrations/001_create_events.sql | docker exec -i observalite-db psql -U user -d observalite


SQL schema:

CREATE TABLE IF NOT EXISTS events (
  id BIGSERIAL PRIMARY KEY,
  service TEXT NOT NULL,
  instance_id TEXT,
  endpoint TEXT NOT NULL,
  method TEXT,
  status INTEGER,
  latency_ms INTEGER,
  size_bytes INTEGER,
  tags JSONB,
  timestamp timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_events_service_endpoint_ts 
  ON events (service, endpoint, timestamp DESC);

⚙️ Environment Variable

The collector reads:

POSTGRES_URL=postgres://user:pass@localhost:5432/observalite

PowerShell
$env:POSTGRES_URL="postgres://user:pass@localhost:5432/observalite"

Linux / Mac
export POSTGRES_URL=postgres://user:pass@localhost:5432/observalite

▶️ Run the Collector

Run the Go server:

go run ./collector/cmd/server/main.go


Expected output:

Collector listening on :8080

📥 POST /v1/events (Send Logs/Events)
PowerShell (Invoke-WebRequest)
Invoke-RestMethod -Method POST "http://localhost:8080/v1/events" `
  -Headers @{ "Content-Type" = "application/json" } `
  -Body '{"service":"auth","instance_id":"auth-1","endpoint":"/login","method":"POST","status":200,"latency_ms":120,"size_bytes":456,"tags":{"env":"dev"}}'

Linux/macOS (curl)
curl -X POST http://localhost:8080/v1/events \
  -H "Content-Type: application/json" \
  -d '{
    "service": "auth",
    "instance_id": "auth-1",
    "endpoint": "/login",
    "method": "POST",
    "status": 200,
    "latency_ms": 120,
    "size_bytes": 456,
    "tags": { "env": "dev" }
  }'

🗄️ Query Events
SELECT * FROM events;

🧪 Example Output (Internal Logs)

When an event is received, logs show:

===> Hit /v1/events
POST /v1/events called
Decoded event: {Service:auth InstanceID:auth-1 Endpoint:/login Method:POST Status:200 LatencyMs:120 SizeBytes:456 Tags:map[env:dev]}
Insert successful

🎯 Roadmap (Day-by-Day Plan)
Day 1 (✔ Done)

Basic collector

PostgreSQL insert

Simple handler

JSON decoding

Migrations

Manual testing

Day 2

Docker Compose for entire system

Health checks

Internal metrics

Basic dashboard UI (optional)

Day 3

Aggregation queries (latency, error rate)

List & filter endpoints

Time-range queries

Day 4

Ingest logs & metrics

API token authentication

Day 5

Frontend dashboard

Charts (status codes, latency, throughput)

🤝 Contributing

PRs are welcome — the project is modular and easy to extend.

📜 License

MIT License.
📘 Observalite — Lightweight Observability Platform (Work-in-Progress)

Observalite is an open-source, lightweight observability platform that I am building from scratch.
The goal is to create a simple, high-performance event collector and analytics system that developers can easily run locally or deploy in production without the complexity of full observability stacks like OpenTelemetry, Loki, or Elastic.

This project starts with a minimal core — a Go-based collector — and will expand step-by-step into a fully functioning observability suite.

🎯 What This Project Aims to Achieve
1. Lightweight Observability Without Heavy Tools

Most existing observability systems require large setups and complex agents.
Observalite aims to deliver:

A tiny collector

Easy event ingestion

Fast queries

Low resource usage

Perfect for small teams, side projects, or personal infra.

2. Build a Real Observability Pipeline From Scratch

The full system will cover:

Event ingestion (HTTP collector)

Log, metric, and tracing support

PostgreSQL-backed storage

Dashboard visualizations

Service-level insights like:

Latency

Error rates

Throughput

Endpoint performance

This project acts as a real-world learning experience in designing distributed telemetry systems.

3. A Modular, Extensible Architecture

The collector is intentionally small and clean, making it easy to extend with:

Authentication

Batching

Aggregation workers

Exporters

Alerts

Tracing support

4. Developer-Friendly, Local-First Setup

Everything can run locally with minimal commands — Postgres + Go collector.
The goal is to deliver a great onboarding experience for developers trying to understand telemetry systems.

🧱 Current Status (Day 1 Complete)

Basic Go collector working

/v1/events POST endpoint

JSON decoding

Insert events into PostgreSQL

Local testing verified

Migrations added

This is the foundation on which the entire observability platform will be built.

🚀 Upcoming Features

Docker Compose for full local stack

Query APIs (filter, aggregate, time-range)

Metrics ingestion

Logs ingestion

Web dashboard (latency charts, errors, endpoints list)

API keys & authentication

In-memory caching for fast analytics
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
CREATE INDEX IF NOT EXISTS idx_events_service_endpoint_ts ON events (service, endpoint, timestamp DESC);

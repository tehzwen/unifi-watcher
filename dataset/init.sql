CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE events (
    id TEXT DEFAULT uuid_generate_v4()::TEXT PRIMARY KEY,
    "cameraId" TEXT,
    "type" TEXT,
    "createdAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "smartDetectTypes" json
);

-- insert some fake events
INSERT INTO events("cameraId", "type", "createdAt", "smartDetectTypes") VALUES('12345', 'smartDetectZone', '2026-03-03 00:25:00.346+00', '["vehicle"]');
INSERT INTO events("cameraId", "type", "createdAt", "smartDetectTypes") VALUES('12345', 'smartDetectZone', '2026-03-03 00:30:00.346+00', '["person"]');

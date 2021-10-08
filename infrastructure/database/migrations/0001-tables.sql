CREATE TABLE IF NOT EXISTS public.alerts (
  id CHAR(50) PRIMARY KEY,
  priority VARCHAR(50) NULL,
  source VARCHAR(250) NULL,
  message VARCHAR(250) NULL,
  report_ack_time BIGINT NOT NULL DEFAULT 0,
  report_close_time BIGINT NOT NULL DEFAULT 0,
  integration_id VARCHAR(50) NULL,
  colleted_at TIMESTAMP NOT NULL DEFAULT now(),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE INDEX alerts_priority ON public.alerts (priority);
CREATE INDEX alerts_integration_id ON public.alerts (integration_id);
CREATE INDEX alerts_source ON public.alerts (source);
CREATE INDEX alerts_created_at ON public.alerts (created_at);

CREATE TABLE IF NOT EXISTS public.integrations (
  id VARCHAR(50) PRIMARY KEY,
  name VARCHAR(250) NOT NULL,
  type VARCHAR(250) NOT NULL,
  enabled BOOLEAN NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS public.teams (
  id VARCHAR(50) PRIMARY KEY,
  name VARCHAR(250) NULL,
  description VARCHAR(250) NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

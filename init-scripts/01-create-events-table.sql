-- Create events table for analytics tracking
CREATE TABLE
  public.events (
    id serial NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    user_id character varying(255) NOT NULL,
    event_type character varying(255) NOT NULL,
    event_url text NULL,
    event_data jsonb NOT NULL
  );

ALTER TABLE
  public.events
ADD
  CONSTRAINT events_pkey PRIMARY KEY (id);

-- Create index on user_id for better query performance
CREATE INDEX idx_events_user_id ON public.events(user_id);

-- Create index on event_type for filtering
CREATE INDEX idx_events_event_type ON public.events(event_type);

-- Create index on created_at for time-based queries
CREATE INDEX idx_events_created_at ON public.events(created_at);

-- Grant permissions to the analytics user
GRANT ALL PRIVILEGES ON TABLE public.events TO analytics_user;
GRANT USAGE, SELECT ON SEQUENCE public.events_id_seq TO analytics_user;

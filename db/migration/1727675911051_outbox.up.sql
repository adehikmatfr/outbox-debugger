BEGIN;

CREATE TABLE IF NOT EXISTS public.event_outbox1 (
    event_outbox_id uuid not null constraint event_outbox1_pk primary key,
    event_group varchar(100) not null,
    event_topic varchar(100) not null,
    event_key varchar(100) not null,
    event_message bytea not null,
    retry_count integer not null,
    last_retry_time_utc timestamp not null,
    next_retry_time_utc timestamp not null,
    status varchar(15) not null,
    hash_value1 varchar(75) not null,
    created_time_utc timestamp not null,
    updated_time_utc timestamp not null,
    row_version uuid not null
);

CREATE INDEX IF NOT EXISTS event_outbox1_next_retry_time_utc_status_index
    on public.event_outbox1 (status, next_retry_time_utc);

CREATE INDEX IF NOT EXISTS event_outbox1_event_idx1_index
    on public.event_outbox1 (event_group, event_topic, event_key);
    
CREATE INDEX IF NOT EXISTS event_outbox1_hash_value1_index
    on public.event_outbox1 (hash_value1);


CREATE TABLE IF NOT EXISTS public.event_outbox2 (
    event_outbox_id uuid not null constraint event_outbox2_pk primary key,
    event_group varchar(100) not null,
    event_topic varchar(100) not null,
    event_key varchar(100) not null,
    event_message bytea not null,
    retry_count integer not null,
    last_retry_time_utc timestamp not null,
    next_retry_time_utc timestamp not null,
    status varchar(15) not null,
    hash_value1 varchar(75) not null,
    created_time_utc timestamp not null,
    updated_time_utc timestamp not null,
    row_version uuid not null
);

CREATE INDEX IF NOT EXISTS event_outbox2_next_retry_time_utc_status_index
    on public.event_outbox2 (status, next_retry_time_utc);

CREATE INDEX IF NOT EXISTS event_outbox2_event_idx1_index
    on public.event_outbox2 (event_group, event_topic, event_key);
    
CREATE INDEX IF NOT EXISTS event_outbox2_hash_value1_index
    on public.event_outbox2 (hash_value1);


CREATE TABLE IF NOT EXISTS public.event_outbox3 (
    event_outbox_id uuid not null constraint event_outbox3_pk primary key,
    event_group varchar(100) not null,
    event_topic varchar(100) not null,
    event_key varchar(100) not null,
    event_message bytea not null,
    retry_count integer not null,
    last_retry_time_utc timestamp not null,
    next_retry_time_utc timestamp not null,
    status varchar(15) not null,
    hash_value1 varchar(75) not null,
    created_time_utc timestamp not null,
    updated_time_utc timestamp not null,
    row_version uuid not null
);

CREATE INDEX IF NOT EXISTS event_outbox3_next_retry_time_utc_status_index
    on public.event_outbox3 (status, next_retry_time_utc);

CREATE INDEX IF NOT EXISTS event_outbox3_event_idx1_index
    on public.event_outbox3 (event_group, event_topic, event_key);
    
CREATE INDEX IF NOT EXISTS event_outbox3_hash_value1_index
    on public.event_outbox3 (hash_value1);


CREATE TABLE IF NOT EXISTS public.event_outbox4 (
    event_outbox_id uuid not null constraint event_outbox4_pk primary key,
    event_group varchar(100) not null,
    event_topic varchar(100) not null,
    event_key varchar(100) not null,
    event_message bytea not null,
    retry_count integer not null,
    last_retry_time_utc timestamp not null,
    next_retry_time_utc timestamp not null,
    status varchar(15) not null,
    hash_value1 varchar(75) not null,
    created_time_utc timestamp not null,
    updated_time_utc timestamp not null,
    row_version uuid not null
);

CREATE INDEX IF NOT EXISTS event_outbox4_next_retry_time_utc_status_index
    on public.event_outbox4 (status, next_retry_time_utc);

CREATE INDEX IF NOT EXISTS event_outbox4_event_idx1_index
    on public.event_outbox4 (event_group, event_topic, event_key);
    
CREATE INDEX IF NOT EXISTS event_outbox4_hash_value1_index
    on public.event_outbox4 (hash_value1);



CREATE TABLE IF NOT EXISTS public.event_outbox5 (
    event_outbox_id uuid not null constraint event_outbox5_pk primary key,
    event_group varchar(100) not null,
    event_topic varchar(100) not null,
    event_key varchar(100) not null,
    event_message bytea not null,
    retry_count integer not null,
    last_retry_time_utc timestamp not null,
    next_retry_time_utc timestamp not null,
    status varchar(15) not null,
    hash_value1 varchar(75) not null,
    created_time_utc timestamp not null,
    updated_time_utc timestamp not null,
    row_version uuid not null
);

CREATE INDEX IF NOT EXISTS event_outbox5_next_retry_time_utc_status_index
    on public.event_outbox5 (status, next_retry_time_utc);

CREATE INDEX IF NOT EXISTS event_outbox5_event_idx1_index
    on public.event_outbox5 (event_group, event_topic, event_key);
    
CREATE INDEX IF NOT EXISTS event_outbox5_hash_value1_index
    on public.event_outbox5 (hash_value1);


COMMIT;
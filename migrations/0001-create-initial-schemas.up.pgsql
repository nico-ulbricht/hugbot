create table users (
  id uuid not null primary key,
  external_id text not null,
  name text not null
);

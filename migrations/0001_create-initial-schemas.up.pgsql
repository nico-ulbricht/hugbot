create table if not exists users (
  id uuid not null primary key,
  external_id text unique not null,
  created_at timestamp with time zone not null,
  updated_at timestamp with time zone not null
);

create table if not exists reactions (
  id uuid not null primary key,
  recipient_id uuid not null references users (id),
  sender_id uuid not null references users (id),
  reference_id text not null,
  amount int not null,
  type text not null,
  created_at timestamp with time zone not null,
  updated_at timestamp with time zone not null
);

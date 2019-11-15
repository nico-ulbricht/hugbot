create table if not exists users (
  id uuid not null primary key,
  external_id text not null,
  name text not null
);

create table if not exists reactions (
  id uuid not null primary key,
  recipient_id uuid not null references users (id),
  sender_id uuid not null references users (id),
  amount int not null,
  type text not null
);

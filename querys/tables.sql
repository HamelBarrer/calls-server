create table if not exists users.users(
	user_id bigserial primary key,
	username varchar not null unique,
	password varchar not null,
	is_active bool not null default true,
	created_at timestamp not null default now(),
	updated_at timestamp default now()
);
create table if not exists users.users(
	user_id bigserial primary key,
	username varchar not null unique,
	password varchar not null,
	avatar varchar,
	is_active bool not null default true,
	created_at timestamp not null default now(),
	updated_at timestamp default now()
);
create table if not exists users.followed_users (
	followed_user_id bigserial primary key,
	user_id int not null references users.users on update cascade on delete cascade,
	follower_user_id int not null references users.users (user_id) on update cascade on delete cascade,
	canceled_follower_user_id int references users.users (user_id) on update cascade on delete cascade,
	reactive_follower_user_id int references users.users (user_id) on update cascade on delete cascade,
	created_at timestamp not null default now(),
	updated_at timestamp default now(),
	constraint valid_followed_user check (user_id <> follower_user_id),
	constraint valid_unique unique (user_id, follower_user_id)
);
create table users (
	id    integer   PRIMARY KEY NOT NULL,
    email character varying     NOT NULL
);

create table orgs (
	id      integer PRIMARY KEY NOT NULL,
	user_id integer             NOT NULL,

	FOREIGN KEY (user_id) REFERENCES users (id)
);

create table apps (
	id integer PRIMARY KEY NOT NULL,

	user_id integer NOT NULL,
	org_id  integer UNIQUE,

	FOREIGN KEY (user_id) REFERENCES users (id),
	FOREIGN KEY (org_id)  REFERENCES orgs  (id)
);

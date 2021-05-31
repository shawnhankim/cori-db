
create table users (
	id    integer            NOT NULL,
    email character varying  NOT NULL
);

ALTER TABLE users ADD CONSTRAINT user_pkey PRIMARY KEY (id);

create table orgs (
	id      integer NOT NULL,
	user_id integer NOT NULL
);

ALTER TABLE orgs ADD CONSTRAINT org_pkey       PRIMARY KEY (id);
ALTER TABLE orgs ADD CONSTRAINT org_users_fkey FOREIGN KEY (user_id) REFERENCES users(id);

create table apps (
	id      integer NOT NULL,
	user_id integer NOT NULL,
	org_id  integer UNIQUE
);

ALTER TABLE apps ADD CONSTRAINT app_pkey       PRIMARY KEY (id);
ALTER TABLE apps ADD CONSTRAINT app_users_fkey FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE apps ADD CONSTRAINT app_orgs_fkey  FOREIGN KEY (org_id)  REFERENCES orgs(id);

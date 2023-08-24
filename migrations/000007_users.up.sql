CREATE TABLE public.users (
	id_user uuid NOT NULL DEFAULT gen_random_uuid(),
	email_user varchar(255) NULL,
	"password" varchar(255) NOT NULL,
	phone_number varchar(20) NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL,
	image_user varchar(255) NULL,
	"role" varchar(50) NULL,
	first_name varchar NULL,
	last_name varchar NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id_user)
);
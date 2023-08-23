CREATE TABLE public.directors (
	id_director uuid NOT NULL DEFAULT gen_random_uuid(),
	name_director varchar(255) NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL,
	CONSTRAINT director_pkey PRIMARY KEY (id_director)
);

INSERT INTO public.directors
(id_director, name_director)
VALUES('b52a8999-3b9d-455c-99b2-97605cfb4868'::uuid, 'Eiichiro Oda');
INSERT INTO public.directors
(id_director, name_director)
VALUES('f045c26e-49d9-4e5a-9ec2-ef85bc28714b'::uuid, 'Manabu Ono');
INSERT INTO public.directors
(id_director, name_director)
VALUES('8a7e984b-72c7-4be2-b38a-b8a0ae1b90d3'::uuid, 'Masaki Yoshimura');
INSERT INTO public.directors
(id_director, name_director)
VALUES('1009481f-747d-477c-8698-d0e89481a71b'::uuid, 'Tomoya Tanaka');
INSERT INTO public.directors
(id_director, name_director)
VALUES('1c35d887-3635-4fc3-a454-0449c5a7f065'::uuid, 'Haruo Sotozaki');


CREATE TABLE public.genres (
	id_genre uuid NOT NULL DEFAULT gen_random_uuid(),
	name_genre varchar(255) NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL,
	CONSTRAINT genre_pkey PRIMARY KEY (id_genre)
);

INSERT INTO public.genres
(id_genre, name_genre)
VALUES('d8cceb39-443d-418a-b564-6a1e4b3ddebd'::uuid, 'action');
INSERT INTO public.genres
(id_genre, name_genre)
VALUES('c85f60f7-30ef-4342-81f3-0e4ee973a3f6'::uuid, 'comedy');
INSERT INTO public.genres
(id_genre, name_genre)
VALUES('a9216527-64a0-4c3f-a66d-7ed3990f273f'::uuid, 'parody');
INSERT INTO public.genres
(id_genre, name_genre)
VALUES('9fba3d8d-b895-441e-899d-05e260c3e042'::uuid, 'fantasy');
INSERT INTO public.genres
(id_genre, name_genre)
VALUES('57ed2517-7166-4d37-b642-a20faeff1f1f'::uuid, 'school');
INSERT INTO public.genres
(id_genre, name_genre)
VALUES('a684293d-843a-4246-afbc-d7867f1e6563'::uuid, 'shounen');
INSERT INTO public.genres
(id_genre, name_genre)
VALUES('1203b743-81a2-43cc-a762-ce923dd923a8'::uuid, 'drama');
INSERT INTO public.genres
(id_genre, name_genre)
VALUES('804b159f-3da5-454f-9786-a9f8aaed4a17'::uuid, 'crime');
INSERT INTO public.genres
(id_genre, name_genre)
VALUES('a6955527-6fc3-4d9d-ad76-d1290fc8cf6d'::uuid, 'super power');
INSERT INTO public.genres
(id_genre, name_genre)
VALUES('e868c846-57d3-4e91-82c7-34d85078d93e'::uuid, 'seinen');


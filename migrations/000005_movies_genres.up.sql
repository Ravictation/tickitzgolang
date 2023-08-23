CREATE TABLE public.movies_genres (
	id_movie_genre uuid NOT NULL DEFAULT gen_random_uuid(),
	id_movie uuid NOT NULL,
	id_genre uuid NOT NULL,
	CONSTRAINT movie_genre_pkey PRIMARY KEY (id_movie_genre),
	CONSTRAINT movie_genre_id_genre_fkey FOREIGN KEY (id_genre) REFERENCES public.genres(id_genre),
	CONSTRAINT movie_genre_id_movie_fkey FOREIGN KEY (id_movie) REFERENCES public.movies(id_movie)
);

INSERT INTO public.movies_genres
(id_movie_genre, id_movie, id_genre)
VALUES('b3e72331-6253-4520-8ca9-3d781e0600e5'::uuid, '063e892d-e9c6-40e5-bb06-d4fb164d6b64'::uuid, '57ed2517-7166-4d37-b642-a20faeff1f1f'::uuid);
INSERT INTO public.movies_genres
(id_movie_genre, id_movie, id_genre)
VALUES('6019750e-a386-4003-bad6-ff8b6b397ed5'::uuid, '063e892d-e9c6-40e5-bb06-d4fb164d6b64'::uuid, 'a6955527-6fc3-4d9d-ad76-d1290fc8cf6d'::uuid);
INSERT INTO public.movies_genres
(id_movie_genre, id_movie, id_genre)
VALUES('95dd9975-74fc-4b3b-a5a5-7fc65b26e742'::uuid, '063e892d-e9c6-40e5-bb06-d4fb164d6b64'::uuid, 'd8cceb39-443d-418a-b564-6a1e4b3ddebd'::uuid);
INSERT INTO public.movies_genres
(id_movie_genre, id_movie, id_genre)
VALUES('1cbca8c3-cfb4-4748-a096-d70916e497a1'::uuid, '60677a43-41ee-4a92-8fbb-d8bc1c4ae770'::uuid, 'd8cceb39-443d-418a-b564-6a1e4b3ddebd'::uuid);
INSERT INTO public.movies_genres
(id_movie_genre, id_movie, id_genre)
VALUES('2786371a-0d49-4874-8317-bdab9f07878d'::uuid, '60677a43-41ee-4a92-8fbb-d8bc1c4ae770'::uuid, '804b159f-3da5-454f-9786-a9f8aaed4a17'::uuid);
INSERT INTO public.movies_genres
(id_movie_genre, id_movie, id_genre)
VALUES('b3865c46-8e96-4189-9de9-da832e7b3fe6'::uuid, '60677a43-41ee-4a92-8fbb-d8bc1c4ae770'::uuid, 'a6955527-6fc3-4d9d-ad76-d1290fc8cf6d'::uuid);
INSERT INTO public.movies_genres
(id_movie_genre, id_movie, id_genre)
VALUES('c13dae61-49b1-4cde-9937-2847220c2302'::uuid, 'a76db5de-c628-4c0c-b897-611e3b382874'::uuid, 'a9216527-64a0-4c3f-a66d-7ed3990f273f'::uuid);
INSERT INTO public.movies_genres
(id_movie_genre, id_movie, id_genre)
VALUES('113a2e68-dbb7-431c-bba8-35607d1674a5'::uuid, 'a76db5de-c628-4c0c-b897-611e3b382874'::uuid, 'd8cceb39-443d-418a-b564-6a1e4b3ddebd'::uuid);
INSERT INTO public.movies_genres
(id_movie_genre, id_movie, id_genre)
VALUES('c27ebb4e-6704-4739-9f85-fe0bf4636f13'::uuid, 'a76db5de-c628-4c0c-b897-611e3b382874'::uuid, '57ed2517-7166-4d37-b642-a20faeff1f1f'::uuid);
INSERT INTO public.movies_genres
(id_movie_genre, id_movie, id_genre)
VALUES('e9219c81-b9f9-459b-bd9d-d079716912d9'::uuid, 'ff2a4960-6099-4ef2-b529-c44ce70f8328'::uuid, '1203b743-81a2-43cc-a762-ce923dd923a8'::uuid);
INSERT INTO public.movies_genres
(id_movie_genre, id_movie, id_genre)
VALUES('a97a1ea3-9df8-4364-86d6-c04424ce0a92'::uuid, 'ff2a4960-6099-4ef2-b529-c44ce70f8328'::uuid, '9fba3d8d-b895-441e-899d-05e260c3e042'::uuid);
INSERT INTO public.movies_genres
(id_movie_genre, id_movie, id_genre)
VALUES('3e06546e-013a-4cc3-8381-e5104b8e34c6'::uuid, 'ff2a4960-6099-4ef2-b529-c44ce70f8328'::uuid, 'd8cceb39-443d-418a-b564-6a1e4b3ddebd'::uuid);

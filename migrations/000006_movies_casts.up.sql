CREATE TABLE public.movies_casts (
	id_movie_cast uuid NOT NULL DEFAULT gen_random_uuid(),
	id_movie uuid NOT NULL,
	id_cast uuid NOT NULL,
	CONSTRAINT movie_cast_pkey PRIMARY KEY (id_movie_cast),
	CONSTRAINT movie_cast_id_cast_fkey FOREIGN KEY (id_cast) REFERENCES public.casts(id_cast),
	CONSTRAINT movie_cast_id_movie_fkey FOREIGN KEY (id_movie) REFERENCES public.movies(id_movie)
);

INSERT INTO public.movies_casts
(id_movie_cast, id_movie, id_cast)
VALUES('2aa5abaa-dd77-4d23-a710-0642af0f4062'::uuid, '60677a43-41ee-4a92-8fbb-d8bc1c4ae770'::uuid, '23f2828b-ba86-4475-b5bb-ef6cc47eb620'::uuid);
INSERT INTO public.movies_casts
(id_movie_cast, id_movie, id_cast)
VALUES('861eb66e-9c38-4309-84d8-5380efebe897'::uuid, '60677a43-41ee-4a92-8fbb-d8bc1c4ae770'::uuid, '974301e9-9dc4-4ac8-9604-5768d8348d29'::uuid);
INSERT INTO public.movies_casts
(id_movie_cast, id_movie, id_cast)
VALUES('3ca4c18e-689f-433a-8d04-c243106fdf29'::uuid, '60677a43-41ee-4a92-8fbb-d8bc1c4ae770'::uuid, 'cdbc33c1-6de7-4612-8d98-003fe56e86a0'::uuid);
INSERT INTO public.movies_casts
(id_movie_cast, id_movie, id_cast)
VALUES('30de282c-6ff7-48aa-b78c-7deb6964fc0e'::uuid, '063e892d-e9c6-40e5-bb06-d4fb164d6b64'::uuid, '77c71e35-bdc9-4cd5-8ab1-5b961f0e747c'::uuid);
INSERT INTO public.movies_casts
(id_movie_cast, id_movie, id_cast)
VALUES('d59128d9-4f26-40fa-a9f5-17ffdb8d42b8'::uuid, '063e892d-e9c6-40e5-bb06-d4fb164d6b64'::uuid, 'c0abcdb6-9747-49d4-a718-19d9543802b6'::uuid);
INSERT INTO public.movies_casts
(id_movie_cast, id_movie, id_cast)
VALUES('aadb0323-a47b-4c51-9db3-5bc401c84e7e'::uuid, 'a76db5de-c628-4c0c-b897-611e3b382874'::uuid, '77c71e35-bdc9-4cd5-8ab1-5b961f0e747c'::uuid);
INSERT INTO public.movies_casts
(id_movie_cast, id_movie, id_cast)
VALUES('143de2a5-fe90-486e-a5d1-88a87c32b70b'::uuid, 'a76db5de-c628-4c0c-b897-611e3b382874'::uuid, 'f49c44e4-48fa-41ed-a0a9-8109fe137470'::uuid);
INSERT INTO public.movies_casts
(id_movie_cast, id_movie, id_cast)
VALUES('0d8a5a1b-3126-437d-9587-fa93b8dc5cec'::uuid, 'a76db5de-c628-4c0c-b897-611e3b382874'::uuid, '6ec40464-4a19-4ca2-93c9-f2bcaa420169'::uuid);
INSERT INTO public.movies_casts
(id_movie_cast, id_movie, id_cast)
VALUES('4ef31581-4a51-4a14-9dbf-48e2a169b0a5'::uuid, 'ff2a4960-6099-4ef2-b529-c44ce70f8328'::uuid, '250d24f9-1fb8-4a85-920b-304f3864f8a0'::uuid);
INSERT INTO public.movies_casts
(id_movie_cast, id_movie, id_cast)
VALUES('611d3ae1-2c97-461c-83ea-30753ded95a3'::uuid, 'ff2a4960-6099-4ef2-b529-c44ce70f8328'::uuid, '974301e9-9dc4-4ac8-9604-5768d8348d29'::uuid);
INSERT INTO public.movies_casts
(id_movie_cast, id_movie, id_cast)
VALUES('094d9218-fc2a-4fc4-80e6-979fbb402ab8'::uuid, '063e892d-e9c6-40e5-bb06-d4fb164d6b64'::uuid, '23f2828b-ba86-4475-b5bb-ef6cc47eb620'::uuid);
INSERT INTO public.movies_casts
(id_movie_cast, id_movie, id_cast)
VALUES('36b76a63-1686-4ab0-b3ad-ffdd63d1464f'::uuid, 'ff2a4960-6099-4ef2-b529-c44ce70f8328'::uuid, 'efe1c949-cf8a-40ea-82d4-b24065099f96'::uuid);

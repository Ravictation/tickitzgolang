CREATE TABLE public.movies (
	id_movie uuid NOT NULL DEFAULT gen_random_uuid(),
	id_director uuid NULL,
	title varchar(255) NOT NULL,
	release_date date NOT NULL,
	duration_hour int4 NOT NULL,
	duration_minute int4 NOT NULL,
	synopsis text NOT NULL,
	image text NOT NULL,
    cover_image text NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL,
	CONSTRAINT movie_pkey PRIMARY KEY (id_movie),
	CONSTRAINT movie_id_director_fkey FOREIGN KEY (id_director) REFERENCES public.directors(id_director)
);

INSERT INTO public.movies
(id_movie, id_director, title, release_date, duration_hour, duration_minute, synopsis, image, cover_image)
VALUES('ff2a4960-6099-4ef2-b529-c44ce70f8328'::uuid, 'f045c26e-49d9-4e5a-9ec2-ef85bc28714b'::uuid, 'Kimetsu no Yaiba', '2019-04-06', 2, 30, 'Diceritakan ketika umat manusia diteror oleh iblis jahat yang melahap jiwa manu...', 'https://res.cloudinary.com/deaia7unw/image/upload/v1692809165/no-product-image-400x400_hrg7mo.png', 'https://res.cloudinary.com/deaia7unw/image/upload/v1693147474/2981201-sw7_yjcswa.jpg');
INSERT INTO public.movies
(id_movie, id_director, title, release_date, duration_hour, duration_minute, synopsis, image, cover_image)
VALUES('a76db5de-c628-4c0c-b897-611e3b382874'::uuid, 'f045c26e-49d9-4e5a-9ec2-ef85bc28714b'::uuid, 'Mashle', '2023-04-18', 3, 30, 'Magic and Muscles. Ini adalah dunia sihir. Keberadaan ilmu sihi...', 'https://res.cloudinary.com/deaia7unw/image/upload/v1692809165/no-product-image-400x400_hrg7mo.png', 'https://res.cloudinary.com/deaia7unw/image/upload/v1693147474/2981201-sw7_yjcswa.jpg');
INSERT INTO public.movies
(id_movie, id_director, title, release_date, duration_hour, duration_minute, synopsis, image, cover_image)
VALUES('063e892d-e9c6-40e5-bb06-d4fb164d6b64'::uuid, '1009481f-747d-477c-8698-d0e89481a71b'::uuid, 'My Home Hero', '2023-03-31', 2, 30, 'Tetsuo Tosu, seorang pegawai biasa, menemukan putrinya, Reika, telah disik...', 'https://res.cloudinary.com/deaia7unw/image/upload/v1692809165/no-product-image-400x400_hrg7mo.png', 'https://res.cloudinary.com/deaia7unw/image/upload/v1693147474/2981201-sw7_yjcswa.jpg');
INSERT INTO public.movies
(id_movie, id_director, title, release_date, duration_hour, duration_minute, synopsis, image, cover_image)
VALUES('60677a43-41ee-4a92-8fbb-d8bc1c4ae770'::uuid, 'b52a8999-3b9d-455c-99b2-97605cfb4868'::uuid, 'One Piece', '1999-10-19', 2, 30, 'Gol D. Roger dikenal sebagai Raja Bajak Laut, ...', 'https://res.cloudinary.com/deaia7unw/image/upload/v1692809165/no-product-image-400x400_hrg7mo.png', 'https://res.cloudinary.com/deaia7unw/image/upload/v1693147474/2981201-sw7_yjcswa.jpg');

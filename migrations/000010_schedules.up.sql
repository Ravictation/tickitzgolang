CREATE TABLE public.schedules (
	id_schedule uuid NOT NULL DEFAULT gen_random_uuid(),
	id_movie uuid NOT NULL,
	id_premier uuid NOT NULL,
    regency text NOT NULL,
	price int4 NOT NULL,
	set_date date NOT NULL,
	CONSTRAINT schedule_pkey PRIMARY KEY (id_schedule),
	CONSTRAINT schedule_id_movie_fkey FOREIGN KEY (id_movie) REFERENCES public.movies(id_movie),
	CONSTRAINT schedule_id_premier_fkey FOREIGN KEY (id_premier) REFERENCES public.premiers(id_premier)
);

INSERT INTO public.schedules
(id_schedule, id_movie, id_premier, regency, price, set_date)
VALUES('1c0696a8-0e7f-49c8-8d48-339a213d9be8'::uuid, '60677a43-41ee-4a92-8fbb-d8bc1c4ae770'::uuid, 'af09f3dd-dd61-4e96-b9af-3330a630726a'::uuid, 'probolinggo', 35000, '2022-08-25');
INSERT INTO public.schedules
(id_schedule, id_movie, id_premier, regency, price, set_date)
VALUES('f6d12b00-369a-487d-963d-ebaeb85430db'::uuid, '60677a43-41ee-4a92-8fbb-d8bc1c4ae770'::uuid, '2af9789b-b486-4af4-91a3-cadd2c0450cb'::uuid, 'probolinggo', 35000, '2022-08-25');
INSERT INTO public.schedules
(id_schedule, id_movie, id_premier, regency, price, set_date)
VALUES('a3f35226-da92-4a64-824f-4000ff0638f5'::uuid, '60677a43-41ee-4a92-8fbb-d8bc1c4ae770'::uuid, 'af09f3dd-dd61-4e96-b9af-3330a630726a'::uuid, 'surabaya', 35000, '2022-08-25');
INSERT INTO public.schedules
(id_schedule, id_movie, id_premier, regency, price, set_date)
VALUES('b520ff46-503d-4347-8224-0798b461daf0'::uuid, '60677a43-41ee-4a92-8fbb-d8bc1c4ae770'::uuid, '2af9789b-b486-4af4-91a3-cadd2c0450cb'::uuid, 'surabaya', 35000, '2022-08-25');

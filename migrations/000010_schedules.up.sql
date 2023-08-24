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
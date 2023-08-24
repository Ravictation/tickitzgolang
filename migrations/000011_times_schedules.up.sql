CREATE TABLE public.times_schedules (
	id_time_schedule uuid NOT NULL DEFAULT gen_random_uuid(),
	id_schedule uuid NOT NULL,
	time_schedule time NOT NULL,
	CONSTRAINT time_schedule_pkey PRIMARY KEY (id_time_schedule),
	CONSTRAINT time_schedule_id_schedule_fkey FOREIGN KEY (id_schedule) REFERENCES public.schedules(id_schedule)
);
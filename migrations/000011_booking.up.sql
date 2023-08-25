CREATE TABLE public.bookings (
	id_booking uuid NOT NULL DEFAULT gen_random_uuid(),
	id_time_schedule uuid NOT NULL,
	id_user uuid NOT NULL,
	seats _text NOT NULL,
	selected_date date NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL,
	CONSTRAINT booking_pkey PRIMARY KEY (id_booking),
	CONSTRAINT booking_id_time_schedule_fkey FOREIGN KEY (id_time_schedule) REFERENCES public.times_schedules(id_time_schedule),
	CONSTRAINT booking_id_user_fkey FOREIGN KEY (id_user) REFERENCES public.users(id_user)
);

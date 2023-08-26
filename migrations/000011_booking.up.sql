CREATE TABLE public.bookings (
	id_booking uuid NOT NULL DEFAULT gen_random_uuid(),
	id_time_schedule uuid NOT NULL,
	id_user uuid NOT NULL,
	seats text NOT NULL,
	total int NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL,
	CONSTRAINT booking_pkey PRIMARY KEY (id_booking),
	CONSTRAINT booking_id_time_schedule_fkey FOREIGN KEY (id_time_schedule) REFERENCES public.times_schedules(id_time_schedule),
	CONSTRAINT booking_id_user_fkey FOREIGN KEY (id_user) REFERENCES public.users(id_user)
);

INSERT INTO public.bookings
(id_booking, id_time_schedule, id_user, seats, total, created_at, updated_at)
VALUES('5c36fb94-dae3-4b30-90d4-a7745f4c6151'::uuid, '85435820-fc0d-4559-9129-a4c8431b4813'::uuid, '0ae1db2e-c038-42da-9530-11918a0537fb'::uuid, 'A1,A2', 70000, '2023-08-25 20:13:17.013', NULL);


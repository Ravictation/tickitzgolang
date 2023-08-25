CREATE TABLE public.booking (
	booking_id uuid NOT NULL DEFAULT gen_random_uuid(),
	user_id uuid NOT NULL,
	booking_date date NULL,
	seats _varchar NULL,
	price numeric(10, 2) NULL,
	CONSTRAINT booking_pkey PRIMARY KEY (booking_id)
);

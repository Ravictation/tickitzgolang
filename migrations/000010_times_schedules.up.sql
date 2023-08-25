CREATE TABLE public.times_schedules (
	id_time_schedule uuid NOT NULL DEFAULT gen_random_uuid(),
	id_schedule uuid NOT NULL,
	time_schedule time NOT NULL,
	CONSTRAINT time_schedule_pkey PRIMARY KEY (id_time_schedule),
	CONSTRAINT time_schedule_id_schedule_fkey FOREIGN KEY (id_schedule) REFERENCES public.schedules(id_schedule)
);

INSERT INTO public.times_schedules
(id_time_schedule, id_schedule, time_schedule)
VALUES('45614fd5-b96d-48c9-b0f7-a8f8ebea8a22'::uuid, '1c0696a8-0e7f-49c8-8d48-339a213d9be8'::uuid, '10:00:00');
INSERT INTO public.times_schedules
(id_time_schedule, id_schedule, time_schedule)
VALUES('85435820-fc0d-4559-9129-a4c8431b4813'::uuid, '1c0696a8-0e7f-49c8-8d48-339a213d9be8'::uuid, '14:00:00');
INSERT INTO public.times_schedules
(id_time_schedule, id_schedule, time_schedule)
VALUES('57df7db0-8673-4b46-874d-43fe72f2f9c0'::uuid, 'f6d12b00-369a-487d-963d-ebaeb85430db'::uuid, '10:00:00');
INSERT INTO public.times_schedules
(id_time_schedule, id_schedule, time_schedule)
VALUES('7ab6e7e0-7cde-4d2a-a330-4d5df28b2f8f'::uuid, 'f6d12b00-369a-487d-963d-ebaeb85430db'::uuid, '14:00:00');
INSERT INTO public.times_schedules
(id_time_schedule, id_schedule, time_schedule)
VALUES('e9404242-ccfd-406a-98cc-2f9c05e9b569'::uuid, 'a3f35226-da92-4a64-824f-4000ff0638f5'::uuid, '10:00:00');
INSERT INTO public.times_schedules
(id_time_schedule, id_schedule, time_schedule)
VALUES('8e548f65-7028-43ec-b4af-a65339079bfb'::uuid, 'a3f35226-da92-4a64-824f-4000ff0638f5'::uuid, '14:00:00');
INSERT INTO public.times_schedules
(id_time_schedule, id_schedule, time_schedule)
VALUES('99041bb6-1bd8-463c-afbf-eb45b1e8b8d3'::uuid, 'b520ff46-503d-4347-8224-0798b461daf0'::uuid, '10:00:00');
INSERT INTO public.times_schedules
(id_time_schedule, id_schedule, time_schedule)
VALUES('6181f332-8aa3-46a8-818b-0b83cb4404e4'::uuid, 'b520ff46-503d-4347-8224-0798b461daf0'::uuid, '14:00:00');

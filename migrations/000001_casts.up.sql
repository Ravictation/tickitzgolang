CREATE TABLE public.casts (
	id_cast uuid NOT NULL DEFAULT gen_random_uuid(),
	name_cast varchar(255) NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL,
	CONSTRAINT casts_pkey PRIMARY KEY (id_cast)
);

INSERT INTO public.casts
(id_cast, name_cast)
VALUES('cdbc33c1-6de7-4612-8d98-003fe56e86a0'::uuid, 'Mayumi Tanaka');
INSERT INTO public.casts
(id_cast, name_cast)
VALUES('6ec40464-4a19-4ca2-93c9-f2bcaa420169'::uuid, 'Kazuya Nakai');
INSERT INTO public.casts
(id_cast, name_cast)
VALUES('efe1c949-cf8a-40ea-82d4-b24065099f96'::uuid, 'Akemi Okamura');
INSERT INTO public.casts
(id_cast, name_cast)
VALUES('250d24f9-1fb8-4a85-920b-304f3864f8a0'::uuid, 'Inori Minase');
INSERT INTO public.casts
(id_cast, name_cast)
VALUES('80c344ae-338e-4b50-9957-5d74183f7950'::uuid, 'Yuuki Sakakihara');
INSERT INTO public.casts
(id_cast, name_cast)
VALUES('c44354a1-6e97-4080-8335-ea64c951e6d2'::uuid, 'Yuuma Uchida');
INSERT INTO public.casts
(id_cast, name_cast)
VALUES('4a6d59af-75e9-450b-b3f6-2560a7341a26'::uuid, 'Kento Itou');
INSERT INTO public.casts
(id_cast, name_cast)
VALUES('f49c44e4-48fa-41ed-a0a9-8109fe137470'::uuid, 'Junichi Suwabe');
INSERT INTO public.casts
(id_cast, name_cast)
VALUES('23f2828b-ba86-4475-b5bb-ef6cc47eb620'::uuid, 'Rumi Ookubo');
INSERT INTO public.casts
(id_cast, name_cast)
VALUES('34390e0d-4226-41e5-94a2-3a08cc8f13de'::uuid, 'Chiaki Kobayashi');
INSERT INTO public.casts
(id_cast, name_cast)
VALUES('77c71e35-bdc9-4cd5-8ab1-5b961f0e747c'::uuid, 'Reina Ueda');
INSERT INTO public.casts
(id_cast, name_cast)
VALUES('e3271aad-6f57-4480-8948-e7fd3be99b9a'::uuid, 'Kaito Ishikawa');
INSERT INTO public.casts
(id_cast, name_cast)
VALUES('c0abcdb6-9747-49d4-a718-19d9543802b6'::uuid, 'Natsuki Hanae');
INSERT INTO public.casts
(id_cast, name_cast)
VALUES('974301e9-9dc4-4ac8-9604-5768d8348d29'::uuid, 'Akari Kitou');
INSERT INTO public.casts
(id_cast, name_cast)
VALUES('7ff38402-a1e3-428b-ae34-af35e25739bc'::uuid, 'Hiro Shimono');


CREATE TABLE public.premiers (
	id_premier uuid NOT NULL DEFAULT gen_random_uuid(),
	name_premier varchar(255) NOT NULL,
	image text NOT NULL,
	count_row_seat int4 NOT NULL,
	count_col_seat int4 NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL,
	CONSTRAINT premier_pkey PRIMARY KEY (id_premier)
);

INSERT INTO public.premiers
(id_premier, name_premier, image, count_row_seat, count_col_seat, created_at, updated_at)
VALUES('af09f3dd-dd61-4e96-b9af-3330a630726a'::uuid, 'cineone21', 'http://res.cloudinary.com/deaia7unw/image/upload/v1692894314/yunfjstaayucyzd9kele.png', 10, 5, '2023-08-24 23:25:14.090', NULL);
INSERT INTO public.premiers
(id_premier, name_premier, image, count_row_seat, count_col_seat, created_at, updated_at)
VALUES('2af9789b-b486-4af4-91a3-cadd2c0450cb'::uuid, 'ebv.id', 'http://res.cloudinary.com/deaia7unw/image/upload/v1692894435/kcgiravpitbdxzpiqdkz.png', 10, 5, '2023-08-24 23:27:15.000', NULL);
INSERT INTO public.premiers
(id_premier, name_premier, image, count_row_seat, count_col_seat, created_at, updated_at)
VALUES('e2b957ec-7f4c-4fdb-be03-82ff72085f16'::uuid, 'hiflix', 'http://res.cloudinary.com/deaia7unw/image/upload/v1692894512/szcxljsryrjrdwonyjai.png', 10, 5, '2023-08-24 23:28:32.059', '2023-08-24 23:30:21.918');

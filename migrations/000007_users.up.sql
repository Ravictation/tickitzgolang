CREATE TABLE public.users (
	id_user uuid NOT NULL DEFAULT gen_random_uuid(),
	email_user varchar(255) NULL,
	"password" varchar(255) NOT NULL,
	phone_number varchar(20) NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL,
	image_user text NULL,
	"role" varchar(50) NULL,
	first_name varchar NULL,
	last_name varchar NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id_user)
);

INSERT INTO public.users
(id_user, email_user, "password", phone_number, created_at, updated_at, image_user, "role", first_name, last_name)
VALUES('301c72b4-6bd6-415b-8d49-abb86cae860d'::uuid, 'ddada@gmail.com', '$2a$10$4kvdwB3TlxxI76975Sk4LOpxDvoVGn68yAdVw2chF.5.Hc2e4dZTC', '', '2023-08-25 14:17:44.256', NULL, 'https://res.cloudinary.com/deaia7unw/image/upload/v1691956997/s8owthmz0dlwzvstolmc.png', 'user', '', '');
INSERT INTO public.users
(id_user, email_user, "password", phone_number, created_at, updated_at, image_user, "role", first_name, last_name)
VALUES('0f30cafd-752a-4852-8795-926167ddbdad'::uuid, 'iniadmin@gmail.com', '$2a$10$m2agtMk0bY0QXJNnJ03gye.f3Bu487Xdmr9B6OSj.pU8gmEc5NUZW', '', '2023-08-25 15:20:59.353', NULL, 'https://res.cloudinary.com/deaia7unw/image/upload/v1691956997/s8owthmz0dlwzvstolmc.png', 'admin', '', '');
INSERT INTO public.users
(id_user, email_user, "password", phone_number, created_at, updated_at, image_user, "role", first_name, last_name)
VALUES('0ae1db2e-c038-42da-9530-11918a0537fb'::uuid, 'iniadmin1@gmail.com', '$2a$10$f2SJo5PFn.cDqsjkFwiOV.T2TY1mycNolECrI3.i9vk37sOmCg0Oe', '', '2023-08-25 15:21:02.959', NULL, 'https://res.cloudinary.com/deaia7unw/image/upload/v1691956997/s8owthmz0dlwzvstolmc.png', 'admin', '', '');
INSERT INTO public.users
(id_user, email_user, "password", phone_number, created_at, updated_at, image_user, "role", first_name, last_name)
VALUES('19f43646-f7e4-4223-895e-6f33f8776762'::uuid, 'iniadmin2@gmail.com', '$2a$10$PY3plnkhKbVj1.6eMExx8OBMAV4eE13RHB/aEPd9zjZflV36BAZ7y', '', '2023-08-25 15:21:08.587', NULL, 'https://res.cloudinary.com/deaia7unw/image/upload/v1691956997/s8owthmz0dlwzvstolmc.png', 'admin', '', '');
INSERT INTO public.users
(id_user, email_user, "password", phone_number, created_at, updated_at, image_user, "role", first_name, last_name)
VALUES('fddd8e69-a956-458a-9901-691435fbc1bd'::uuid, 'iniuser@gmail.com', '$2a$10$UdCQsAPqtRurn0tFnCor2uzQcF45AfFVFTyUhp.J1gOBshhG5TWLy', '', '2023-08-25 15:22:17.646', NULL, 'https://res.cloudinary.com/deaia7unw/image/upload/v1691956997/s8owthmz0dlwzvstolmc.png', 'user', '', '');
INSERT INTO public.users
(id_user, email_user, "password", phone_number, created_at, updated_at, image_user, "role", first_name, last_name)
VALUES('701762a8-8309-45ec-8ad2-5283551c703f'::uuid, 'iniuser1@gmail.com', '$2a$10$VoIUf09kctSYl.ORpJtGQeYcTXZqIeH8O3wnTDKlYd.gmMhGyN3Gm', '', '2023-08-25 15:22:19.795', NULL, 'https://res.cloudinary.com/deaia7unw/image/upload/v1691956997/s8owthmz0dlwzvstolmc.png', 'user', '', '');

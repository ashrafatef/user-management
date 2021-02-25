-- CREATE TABLE public.organization_roles
-- (
--     id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
--     name character varying(50) COLLATE pg_catalog."default" NOT NULL,
--     description character varying(50) COLLATE pg_catalog."default",
--     organization_id integer NOT NULL,
--     created_at timestamp with time zone,
--     updated_at timestamp with time zone,
--     CONSTRAINT organization_roles_pkey PRIMARY KEY (id),
--     CONSTRAINT organization_id FOREIGN KEY (id)
--         REFERENCES public.organizations (id) MATCH SIMPLE
--         ON UPDATE CASCADE
--         ON DELETE CASCADE
--         NOT VALID
-- )

-- TABLESPACE pg_default;

-- ALTER TABLE public.organization_roles
--     OWNER to botter;
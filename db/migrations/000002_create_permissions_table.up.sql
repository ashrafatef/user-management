CREATE TYPE public.permission_type AS ENUM
    ('read', 'create', 'update', 'delete');

CREATE TYPE public.permission_category AS ENUM
    ('builder', 'nlp', 'wordspotting' ,'insights','settings','administration','roles','users','bots');

ALTER TYPE public.permission_type
    OWNER TO postgres;
    ALTER TYPE public.permission_category
    OWNER TO postgres;

CREATE TABLE public.permissions
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    name character varying COLLATE pg_catalog."default" NOT NULL,
    category permission_category NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    type permission_type NOT NULL,
    CONSTRAINT permissions_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.permissions
    OWNER to botter;
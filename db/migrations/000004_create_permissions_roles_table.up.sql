CREATE TABLE public.permissions_roles
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    role_id bigint NOT NULL,
    permission_id bigint NOT NULL,
    created_at time with time zone NOT NULL DEFAULT now(),
    updated_at time with time zone NOT NULL DEFAULT now(),
    CONSTRAINT permissions_roles_pkey PRIMARY KEY (id),
    CONSTRAINT permissions FOREIGN KEY (id)
        REFERENCES public.permissions (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT roles FOREIGN KEY (id)
        REFERENCES public.organization_roles (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE public.permissions_roles
    OWNER to botter;
-- Index: fki_permissions

-- DROP INDEX public.fki_permissions;

CREATE INDEX fki_permissions
    ON public.permissions_roles USING btree
    (permission_id ASC NULLS LAST)
    TABLESPACE pg_default;
-- Index: fki_roles

-- DROP INDEX public.fki_roles;

CREATE INDEX fki_roles
    ON public.permissions_roles USING btree
    (role_id ASC NULLS LAST)
    TABLESPACE pg_default;
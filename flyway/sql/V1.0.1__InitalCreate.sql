CREATE TABLE public.project
(
    id serial,
    created_at timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp with time zone NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id)
);

CREATE TABLE public.version
(
    id serial NOT NULL,
    title character varying(30) NOT NULL,
    description character varying(300) NOT NULL,
    project_id integer NOT NULL,
    PRIMARY KEY(id)
)
    INHERITS (public.project);

ALTER TABLE public.version
    ADD CONSTRAINT fk_project FOREIGN KEY (project_id)
    REFERENCES public.project (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

CREATE TABLE public.project_information
(
    id serial,
    order_id integer NOT NULL,
    name character varying(45) NOT NULL,
    description character varying(300),
    manager character varying(45) NOT NULL,
    client character varying(45) NOT NULL,
    sector character varying(45) NOT NULL,
    project_id integer NOT NULL,
    version_id integer NOT NULL,
    PRIMARY KEY(id)
)
    INHERITS (public.project);

ALTER TABLE public.project_information
    ADD CONSTRAINT fk_project FOREIGN KEY (project_id)
    REFERENCES public.project (id) MATCH SIMPLE
    ON UPDATE RESTRICT
    ON DELETE CASCADE
    NOT VALID;

ALTER TABLE public.project_information
    ADD CONSTRAINT fk_version FOREIGN KEY (version_id)
    REFERENCES public.version (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

CREATE TABLE public.components
(
    id serial,
    component_id integer NOT NULL,
    name character varying(45) NOT NULL,
    profile character varying(40) NOT NULL,
    material character varying(40) NOT NULL,
    co real NOT NULL,
    level int NOT NULL,
    type character varying(45) NOT NULL,
    version_id integer NOT NULL,
    PRIMARY KEY(id)
)
    INHERITS (public.project);

ALTER TABLE public.components
    ADD CONSTRAINT pk_version FOREIGN KEY (version_id)
    REFERENCES public.version (id) MATCH SIMPLE
    ON UPDATE RESTRICT
    ON DELETE CASCADE
    NOT VALID;




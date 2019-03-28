-- 用户表
CREATE TABLE public.users
(
    user_id bigint NOT NULL DEFAULT nextval('users_user_id_seq'::regclass),
    user_name character varying COLLATE pg_catalog."default" NOT NULL,
    pass_word character varying COLLATE pg_catalog."default" NOT NULL,
    create_date timestamp without time zone NOT NULL DEFAULT now(),
    CONSTRAINT users_pkey PRIMARY KEY (user_id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.users
    OWNER to postgres;

-- 文件表
CREATE TABLE public.files
(
    file_id bigint NOT NULL DEFAULT nextval('files_file_id_seq'::regclass),
    file_name character varying COLLATE pg_catalog."default" NOT NULL,
    create_date timestamp without time zone NOT NULL DEFAULT now(),
    file_data bytea NOT NULL,
    CONSTRAINT files_pkey PRIMARY KEY (file_id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.files
    OWNER to postgres;
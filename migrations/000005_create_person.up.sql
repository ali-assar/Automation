CREATE TABLE IF NOT EXISTS person
(
    national_id_number character varying(255) NOT NULL,
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL,
    family_info_id bigint,
    contact_info_id bigint,
    skills_id bigint,
    physical_info_id bigint,
    birth_date date NOT NULL,
    religion_id bigint,
    person_type_id bigint,
    military_details_id bigint,
    deleted_at bigint NOT NULL,
    CONSTRAINT person_pkey PRIMARY KEY (national_id_number)
);
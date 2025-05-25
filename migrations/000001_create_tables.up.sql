CREATE TABLE actiontype (
    id BIGINT PRIMARY KEY,
    action_name VARCHAR(255) NOT NULL,
    deleted_at BIGINT NOT NULL
);

CREATE TABLE bloodgroup (
    id BIGINT PRIMARY KEY,
    "group" VARCHAR(255) NOT NULL UNIQUE,
    deleted_at BIGINT NOT NULL
);

CREATE TABLE gender (
    id BIGINT PRIMARY KEY,
    gender VARCHAR(255) NOT NULL,
    deleted_at BIGINT NOT NULL
);

CREATE TABLE persontype (
    id BIGINT PRIMARY KEY,
    type VARCHAR(255) NOT NULL,
    deleted_at BIGINT NOT NULL
);

CREATE TABLE physicalstatus (
    id BIGINT PRIMARY KEY,
    status VARCHAR(255) NOT NULL,
    description JSON NOT NULL,
    deleted_at BIGINT NOT NULL
);

CREATE TABLE rank (
    id BIGINT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    deleted_at BIGINT NOT NULL
);

CREATE TABLE religion (
    id BIGINT PRIMARY KEY,
    religion_name VARCHAR(255) NOT NULL,
    religion_type VARCHAR(255) NOT NULL,
    deleted_at BIGINT NOT NULL
);

CREATE TABLE role (
    id BIGINT PRIMARY KEY,
    type VARCHAR(255) NOT NULL,
    deleted_at BIGINT NOT NULL
);

CREATE TABLE education (
    id BIGINT PRIMARY KEY,
    education_level_id BIGINT NOT NULL,
    field_of_study BIGINT NOT NULL,
    description VARCHAR(255) NOT NULL,
    university VARCHAR(255) NOT NULL,
    start_date BIGINT NOT NULL,
    end_date BIGINT NOT NULL,
    deleted_at BIGINT NOT NULL
);

CREATE TABLE skills (
    id BIGINT PRIMARY KEY,
    education_id BIGINT NOT NULL REFERENCES education(id),
    languages JSON NOT NULL,
    skills_description TEXT NOT NULL,
    certificates TEXT NOT NULL,
    deleted_at BIGINT NOT NULL
);

CREATE TABLE physicalinfo (
    id BIGINT PRIMARY KEY,
    height INT NOT NULL,
    weight INT NOT NULL,
    eye_color VARCHAR(255) NOT NULL,
    blood_group_id BIGINT NOT NULL REFERENCES bloodgroup(id),
    gender_id BIGINT NOT NULL REFERENCES gender(id),
    physical_status_id BIGINT NOT NULL REFERENCES physicalstatus(id),
    deleted_at BIGINT NOT NULL
);

CREATE TABLE militarydetails (
    id BIGINT PRIMARY KEY,
    rank_id BIGINT NOT NULL REFERENCES rank(id),
    service_start_date BIGINT NOT NULL,
    service_dispatch_date BIGINT NOT NULL,
    service_unit BIGINT NOT NULL,
    battalion_unit BIGINT NOT NULL,
    company_unit BIGINT NOT NULL,
    deleted_at BIGINT NOT NULL
);

CREATE TABLE familyinfo (
    id BIGINT PRIMARY KEY,
    father_details JSON NOT NULL,
    mother_details JSON NOT NULL,
    childs_details JSON NOT NULL,
    husband_details JSON NOT NULL,
    deleted_at BIGINT NOT NULL
);

CREATE TABLE contactinfo (
    id BIGINT PRIMARY KEY,
    address VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) NOT NULL,
    emergency_phone_number VARCHAR(255) NOT NULL,
    landline_phone VARCHAR(255) NOT NULL,
    email_address VARCHAR(255) NOT NULL,
    social_media JSON NOT NULL,
    deleted_at BIGINT NOT NULL
);

CREATE TABLE person (
    national_id_number VARCHAR(255) PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    family_info_id BIGINT NOT NULL REFERENCES familyinfo(id),
    physical_info_id BIGINT NOT NULL REFERENCES physicalinfo(id),
    contact_info_id BIGINT NOT NULL REFERENCES contactinfo(id),
    skills_id BIGINT NOT NULL REFERENCES skills(id),
    birth_date DATE NOT NULL,
    religion_id BIGINT NOT NULL REFERENCES religion(id),
    person_type_id BIGINT NOT NULL REFERENCES persontype(id),
    military_details_id BIGINT NOT NULL REFERENCES militarydetails(id),
    deleted_at BIGINT NOT NULL
);

CREATE TABLE admin (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    national_id_number VARCHAR(255) NOT NULL REFERENCES person(national_id_number),
    user_name VARCHAR(255) NOT NULL,
    hash_password VARCHAR(255) NOT NULL,
    role_id BIGINT NOT NULL REFERENCES role(id),
    deleted_at BIGINT NOT NULL
);

CREATE TABLE credentials (
    id SERIAL PRIMARY KEY,
    admin_id UUID NOT NULL REFERENCES admin(id) UNIQUE,
    dynamic_token VARCHAR(255),
    static_token VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);
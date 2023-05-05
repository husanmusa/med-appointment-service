create table if not exists appointment
(
    id           uuid primary key not null,
    appoint_id   int              not null,
    doctor_id    uuid             not null,
    patient_id   uuid             not null,
    appoint_date varchar,
    appoint_time varchar,
    status       int,
    created_at   timestamp default current_timestamp,
    updated_at   timestamp default current_timestamp,
    unique ( doctor_id, appoint_date, appoint_time)
);

create table if not exists doctor
(
    id         uuid primary key not null,
    name       varchar,
    age        int,
    role       varchar,
    polyclinic varchar,
    gender     int,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

create table if not exists patient
(
    id            uuid primary key not null,
    name          varchar,
    age           int,
    gender        int,
    login         varchar,
    password      varchar,
    access_token  text,
    refresh_token text,
    created_at    timestamp default current_timestamp,
    updated_at    timestamp default current_timestamp
);
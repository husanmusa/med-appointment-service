create table appointment
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
CREATE TABLE IF NOT EXISTS groups_schedule
(
    id          SERIAL PRIMARY KEY,
    group_name  VARCHAR(50) NOT NULL UNIQUE,
    update_time DATE        NOT NULL,
    info        jsonb       NOT NULL
);

CREATE TABLE IF NOT EXISTS teachers_schedule
(
    id           SERIAL PRIMARY KEY,
    teacher_name VARCHAR(50) NOT NULL UNIQUE,
    update_time  DATE        NOT NULL,
    info         jsonb       NOT NULL
);
CREATE TABLE IF NOT EXISTS groups_schedules
(
    id                      SERIAL PRIMARY KEY,
    group_name              VARCHAR(50) NOT NULL UNIQUE,
    first_week_update_time  TIMESTAMP   NOT NULL,
    second_week_update_time TIMESTAMP   NOT NULL,
    full_schedule           jsonb       NOT NULL
);

CREATE TABLE IF NOT EXISTS teachers_schedules
(
    id                      SERIAL PRIMARY KEY,
    teacher_name            VARCHAR(50) NOT NULL UNIQUE,
    first_week_update_time  TIMESTAMP   NOT NULL,
    second_week_update_time TIMESTAMP   NOT NULL,
    full_schedule           jsonb       NOT NULL
);
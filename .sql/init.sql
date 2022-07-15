create schema if not exists processing;
use processing;
CREATE TABLE persons(  
    person_id int NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'Primary Key'
);

create table terminal_alias
(
    terminal_id     int null,
    dst_terminal_id int null
);

create index IDX_DST_ID
    on terminal_alias (dst_terminal_id);

create table solar_panels_data
(
    person_id           int      not null,
    date                datetime not null,
    battery_voltage     float    null,
    solar_voltage       float    null,
    solar_current       float    null,
    load_current        float    null,
    primary key (person_id, date)
);
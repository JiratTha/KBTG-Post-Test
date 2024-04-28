DROP SCHEMA IF EXISTS project1 CASCADE;
CREATE SCHEMA project1;
DROP TABLE IF EXISTS project1.allowance;
create table project1.allowance
(
    allowance_type varchar not null
        constraint allowance_pk
            primary key,
    amount         double precision
);

alter table project1.allowance
    owner to postgres;

INSERT INTO project1.allowance (allowance_type, amount) VALUES ('donation', 100000);
INSERT INTO project1.allowance (allowance_type, amount) VALUES ('k-receipt', 50000);


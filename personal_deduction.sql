DROP TABLE IF EXISTS project1.personal_deduction;
create table project1.personal_deduction
(
    personal_deduction varchar not null
        constraint personal_deduction_pk
            primary key,
    amount              double precision
);

alter table project1.personal_deduction
    owner to postgres;

INSERT INTO project1.personal_deduction (personal_deduction, amount) VALUES ('personalDeduction', 60000);
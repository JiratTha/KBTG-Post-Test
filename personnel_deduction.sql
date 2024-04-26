DROP TABLE IF EXISTS project1.personnel_deduction;
create table project1.personnel_deduction
(
    personnel_deduction varchar not null
        constraint personnel_deduction_pk
            primary key,
    amount              double precision
);

alter table project1.personnel_deduction
    owner to postgres;

INSERT INTO project1.personnel_deduction (personnel_deduction, amount) VALUES ('personnelDeduction', 60000);

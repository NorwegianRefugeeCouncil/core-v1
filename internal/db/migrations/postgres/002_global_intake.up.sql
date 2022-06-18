alter table individuals
    add column preferred_name varchar(255) not null default '';
alter table individuals
    add column is_minor bool not null default false;
alter table individuals
    add column presents_protection_concerns bool not null default false;
alter table individuals
    add column physical_impairment varchar(32) not null default '';
alter table individuals
    add column sensory_impairment varchar(32) not null default '';
alter table individuals
    add column mental_impairment varchar(32) not null default '';
alter table individuals
    add column displacement_status varchar(64) not null default '';

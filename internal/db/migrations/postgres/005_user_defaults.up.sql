alter table individuals
    alter column full_name set not null;

alter table individuals
    alter column full_name set default '';

alter table individuals
    alter column phone_number set not null;

alter table individuals
    alter column phone_number set default '';

alter table individuals
    alter column normalized_phone_number set not null;

alter table individuals
    alter column normalized_phone_number set default '';

alter table individuals
    alter column email set not null;

alter table individuals
    alter column email set default '';

alter table individuals
    alter column address set not null;

alter table individuals
    alter column address set default '';

alter table individuals
    alter column gender set not null;

alter table individuals
    alter column gender set default '';


create table if not exists customers (
       id            bigserial          not null unique,
       name          varchar(20)        not null,
       address1      varchar(50)        not null,
       address2      varchar(50)        not null,       
       address3      varchar(50)        not null,
       created_at    timestamp          not null,
       updated_at    timestamp          not null,
       primary key (id)
);

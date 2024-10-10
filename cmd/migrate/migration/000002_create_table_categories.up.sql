create table categories (
    id serial primary key,
    name varchar(100) unique
);

-- insert some dummy data
insert into categories(name)
    values ('Travel'), ('Entertainment'), ('Music'), ('Cars'), ('Lifestyle'), ('Movie');
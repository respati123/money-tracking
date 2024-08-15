CREATE TABLE roles (
    id serial4 not null primary key ,
    uuid UUID NOT NULL Default uuid_generate_v4(),
    role_code int NOT NULL,
    name VARCHAR(255) NOT NULL,
    alias VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by int null,
    updated_by int null,
    deleted_by int null
);

alter table users add column role_code int null;

alter table roles add constraint role_code UNIQUE (role_code);
alter table users add constraint users_role_code_fkey FOREIGN KEY (role_code) REFERENCES roles(role_code);
alter table roles add constraint role_name UNIQUE (name);

insert into roles (role_code, name, alias, created_at, created_by)
values (1000, 'superadmin', 'superadmin', now(), 5000);


insert INTO users (user_code, email, password, phone_number, is_active, created_at, role_code)
VALUES (5000, 'tyorespati@gmail.com', '$2a$14$Rydml1m0r2tipLxzVtiu5.stEznOVeC0pDEddpPMIIJjDrN5A7bzu', '082143234019',
        true, now(), 1000);

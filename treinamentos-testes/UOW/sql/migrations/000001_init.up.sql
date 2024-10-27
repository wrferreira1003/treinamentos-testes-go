CREATE TABLE categories (
    id varchar(36) not null primary key,
    name text NOT NULL,
    description text
);

CREATE TABLE courses (
    id varchar(36) not null primary key,
    category_id varchar(36) not null,
    name text NOT NULL,
    description text,
    thumbnail text,
    price decimal(10,2) not null,
    foreign key (category_id) references categories (id)
);


create table generatedproductsubcategories
(
    id          int auto_increment
        primary key,
    name        varchar(255)                        null,
    category_id int                                 null,
    created_at  timestamp default CURRENT_TIMESTAMP not null,
    updated_at  timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
);


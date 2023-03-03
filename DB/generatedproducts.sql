create table generatedproducts
(
    id          int auto_increment
        primary key,
    name        varchar(255)                        null,
    description text                                null,
    price       decimal(10, 2)                      null,
    image       varchar(255)                        null,
    category    varchar(255)                        null,
    subcategory varchar(255)                        null,
    created_at  timestamp default CURRENT_TIMESTAMP not null,
    updated_at  timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
);

INSERT INTO aigen.generatedproducts (id, name, description, price, image, category, subcategory, created_at, updated_at) VALUES (1, 'Test', 'Hello World', 0.00, 'generatedProducts/AIGEN_hXtSf.jpg', 'Design', 'Shirts', '2023-02-23 21:08:54', '2023-02-23 21:08:54');
INSERT INTO aigen.generatedproducts (id, name, description, price, image, category, subcategory, created_at, updated_at) VALUES (2, 'Test', 'Programmers roasting', 0.00, 'generatedProducts/AIGEN_JWvaD.jpg', 'Design', 'Shirts', '2023-02-23 21:11:20', '2023-02-23 21:11:20');
INSERT INTO aigen.generatedproducts (id, name, description, price, image, category, subcategory, created_at, updated_at) VALUES (3, 'Test', 'Programmers roasting', 0.00, 'generatedProducts/AIGEN_fCRYw.jpg', 'Design', 'Shirts', '2023-02-23 22:16:57', '2023-02-23 22:16:57');
INSERT INTO aigen.generatedproducts (id, name, description, price, image, category, subcategory, created_at, updated_at) VALUES (4, 'Test', 'Programmers roasting', 0.00, 'generatedProducts/AIGEN_zYuWA.jpg', 'Design', 'Shirts', '2023-02-23 22:16:59', '2023-02-23 22:16:59');
INSERT INTO aigen.generatedproducts (id, name, description, price, image, category, subcategory, created_at, updated_at) VALUES (5, 'Test', 'Programmers roasting', 0.00, 'generatedProducts/AIGEN_mujQE.jpg', 'Design', 'Shirts', '2023-02-23 22:17:01', '2023-02-23 22:17:01');
INSERT INTO aigen.generatedproducts (id, name, description, price, image, category, subcategory, created_at, updated_at) VALUES (6, 'Test', 'Programmers roasting', 0.00, 'generatedProducts/AIGEN_nXXJt.jpg', 'Design', 'Shirts', '2023-02-23 22:22:54', '2023-02-23 22:22:54');
INSERT INTO aigen.generatedproducts (id, name, description, price, image, category, subcategory, created_at, updated_at) VALUES (7, 'Test', 'Programmers roasting', 0.00, 'generatedProducts/AIGEN_WgzsM.jpg', 'Design', 'Shirts', '2023-02-23 22:22:56', '2023-02-23 22:22:56');
INSERT INTO aigen.generatedproducts (id, name, description, price, image, category, subcategory, created_at, updated_at) VALUES (8, 'Test', 'Programmers roasting', 0.00, 'generatedProducts/AIGEN_Ggrdy.jpg', 'Design', 'Shirts', '2023-02-23 22:22:59', '2023-02-23 22:22:59');
INSERT INTO aigen.generatedproducts (id, name, description, price, image, category, subcategory, created_at, updated_at) VALUES (9, 'Test', 'Running on mars', 0.00, 'generatedProducts/AIGEN_EEGwA.jpg', 'Design', 'Shirts', '2023-02-23 23:21:21', '2023-02-23 23:21:21');
INSERT INTO aigen.generatedproducts (id, name, description, price, image, category, subcategory, created_at, updated_at) VALUES (10, 'Test', 'Running on mars', 0.00, 'generatedProducts/AIGEN_rLbAn.jpg', 'Design', 'Shirts', '2023-02-23 23:21:25', '2023-02-23 23:21:25');
INSERT INTO aigen.generatedproducts (id, name, description, price, image, category, subcategory, created_at, updated_at) VALUES (11, 'Test', 'Running on mars', 0.00, 'generatedProducts/AIGEN_bupPa.jpg', 'Design', 'Shirts', '2023-02-23 23:21:28', '2023-02-23 23:21:28');
INSERT INTO aigen.generatedproducts (id, name, description, price, image, category, subcategory, created_at, updated_at) VALUES (12, 'Test', 'Programmers roasting', 0.00, 'generatedProducts/AIGEN_zKHaP.jpg', 'Design', 'Shirts', '2023-02-23 23:39:41', '2023-02-23 23:39:41');
INSERT INTO aigen.generatedproducts (id, name, description, price, image, category, subcategory, created_at, updated_at) VALUES (13, 'Test', 'Programmers roasting', 0.00, 'generatedProducts/AIGEN_ehdYO.jpg', 'Design', 'Shirts', '2023-02-23 23:39:45', '2023-02-23 23:39:45');
INSERT INTO aigen.generatedproducts (id, name, description, price, image, category, subcategory, created_at, updated_at) VALUES (14, 'Test', 'Programmers roasting', 0.00, 'generatedProducts/AIGEN_MMZVP.jpg', 'Design', 'Shirts', '2023-02-23 23:39:48', '2023-02-23 23:39:48');

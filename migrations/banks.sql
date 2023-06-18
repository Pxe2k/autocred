INSERT INTO banks (title, id) VALUES ('МФО Orbis Finance', 1);
INSERT INTO banks (title, id) VALUES ('Банк Центр Кредит', 1);
INSERT INTO banks (title, id) VALUES ('Евразийский Банк', 2);
INSERT INTO banks (title, id) VALUES ('Шинхан Банк', 3);
INSERT INTO banks (title) VALUES ('Береке Банк');

INSERT INTO bank_products (title, bank_id) VALUES ('С пробегом_Аннуитет', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('С пробегом_Дифф_6М', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('С пробегом_Дифф_12М', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('С пробегом_Дифф_24М', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('С пробегом_Дифф_36М', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('С пробегом_Дифф_48М', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('С пробегом_Дифф_60М', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('С пробегом_Дифф_72М', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('С пробегом_Дифф_84М', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('С пробегом_Аннуитет_ПВ10%', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('С пробегом_Дифф_6М_ПВ10%', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('С пробегом_Дифф_12М_ПВ10%', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('С пробегом_Дифф_24М_ПВ10%', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('С пробегом_Дифф_36М_ПВ10%', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('С пробегом_Дифф_48М_ПВ10%', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('С пробегом_Дифф_60М_ПВ10%', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('С пробегом_Дифф_72М_ПВ10%', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('С пробегом_Дифф_84М_ПВ10%', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('Новое_Аннуитет', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('Новое_Дифф_6М', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('Новое_Дифф_12М', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('Новое_Дифф_24М', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('Новое_Дифф_36М', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('Новое_Дифф_48М', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('Новое_Дифф_60М', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('Новое_Дифф_72М', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('Новое_Дифф_84М', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('Новое_Аннуитет_ПВ10%', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('Новое_Дифф_6М_ПВ10%', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('Новое_Дифф_12М_ПВ10%', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('Новое_Дифф_24М_ПВ10%', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('Новое_Дифф_36М_ПВ10%', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('Новое_Дифф_48М_ПВ10%', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('Новое_Дифф_60М_ПВ10%', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('Новое_Дифф_72М_ПВ10%', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('Новое_Дифф_84М_ПВ10%', 1);
INSERT INTO bank_products (title, bank_id) VALUES ('Новое_Аннуитет_Субсидия', 1);


INSERT INTO life_insurances (title, percent, bank_id) VALUES ('ORBIS Life 30', 1.10, 1);
INSERT INTO life_insurances (title, percent, bank_id) VALUES ('ORBIS Life 40', 1.40, 1);
INSERT INTO life_insurances (title, percent, bank_id) VALUES ('ORBIS Life 60', 3, 1);

INSERT INTO kaskos (title, percent, bank_id) VALUES ('Orbis Small', 1.50, 1);
INSERT INTO kaskos (title, percent, bank_id) VALUES ('Orbis Medium', 2.15, 1);
INSERT INTO kaskos (title, percent, bank_id) VALUES ('Orbis Large', 2.90, 1);
INSERT INTO kaskos (title, percent, bank_id) VALUES ('VIP', 3, 1);
INSERT INTO kaskos (title, percent, bank_id) VALUES ('VIP Plus', 3.40, 1);

INSERT INTO road_helps (title, price, bank_id) VALUES ('ORBIS ASSISTANCE', 150000, 1);
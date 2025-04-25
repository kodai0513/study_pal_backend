CREATE SEQUENCE article_page_id_seq START 1 INCREMENT 1;
ALTER TABLE articles ALTER COLUMN page_id SET DEFAULT nextval('article_page_id_seq');
ALTER SEQUENCE article_page_id_seq OWNED BY articles.page_id;
 
INSERT INTO
    users(id, name, nick_name, email, password, created_at, updated_at)
VALUES
    ('1b142a47-765f-46ce-be5c-5d37c8ffbca5', 'test', 'テスト', 'test@example.com', '$2a$10$c7jsf8NAbqZNDmYkhso8tOt/Z5bguAjAAljByPRUjMEvovfujtvwO', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
 
INSERT INTO roles(id, name, created_at, updated_at) VALUES ('e77de427-b526-402c-9c9f-4b4b8281ecdd', 'オーナー', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO roles(id, name, created_at, updated_at) VALUES ('773c1d82-ce3d-45cd-afe7-58ced56ed630', 'メンバー', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO roles(id, name, created_at, updated_at) VALUES ('a7dcf2c2-1464-4edf-8734-1dc32a1c5c32', 'ゲスト', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- workbook_permission
INSERT INTO permissions(id, name, created_at, updated_at) VALUES ('53dcebb5-a019-4a41-b152-a5eca98ca472', 'update:workbooks', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO permissions(id, name, created_at, updated_at) VALUES ('e4a0eb0b-2742-429c-af77-0d1b963cdfa1', 'delete:workbooks', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- problem_permission
INSERT INTO permissions(id, name, created_at, updated_at) VALUES ('e08970b5-2999-4bd8-b947-80877c604837', 'create:problems', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- description_problem_permission
INSERT INTO permissions(id, name, created_at, updated_at) VALUES ('6800b0c8-ca99-4ec5-8e9b-8d8b3027c44c', 'update:description-problems', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO permissions(id, name, created_at, updated_at) VALUES ('f02de772-cee1-4fe7-8f7b-6e3d3c77233f', 'delete:description-problems', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- selection_problem_permission
INSERT INTO permissions(id, name, created_at, updated_at) VALUES ('c924478b-8da3-421c-8a77-56bcfcb2f69e', 'update:selection-problems', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO permissions(id, name, created_at, updated_at) VALUES ('dd26376f-12c8-4084-af34-c84f80df7412', 'delete:selection-problems', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- true_or_false_problem_permission
INSERT INTO permissions(id, name, created_at, updated_at) VALUES ('faba8a49-e6e7-43c4-a827-bd8ab70ea917', 'update:true-or-false-problems', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO permissions(id, name, created_at, updated_at) VALUES ('106c3718-ce9e-4374-a5d5-91793ab0e974', 'delete:true-or-false-problems', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- workbook_category_permission
INSERT INTO permissions(id, name, created_at, updated_at) VALUES ('c6e51146-8b47-4bae-9724-71b55e298a6c', 'read:workbook-categories', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO permissions(id, name, created_at, updated_at) VALUES ('c9d0072d-4e41-4451-afdf-dff90d2f9a87', 'update:workbook-categories', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- オーナー
INSERT INTO permission_roles(role_id, permission_id) VALUES ('e77de427-b526-402c-9c9f-4b4b8281ecdd', '53dcebb5-a019-4a41-b152-a5eca98ca472');
INSERT INTO permission_roles(role_id, permission_id) VALUES ('e77de427-b526-402c-9c9f-4b4b8281ecdd', 'e4a0eb0b-2742-429c-af77-0d1b963cdfa1');
INSERT INTO permission_roles(role_id, permission_id) VALUES ('e77de427-b526-402c-9c9f-4b4b8281ecdd', 'e08970b5-2999-4bd8-b947-80877c604837');
INSERT INTO permission_roles(role_id, permission_id) VALUES ('e77de427-b526-402c-9c9f-4b4b8281ecdd', '6800b0c8-ca99-4ec5-8e9b-8d8b3027c44c');
INSERT INTO permission_roles(role_id, permission_id) VALUES ('e77de427-b526-402c-9c9f-4b4b8281ecdd', 'f02de772-cee1-4fe7-8f7b-6e3d3c77233f');
INSERT INTO permission_roles(role_id, permission_id) VALUES ('e77de427-b526-402c-9c9f-4b4b8281ecdd', 'c924478b-8da3-421c-8a77-56bcfcb2f69e');
INSERT INTO permission_roles(role_id, permission_id) VALUES ('e77de427-b526-402c-9c9f-4b4b8281ecdd', 'dd26376f-12c8-4084-af34-c84f80df7412');
INSERT INTO permission_roles(role_id, permission_id) VALUES ('e77de427-b526-402c-9c9f-4b4b8281ecdd', 'faba8a49-e6e7-43c4-a827-bd8ab70ea917');
INSERT INTO permission_roles(role_id, permission_id) VALUES ('e77de427-b526-402c-9c9f-4b4b8281ecdd', '106c3718-ce9e-4374-a5d5-91793ab0e974');
INSERT INTO permission_roles(role_id, permission_id) VALUES ('e77de427-b526-402c-9c9f-4b4b8281ecdd', 'c6e51146-8b47-4bae-9724-71b55e298a6c');
INSERT INTO permission_roles(role_id, permission_id) VALUES ('e77de427-b526-402c-9c9f-4b4b8281ecdd', 'c9d0072d-4e41-4451-afdf-dff90d2f9a87');

-- メンバー
INSERT INTO permission_roles(role_id, permission_id) VALUES ('773c1d82-ce3d-45cd-afe7-58ced56ed630', 'e08970b5-2999-4bd8-b947-80877c604837');
INSERT INTO permission_roles(role_id, permission_id) VALUES ('773c1d82-ce3d-45cd-afe7-58ced56ed630', 'e08970b5-2999-4bd8-b947-80877c604837');
INSERT INTO permission_roles(role_id, permission_id) VALUES ('773c1d82-ce3d-45cd-afe7-58ced56ed630', '6800b0c8-ca99-4ec5-8e9b-8d8b3027c44c');
INSERT INTO permission_roles(role_id, permission_id) VALUES ('773c1d82-ce3d-45cd-afe7-58ced56ed630', 'f02de772-cee1-4fe7-8f7b-6e3d3c77233f');
INSERT INTO permission_roles(role_id, permission_id) VALUES ('773c1d82-ce3d-45cd-afe7-58ced56ed630', 'c924478b-8da3-421c-8a77-56bcfcb2f69e');
INSERT INTO permission_roles(role_id, permission_id) VALUES ('773c1d82-ce3d-45cd-afe7-58ced56ed630', 'dd26376f-12c8-4084-af34-c84f80df7412');
INSERT INTO permission_roles(role_id, permission_id) VALUES ('773c1d82-ce3d-45cd-afe7-58ced56ed630', 'faba8a49-e6e7-43c4-a827-bd8ab70ea917');
INSERT INTO permission_roles(role_id, permission_id) VALUES ('773c1d82-ce3d-45cd-afe7-58ced56ed630', '106c3718-ce9e-4374-a5d5-91793ab0e974');
INSERT INTO permission_roles(role_id, permission_id) VALUES ('773c1d82-ce3d-45cd-afe7-58ced56ed630', 'c6e51146-8b47-4bae-9724-71b55e298a6c');
INSERT INTO permission_roles(role_id, permission_id) VALUES ('773c1d82-ce3d-45cd-afe7-58ced56ed630', 'c9d0072d-4e41-4451-afdf-dff90d2f9a87');

-- ゲスト
INSERT INTO permission_roles(role_id, permission_id) VALUES ('a7dcf2c2-1464-4edf-8734-1dc32a1c5c32', 'c6e51146-8b47-4bae-9724-71b55e298a6c');
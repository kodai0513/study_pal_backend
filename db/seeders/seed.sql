CREATE SEQUENCE article_page_id_seq START 1 INCREMENT 1;
ALTER TABLE articles ALTER COLUMN page_id SET DEFAULT nextval('article_page_id_seq');
ALTER SEQUENCE article_page_id_seq OWNED BY articles.page_id;
 
INSERT INTO
    users(id, name, nick_name, email, password, created_at, updated_at)
VALUES
    ('1b142a47-765f-46ce-be5c-5d37c8ffbca5', 'test', 'テスト', 'test@example.com', '$2a$10$c7jsf8NAbqZNDmYkhso8tOt/Z5bguAjAAljByPRUjMEvovfujtvwO', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
 
INSERT INTO
    roles(id, name, created_at, updated_at)
VALUES
    ('e77de427-b526-402c-9c9f-4b4b8281ecdd', '管理者', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('773c1d82-ce3d-45cd-afe7-58ced56ed630', '編集者', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('a7dcf2c2-1464-4edf-8734-1dc32a1c5c32', '閲覧者', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
 
INSERT INTO
    permissions(id, name, created_at, updated_at)
VALUES
    ('19d2ea70-93a3-4756-bcfd-8e79acadb156', 'create:workbooks', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('53dcebb5-a019-4a41-b152-a5eca98ca472', 'update:workbooks', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('e4a0eb0b-2742-429c-af77-0d1b963cdfa1', 'delete:workbooks', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('e08970b5-2999-4bd8-b947-80877c604837', 'create:problems', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('6800b0c8-ca99-4ec5-8e9b-8d8b3027c44c', 'update:description-problems', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('f02de772-cee1-4fe7-8f7b-6e3d3c77233f', 'delete:description-problems', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('c924478b-8da3-421c-8a77-56bcfcb2f69e', 'update:selection-problems', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('dd26376f-12c8-4084-af34-c84f80df7412', 'delete:selection-problems', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('faba8a49-e6e7-43c4-a827-bd8ab70ea917', 'update:true-or-false-problems', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('106c3718-ce9e-4374-a5d5-91793ab0e974', 'delete:true-or-false-problems', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
 
INSERT INTO
    permission_roles(role_id, permission_id)
VALUES
    ('e77de427-b526-402c-9c9f-4b4b8281ecdd', '19d2ea70-93a3-4756-bcfd-8e79acadb156'),
    ('e77de427-b526-402c-9c9f-4b4b8281ecdd', '53dcebb5-a019-4a41-b152-a5eca98ca472'),
    ('e77de427-b526-402c-9c9f-4b4b8281ecdd', 'e4a0eb0b-2742-429c-af77-0d1b963cdfa1'),
    ('e77de427-b526-402c-9c9f-4b4b8281ecdd', 'e08970b5-2999-4bd8-b947-80877c604837'),
    ('e77de427-b526-402c-9c9f-4b4b8281ecdd', '6800b0c8-ca99-4ec5-8e9b-8d8b3027c44c'),
    ('e77de427-b526-402c-9c9f-4b4b8281ecdd', 'f02de772-cee1-4fe7-8f7b-6e3d3c77233f'),
    ('e77de427-b526-402c-9c9f-4b4b8281ecdd', 'c924478b-8da3-421c-8a77-56bcfcb2f69e'),
    ('e77de427-b526-402c-9c9f-4b4b8281ecdd', 'dd26376f-12c8-4084-af34-c84f80df7412'),
    ('e77de427-b526-402c-9c9f-4b4b8281ecdd', 'faba8a49-e6e7-43c4-a827-bd8ab70ea917'),
    ('e77de427-b526-402c-9c9f-4b4b8281ecdd', '106c3718-ce9e-4374-a5d5-91793ab0e974'),
    ('773c1d82-ce3d-45cd-afe7-58ced56ed630', 'e08970b5-2999-4bd8-b947-80877c604837'),
    ('773c1d82-ce3d-45cd-afe7-58ced56ed630', '6800b0c8-ca99-4ec5-8e9b-8d8b3027c44c'),
    ('773c1d82-ce3d-45cd-afe7-58ced56ed630', 'f02de772-cee1-4fe7-8f7b-6e3d3c77233f'),
    ('773c1d82-ce3d-45cd-afe7-58ced56ed630', 'c924478b-8da3-421c-8a77-56bcfcb2f69e'),
    ('773c1d82-ce3d-45cd-afe7-58ced56ed630', 'dd26376f-12c8-4084-af34-c84f80df7412'),
    ('773c1d82-ce3d-45cd-afe7-58ced56ed630', 'faba8a49-e6e7-43c4-a827-bd8ab70ea917'),
    ('773c1d82-ce3d-45cd-afe7-58ced56ed630', '106c3718-ce9e-4374-a5d5-91793ab0e974');
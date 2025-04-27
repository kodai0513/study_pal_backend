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

-- workbook_invitation_members
INSERT INTO permissions(id, name, created_at, updated_at) VALUES ('b5824a6c-b218-4d84-bd52-2d90cf60e72d', 'read:workbook-invitation-members', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO permissions(id, name, created_at, updated_at) VALUES ('bf3f4e4f-baa5-484c-bd39-ed695653dae8', 'create:workbook-invitation-members', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO permissions(id, name, created_at, updated_at) VALUES ('c367ffa4-3406-4fb4-aa80-9f2eb9fb3e18', 'update:workbook-invitation-members', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO permissions(id, name, created_at, updated_at) VALUES ('3a885be1-4fc9-49db-ac8d-6290c7860369', 'delete:workbook-invitation-members', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- workbook_members
INSERT INTO permissions(id, name, created_at, updated_at) VALUES ('7ebe3d52-7be5-4fb7-8897-478893eafe0e', 'read:workbook-members', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO permissions(id, name, created_at, updated_at) VALUES ('eee6abd0-a010-43a9-acd8-160275a2e7fc', 'update:workbook-members', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO permissions(id, name, created_at, updated_at) VALUES ('dddaeb1c-4a78-41cb-8921-f75f2212e93d', 'delete:workbook-members', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- オーナー
DO $$
DECLARE
    permission_record RECORD;
BEGIN
    FOR permission_record IN
        SELECT id, name FROM permissions
    LOOP
        INSERT INTO permission_roles (role_id, permission_id)
        VALUES ('e77de427-b526-402c-9c9f-4b4b8281ecdd', permission_record.id);
    END LOOP;
END
$$;

-- メンバー
DO $$
DECLARE
    permission_record RECORD;
BEGIN
    FOR permission_record IN
        SELECT id, name FROM permissions
    LOOP
        IF
            permission_record.name NOT LIKE '%:workbooks%' AND
            permission_record.name != 'create:workbook-invitation-members' AND
            permission_record.name != 'update:workbook-invitation-members' AND
            permission_record.name != 'delete:workbook-invitation-members' AND
            permission_record.name != 'update:workbook-members' AND
            permission_record.name != 'delete:workbook-members'
        THEN
            INSERT INTO permission_roles (role_id, permission_id)
            VALUES ('773c1d82-ce3d-45cd-afe7-58ced56ed630', permission_record.id);
        END IF;
    END LOOP;
END
$$;

-- ゲスト
DO $$
DECLARE
    permission_record RECORD;
BEGIN
    FOR permission_record IN
        SELECT id, name FROM permissions
    LOOP
        IF
            permission_record.name LIKE '%read:%'
        THEN
            INSERT INTO permission_roles (role_id, permission_id)
            VALUES ('a7dcf2c2-1464-4edf-8734-1dc32a1c5c32', permission_record.id);
        END IF;
    END LOOP;
END
$$;
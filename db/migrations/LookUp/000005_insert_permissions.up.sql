-- Adminstration permission
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Access Administration', 'read', 'administration');

-- Bots Permissions
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Access Bots', 'read', 'administration');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('List Bots', 'read', 'administration');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Add Bot', 'create', 'administration');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Update Bot', 'update', 'administration');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Delete Bots', 'delete', 'administration');

-- Roles Permissions
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Access Roles', 'read', 'administration');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('List Roles', 'read', 'administration');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Add Role', 'create', 'administration');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Update Role', 'update', 'administration');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Delete Role', 'delete', 'administration');

-- Users Permissions
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Access Users', 'read', 'administration');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('List Users', 'read', 'administration');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Add Users', 'create', 'administration');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Update Users', 'update', 'administration');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Delete Users', 'delete', 'administration');



-- Cards 
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Add Card', 'create', 'builder');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Delete Card', 'delete', 'builder');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Update Card', 'update', 'builder');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Sort Card', 'update', 'builder');

-- Blocks
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Access Block', 'read', 'builder');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Add Block', 'create', 'builder');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Delete Block', 'delete', 'builder');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Copy Block', 'create', 'builder');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Update Block', 'update', 'builder');

-- Groups 
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Access Group', 'read', 'builder');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('List Group', 'read', 'builder');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Add Group', 'create', 'builder');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Delete Group', 'delete', 'builder');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Update Group', 'update', 'builder');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Sort Groups', 'update', 'builder');


-- NLP Permissions
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Access NLP Rules', 'read', 'nlp');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('List NLP Rules', 'read', 'nlp');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Add NLP Rule', 'create', 'nlp');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Delete NLP Rule', 'delete', 'nlp');

-- Word Spotting Permissions
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Access Word Spotting', 'read', 'wordspotting');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('List Word Spotting', 'read', 'wordspotting');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Add Word Spotting', 'create', 'wordspotting');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Update Word Spotting', 'update', 'wordspotting');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Delete Word Spotting', 'delete', 'wordspotting');

-- Settings Permissions
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Access Settings', 'read', 'settings');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Access General Settings', 'read', 'settings');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Update General', 'update', 'settings');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Access Channels Settings', 'read', 'settings');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Add Channel', 'create', 'settings');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Update Channel', 'update', 'settings');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Delete Channel', 'delete', 'settings');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Access NLP Settings', 'read', 'settings');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Add NLP', 'create', 'settings');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Update NLP', 'update', 'settings');
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Delete NLP', 'delete', 'settings');


-- Insights Permissions
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Access Insights', 'read', 'insights');


-- Builder Permissions
INSERT INTO permissions
    (name, type, category)
VALUES
    ('Access Builder', 'read', 'builder');
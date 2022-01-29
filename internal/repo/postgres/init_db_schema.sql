/*
 
инициализация БД

*/

-- =============================
-- under user postgres:

CREATE USER appuser WITH PASSWORD 'appp@$$w0rd';

CREATE DATABASE app_db
    TEMPLATE = 'template0'
    ENCODING = 'utf-8'
    LC_COLLATE = 'C.UTF-8'
    LC_CTYPE = 'C.UTF-8';

REVOKE ALL ON DATABASE app_db FROM PUBLIC;
ALTER DATABASE app_db OWNER TO appuser;

-- =============================
-- connect to database:

\c app_db

-- =============================
-- under user appuser:

SET ROLE appuser;

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    uuid CHAR(36) UNIQUE NOT NULL,
    user_name VARCHAR(100) NOT NULL,
    email VARCHAR(254) UNIQUE NOT NULL,
    
    PRIMARY KEY(uuid)
);

DROP TABLE IF EXISTS groups;
CREATE TABLE groups (
    uuid CHAR(36) NOT NULL,
    group_name VARCHAR(50) NOT NULL,
    group_type VARCHAR(254) UNIQUE NOT NULL,
    descr VARCHAR(254) NOT NULL,
    
    PRIMARY KEY(uuid)
);


DROP TABLE IF EXISTS membership;
CREATE TABLE membership (
    group_uuid CHAR(36) REFERENCES groups (uuid),
    user_uuid  CHAR(36) REFERENCES users (uuid),

    constraint membership_fk_group_uuid FOREIGN KEY (group_uuid) references groups (uuid) on delete restrict,
    constraint membership_fk_user_uuid FOREIGN KEY (user_uuid) references users (uuid) on delete restrict
);

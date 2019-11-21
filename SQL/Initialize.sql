DROP TABLE PICTURES;
DROP TABLE APPUSER;

CREATE TABLE APPUSER (
    ID         SERIAL        PRIMARY KEY,
    CREATED_AT TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,
    USERNAME   VARCHAR(30)   UNIQUE,
    EMAIL      VARCHAR(255)  UNIQUE,
    PASSWORD   VARCHAR(1024)
);

INSERT INTO APPUSER (EMAIL, USERNAME, PASSWORD) VALUES ('admin@scinna.dev', 'admin', 'nopass');

CREATE TABLE PICTURES (
    ID         SERIAL         PRIMARY KEY,
    CREATED_AT TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    TITLE      VARCHAR(30),
    URL_ID     VARCHAR(30),
    DESCRIPT   VARCHAR(1024),
    VISIBILITY INTEGER,
    CREATOR    INTEGER        REFERENCES APPUSER(ID)
);

INSERT INTO PICTURES (TITLE, URL_ID, DESCRIPT, VISIBILITY, CREATOR)
VALUES (
    'Ma photo',
    'dJe8g2-FjC',
    'Ceci est une photo que j.ai uploadé',
    0,
    (SELECT ID FROM APPUSER WHERE USERNAME = 'admin')
);

INSERT INTO PICTURES (TITLE, URL_ID, DESCRIPT, VISIBILITY, CREATOR)
VALUES (
    'Ma photo 2',
    'sFueG_Hl23',
    'Autre de mes photos, mais non-listée',
    1,
    (SELECT ID FROM APPUSER WHERE USERNAME = 'admin')
);

INSERT INTO PICTURES (TITLE, URL_ID, DESCRIPT, VISIBILITY, CREATOR)
VALUES (
    'Ma photo 3',
    'qKJbcD_73f',
    'Une petite dernière en privée',
    2,
    (SELECT ID FROM APPUSER WHERE USERNAME = 'admin')
);
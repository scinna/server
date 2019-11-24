DROP TABLE PICTURES;
DROP TABLE APPUSER;

CREATE TABLE APPUSER (
    ID         SERIAL        PRIMARY KEY,
    CREATED_AT TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,
    USERNAME   VARCHAR(30)   UNIQUE,
    EMAIL      VARCHAR(255)  UNIQUE,
    PASSWORD   VARCHAR(1024)
);

-- Default password is 'password'
INSERT INTO APPUSER (EMAIL, USERNAME, PASSWORD) VALUES ('admin@scinna.dev', 'admin', '$argon2id$v=19$m=65536,t=3,p=2$D1hPKoAbrexDtJd6uEf3Cg$d1puA1YPJgUkvvTaotKpRWOT2xIMMIUknyl6IeWJsfQ');

CREATE TABLE PICTURES (
    ID         SERIAL         PRIMARY KEY,
    CREATED_AT TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    TITLE      VARCHAR(30),
    URL_ID     VARCHAR(30),
    DESCRIPT   VARCHAR(1024),
    VISIBILITY INTEGER,
    CREATOR    INTEGER        REFERENCES APPUSER(ID),
    EXT        VARCHAR(8)
);

INSERT INTO PICTURES (TITLE, URL_ID, DESCRIPT, VISIBILITY, CREATOR, EXT)
VALUES (
    'Ma photo',
    'dJe8g2-FjC',
    'Ceci est une photo que j.ai uploadé',
    0,
    (SELECT ID FROM APPUSER WHERE USERNAME = 'admin'),
    'png'
);

INSERT INTO PICTURES (TITLE, URL_ID, DESCRIPT, VISIBILITY, CREATOR, EXT)
VALUES (
    'Ma photo 2',
    'sFueG_Hl23',
    'Autre de mes photos, mais non-listée',
    1,
    (SELECT ID FROM APPUSER WHERE USERNAME = 'admin'),
    'jpg'
);

INSERT INTO PICTURES (TITLE, URL_ID, DESCRIPT, VISIBILITY, CREATOR, EXT)
VALUES (
    'Ma photo 3',
    'qKJbcD_73f',
    'Une petite dernière en privée',
    2,
    (SELECT ID FROM APPUSER WHERE USERNAME = 'admin'),
    'jpg'
);
CREATE TABLE users (
    id SERIAL NOT NULL ,
    firstname VARCHAR(100) NOT NULL ,
    lastname VARCHAR(100) ,
    picture VARCHAR(256) ,
    nickname VARCHAR(100),
    is_staff BOOLEAN DEFAULT FALSE NOT NULL ,
    email VARCHAR(256) NOT NULL ,
    is_email_confirmed BOOLEAN DEFAULT FALSE NOT NULL ,
    phone VARCHAR(50) ,
    password VARCHAR(256) NOT NULL,
    created_at TIMESTAMP DEFAULT now() NOT NULL ,
    updated_at TIMESTAMP,
    CONSTRAINT users_id_pk PRIMARY KEY (id) ,
    CONSTRAINT users_nickname_uq UNIQUE (nickname),
    CONSTRAINT users_email_uq UNIQUE (email)
);

COMMENT ON TABLE users IS 'Save users data';
CREATE TABLE roles (
    id SERIAL NOT NULL ,
    name VARCHAR(100) NOT NULL ,
    is_active BOOLEAN DEFAULT TRUE NOT NULL ,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    CONSTRAINT roles_id_pk PRIMARY KEY (id),
    CONSTRAINT roles_name_uq UNIQUE (name)
);
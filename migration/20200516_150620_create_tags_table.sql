CREATE TABLE tags (
    id SERIAL NOT NULL ,
    name VARCHAR(50) NOT NULL ,
    is_active BOOLEAN DEFAULT FALSE NOT NULL ,
    created_at TIMESTAMP DEFAULT now() NOT NULL ,
    updated_at TIMESTAMP,
    CONSTRAINT tags_id_pk PRIMARY KEY (id),
    CONSTRAINT tags_name_uq UNIQUE (name)
);

COMMENT ON TABLE tags IS 'Tags for created tasks';
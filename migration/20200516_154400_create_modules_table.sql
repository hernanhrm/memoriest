CREATE TABLE modules(
   id SERIAL NOT NULL ,
   name VARCHAR(100) NOT NULL ,
   description VARCHAR(256) ,
   is_active BOOLEAN DEFAULT TRUE NOT NULL ,
   created_at TIMESTAMP NOT NULL DEFAULT now(),
   updated_at TIMESTAMP,
   CONSTRAINT modules_id_pk PRIMARY KEY (id),
   CONSTRAINT modules_name_uq UNIQUE (name)
);

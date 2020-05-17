CREATE TABLE module_roles (
    id SERIAL NOT NULL ,
    module_id INTEGER NOT NULL ,
    role_id INTEGER NOT NULL ,
    read BOOLEAN NOT NULL DEFAULT TRUE,
    post BOOLEAN NOT NULL DEFAULT TRUE,
    modify BOOLEAN NOT NULL DEFAULT TRUE,
    erease BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP ,
    CONSTRAINT module_roles_id_pk PRIMARY KEY (id),
    CONSTRAINT module_roles_module_id_role_id_uq UNIQUE (module_id, role_id),
    CONSTRAINT module_roles_module_id_fk FOREIGN KEY (module_id) REFERENCES modules(id) ON UPDATE RESTRICT ON DELETE CASCADE ,
    CONSTRAINT module_roles_role_id_fk FOREIGN KEY (role_id) REFERENCES roles(id) ON UPDATE RESTRICT ON DELETE CASCADE
);
CREATE TABLE user_roles (
    id SERIAL NOT NULL ,
    user_id INTEGER NOT NULL ,
    role_id INTEGER NOT NULL ,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    CONSTRAINT user_roles_id_pk PRIMARY KEY (id),
    CONSTRAINT user_roles_user_id_role_id_uq UNIQUE (user_id, role_id),
    CONSTRAINT user_roles_user_id_fk FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE RESTRICT ON DELETE CASCADE ,
    CONSTRAINT user_roles_role_id_fk FOREIGN KEY (role_id) REFERENCES roles(id) ON UPDATE RESTRICT ON DELETE CASCADE
);
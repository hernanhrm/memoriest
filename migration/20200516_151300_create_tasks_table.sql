CREATE TABLE tasks (
    id SERIAL NOT NULL ,
    name VARCHAR(50) NOT NULL ,
    slug VARCHAR(256) NOT NULL ,
    description TEXT,
    due_date DATE,
    due_time TIME,
    is_completed BOOLEAN DEFAULT FALSE NOT NULL ,
    is_parent_task BOOLEAN DEFAULT TRUE NOT NULL ,
    created_at TIMESTAMP DEFAULT now() NOT NULL ,
    updated_at TIMESTAMP,
    CONSTRAINT tasks_id_pk PRIMARY KEY (id),
    CONSTRAINT tasks_slug_uq UNIQUE (slug)
);

COMMENT ON TABLE tasks IS 'Save tasks for users';

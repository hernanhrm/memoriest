CREATE TABLE user_tasks (
    id SERIAL NOT NULL ,
    user_id INTEGER NOT NULL ,
    task_id INTEGER NOT NULL ,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    CONSTRAINT user_tasks_id_pk PRIMARY KEY (id),
    CONSTRAINT user_tasks_user_id_task_id_uq UNIQUE (user_id, task_id),
    CONSTRAINT user_tasks_user_id_fk FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE RESTRICT ON DELETE CASCADE ,
    CONSTRAINT user_tasks_task_id_fk FOREIGN KEY (task_id) REFERENCES tasks(id) ON UPDATE RESTRICT ON DELETE CASCADE
);

COMMENT ON TABLE user_tasks IS 'Store all user tasks';
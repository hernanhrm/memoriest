CREATE TABLE subtasks (
    id SERIAL NOT NULL ,
    task_id INTEGER NOT NULL ,
    subtask_id INTEGER NOT NULL ,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    udpated_at TIMESTAMP,
    CONSTRAINT subtasks_id_pk PRIMARY KEY (id),
    CONSTRAINT subtasks_task_idsubtask_id_uq UNIQUE (task_id, subtask_id),
    CONSTRAINT subtasks_task_id_fk FOREIGN KEY (task_id) REFERENCES tasks(id) ON UPDATE RESTRICT ON DELETE CASCADE ,
    CONSTRAINT subtasks_subtask_id FOREIGN KEY (subtask_id) REFERENCES tasks(id) ON UPDATE RESTRICT ON DELETE CASCADE
);

COMMENT ON TABLE subtasks IS 'Save subtasks of a task';
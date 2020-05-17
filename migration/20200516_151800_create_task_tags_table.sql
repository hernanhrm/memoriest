CREATE TABLE task_tags (
    id SERIAL NOT NULL ,
    task_id INTEGER NOT NULL ,
    tag_id INTEGER NOT NULL ,
    created_at TIMESTAMP DEFAULT now() NOT NULL ,
    updated_at TIMESTAMP,
    CONSTRAINT task_tags_id_pk PRIMARY KEY (id),
    CONSTRAINT task_tags_task_id_tag_id_uq UNIQUE (task_id, tag_id),
    CONSTRAINT task_tags_task_id_fk FOREIGN KEY (task_id) REFERENCES tasks(id) ON UPDATE RESTRICT ON DELETE CASCADE ,
    CONSTRAINT task_tags_tag_id_fk FOREIGN KEY (tag_id) REFERENCES tags(id) ON UPDATE RESTRICT ON DELETE CASCADE
);

COMMENT ON TABLE task_tags IS 'Allow to store tags for a task';
create table t_task
(
	id int auto_increment comment '任务ID'
		primary key,
	parent_id int null comment '父任务ID',
	name varchar(50) null comment '任务名称',
	content varchar(1024) null comment '任务描述'
)
comment '任务表' engine=InnoDB
;

create index t_task_parent_id_index
	on t_task (parent_id)
;
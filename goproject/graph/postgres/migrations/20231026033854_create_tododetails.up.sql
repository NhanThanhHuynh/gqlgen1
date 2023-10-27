create table todo_details (
    id varchar(255) primary key not null ,
    text varchar(255) null,
    user_id varchar(255) references users(id) on delete cascade not null
)
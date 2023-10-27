create table todos (
    id  varchar(255) primary key not null,
    text  varchar(255) null,
    done  boolean null,
    user_id varchar(255) references users(id) ON DELETE CASCADE NOT NULL
)
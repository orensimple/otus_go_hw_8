create table events (
    id int primary key,
    owner text,
    title text,
    text text,
    start_time timestamp not null,
    end_time timestamp
)

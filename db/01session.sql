create table usuario(
    username varchar(80) not null,
    password varchar(80) not null,
    state boolean default true,
    created_at timestamp default CURRENT_TIMESTAMP
    owner_entity varchar(80) default '',
);
insert into usuario(username,password,owner_entity) values('kevin002','maira002','admin');
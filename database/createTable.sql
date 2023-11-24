create table if not exists snippets (
    id integer primary key AUTO_INCREMENT,
    title varchar(100) not null,
    content text not null,
    created datetime not null,
    expires datetime not null
);

create index idx_snippets_created on snippets(created);
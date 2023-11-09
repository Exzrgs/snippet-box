create user 'web'@'localhost';
grant select, insert, update, delete on snippetbox. * to 'web'@'localhost';
alter user 'web'@'localhost' identified by 'pass';
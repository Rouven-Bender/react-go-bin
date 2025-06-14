create table content (
	id char(37),
	ctype int,
	data varchar(2048)
);
create table creds (
	userid integer primary key,
	key_hash char(44)
);
create unique index idx_key_hash on creds(key_hash)

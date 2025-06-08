create table content (
	id char(37),
	ctype int,
	data varchar(2048)
);
create table creds (
	key_hash char(44) primary key
)

CREATE EXTENSION pgcrypto;

create or replace function fn_login(
    _username varchar(80),
    _password varchar(80)
) returns setof usuario as
$$
begin
    return query select  
		u.username,
		cast('***************' as varchar(80)) as password,
		u.state, u.create_at , u.owner_entity
		from usuario u where u.username = _username and u.password=crypt(_password,u.password);
end;
$$
language plpgsql;

-- creacion de usuario
create or replace function fn_create_user(
    _username varchar(80),
    _password varchar(80)
) returns setof usuario as
$$
declare
    
begin
    if (select count(u.*) from usuario u where u.username = _username)>0 then
        raise exception '{message","usuario {%}, ya existe"}',_username;
    end if;
    insert into usuario(username,password) values(_username,crypt(_password, gen_salt('bf', 8)));
    return query select
		u.username,
		cast('***************' as varchar(80)) as password,
		u.state, u.create_at, u.owner_entity
	from usuario u where username=_username;
end;
$$
language plpgsql;
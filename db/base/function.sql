-- text format, response
create or replace function f_response(code char(3),message varchar(250))
returns text as
$$
begin
    return FORMAT('{"code":"%s","message":"%s"}',code,message);
end;
$$
language plpgsql;

select * from usuario
create or replace function fn_actividades(
    _fecha_inicio date,
    _fecha_fin date
)
    returns table
            (
                id           int,
                titulo       varchar(100),
                hora_inicio  int,
                hora_fin     int,
                owner_dni    char(8),
                fecha        date,
                owner_nombre text
            )
as
$$
begin
    if _fecha_inicio > _fecha_fin then
        return;
    end if;
    return query select a.id,
                        a.titulo,
                        a.hora_inicio,
                        a.hora_fin,
                        a.owner_dni,
                        p.fecha,
                        concat(e.nombre, ' ', e.apellido_paterno, ' ', e.apellido_materno) as owner_nombre
                 from actividad a
                          inner join
                      programacion p on a.programacion_id = p.id
                          inner join entity e on a.owner_dni = e.dni
                 where p.fecha >= _fecha_inicio
                   and p.fecha < _fecha_fin order by p.fecha;
end;
$$ language plpgsql;
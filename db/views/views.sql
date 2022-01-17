create view vw_instructor as
select e.*, i.state, i.detalle, i.fecha_instructor
from entity e
         inner join instructor i on e.dni = i.dni;
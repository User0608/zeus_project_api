--- tablas constantes
create table nivel_estudio
(
    id    int         not null,
    nivel varchar(80) not null
);
alter table nivel_estudio
    add constraint pk_nivel_estudio primary key (id);

create table estado_civil
(
    id     int         not null,
    estado varchar(80) not null
);
alter table estado_civil
    add constraint pk_estado_civil primary key (id);

insert into estado_civil(id, estado)
values (1, 'soltero'),
       (2, 'casado'),
       (3, 'divorciado'),
       (4, 'separación en proceso judicial'),
       (5, 'viudo'),
       (6, 'concubinato');
       
insert into nivel_estudio(id, nivel)
values (1, 'sin nivel'),
       (2,'primaria incompleta'),
       (3,'primaria'),
       (4,'secundaria incompleta'),
       (5,'secundaria'),
       (6,'superior');
------------------------------------------------------------

create table entity
(
    dni              char(8)     not null,
    nombre           varchar(80) not null,
    apellido_paterno varchar(80) not null,
    apellido_materno varchar(80) not null,
    direccion        varchar(80) not null,
    telefono         varchar(15) not null,
    email            varchar(80)          default '-',
    fecha_nacimiento date        not null,
    nivel_estudio_id int         not null default 1,
    estado_civil_id  int         not null default 1,
    fecha_registro   timestamp   not null default CURRENT_TIMESTAMP
);
alter table entity
    add constraint pk_entity primary key (dni),
    add constraint fk_entity__nivel_estudio foreign key (nivel_estudio_id) references nivel_estudio (id),
    add constraint fk_entity__estado_civil foreign key (estado_civil_id) references estado_civil (id);
    
create table primer_jefe
(
    dni            char(8) not null,
    state          boolean not null default true,
    fecha_registro timestamp        default current_timestamp
);
alter table primer_jefe
    add constraint pk_primer_jefe primary key (dni),
    add constraint fk_primer_jefe__entity foreign key (dni) references entity(dni);

create table segundo_jefe(
    dni char(8) not null,
    state boolean default false,
    fecha_registro timestamp default  current_timestamp,
    constraint pk_segundo_jefe primary key(dni),
    constraint fk_segundo_jefe__entity foreign key (dni)
                         references entity(dni)
);
create table jefe_instruccion
(
    dni            char(8) not null,
    state          boolean not null default true,
    detalle        varchar          default '',
    fecha_registro timestamp        default current_timestamp,
    constraint pk_jefe_instruccion primary key (dni),
    constraint fk_jefe_instruccion__entity foreign key (dni)
        references entity (dni)
);
create table instructor
(
    dni            char(8) not null,
    state          boolean not null default true,
    detalle        varchar          default '',
    fecha_registro timestamp        default current_timestamp,
    constraint pk_instructor primary key (dni),
    constraint fk_pk_instructor__entity foreign key (dni)
        references entity (dni)
);

create table modulo_esbas
(
    nombre         varchar(80) not null,
    descripcion    varchar(500),
    numero_horas   int default 0,
    numero_alumnos int default 0,
    comentarios    varchar(500),
    create_by      char(8)     not null,
    constraint pk_modulo_esbas primary key (nombre),
    constraint fk_modulo_esbas__jefe_instruccion
        foreign key (create_by) references jefe_instruccion (dni)
);
create table leccion
(
    nombre       varchar(80)  not null,
    numero_horas int default 0,
    descripcion  varchar(250) not null,
    modulo_name  varchar(80)  not null,
    constraint pk_lession primary key (nombre),
    constraint fk_lession__modulo_esbas
        foreign key (modulo_name) references modulo_esbas (nombre)
);
create table ciclo(
    nombre varchar(80) not null,
    fecha timestamp default current_timestamp,
    detalles varchar(250),
    constraint pk_ciclo primary key (nombre)
);
create table curso_ciclo
(
    id      serial,
    nombre_ciclo  varchar(80) not null,
    nombre_modulo varchar(80) not null,
    constraint pk_curso_ciclo primary key (id),
    constraint fk_curso_ciclo__ciclo
        foreign key (nombre_ciclo) references ciclo (nombre),
    constraint fk_curso_ciclo__modulo
        foreign key (nombre_modulo) references modulo_esbas (nombre)
);
create table detalle_instructor
(
    curso_id       int     not null,
    instructor_dni char(8) not null,
    constraint pk_detalle_instructor primary key (curso_id, instructor_dni),
    constraint fk_detalle_instructor__curso
        foreign key (curso_id) references curso_ciclo (id),
    constraint fk_detalle_instructor__instructor
        foreign key (instructor_dni) references instructor (dni)
);

create table memorando
(
    codigo      varchar(80)  not null,
    parte_del   varchar(200) not null,
    dirigido_al varchar(200) not null,
    asunto      varchar(500) not null,
    fecha       timestamp default current_timestamp,
    contenido   text         not null,
    create_by   char(8)      not null,
    constraint pk_memorando primary key (codigo),
    constraint fk_memorando__entity foreign key (create_by)
        references entity (dni)
);
create table informe
(
    codigo      varchar(80)  not null,
    parte_del   varchar(200) not null,
    dirigido_al varchar(200) not null,
    asunto      varchar(500) not null,
    fecha       timestamp default current_timestamp,
    contenido   text         not null,
    create_by   char(8)      not null,
    constraint pk_informe primary key (codigo),
    constraint fk_informe__entity foreign key (create_by)
        references entity (dni)
);
create table oficio
(
    codigo      varchar(80)  not null,
    dirigido_al varchar(200) not null,
    asunto      varchar(500) not null,
    fecha       timestamp default current_timestamp,
    contenido   text         not null,
    create_by   char(8)      not null,
    constraint pk_oficio primary key (codigo),
    constraint fk_oficio__entity foreign key (create_by)
        references entity (dni)
);
create table convocatoria
(
    nombre      varchar(80)  not null,
    fecha       timestamp default current_timestamp,
    descripcion varchar(280) not null,
    creado_pr   varchar(80)  not null,
    constraint pk_convocatoria primary key (nombre)
);
create table aspirante
(
    dni                 char(8)     not null,
    deleted             bool      default false,
    fecha_registro      timestamp default current_timestamp,
    horas               int       default 0,
    password            varchar(80) not null,
    is_pasword_default  bool      default true,
    password_update     timestamp default current_timestamp,
    convocatoria_nombre varchar(80) not null,
    constraint pk_aspirante primary key (dni),
    constraint fk_aspirante__entity foreign key (dni)
        references entity (dni),
    constraint fk_aspirante__convocatoria foreign key (convocatoria_nombre)
        references convocatoria (nombre)
);

create table matricula
(
    id                   serial  not null,
    aspirante_dni        char(8) not null,
    calificacion         decimal(4, 2) default 0,
    calificaion_pracitas decimal(4, 2) default 0,
    nota_final           decimal(4, 2) default 0,
    numero_faltas        int           default 0,
    fecha                timestamp     default current_timestamp,
    curso_id             int     not null,
    constraint pk_matricula primary key (id),
    constraint fk_matricula__ciclo_curso foreign key (curso_id)
        references curso_ciclo (id),
    constraint fk_matricula__aspirante foreign key (aspirante_dni)
        references aspirante (dni)
);
create table asistencia
(
    id           serial      not null,
    leccion      varchar(80) not null,
    estado       varchar default 'A',
    matricula_id int         not null,
    constraint pk_asistencia primary key (id),
    constraint ckh_estado check ( estado in ('A', 'F', 'a', 'f') ),
    constraint fk_asistencia__matricula foreign key (matricula_id)
        references matricula (id)
);
create table cronograma
(
    id          serial,
    nombre      varchar(100) not null,
    descripcion varchar(280) default '',
    created_by  char(8)      not null,
    constraint pk_cronograma primary key (id),
    constraint fk_cronograma__primer_jefe foreign key (created_by)
        references primer_jefe (dni)
);
create table programacion
(
    id            serial,
    nombre        varchar(100) not null,
    detalle       varchar(280) default '',
    fecha         date         default current_date,
    cronograma_id int          not null,
    constraint pk_programacion primary key (id),
    constraint fk_programacion__cronograma foreign key (cronograma_id) references cronograma (id)
);
create table actividad
(
    id              serial,
    titulo          varchar(100) not null,
    detalle         varchar(280) default '',
    hora_inicio     int          not null,
    hora_fin        int          not null,
    owner_dni      char(8)      default '00000000',
    programacion_id int          not null,
    constraint chk_hora_inicio check (hora_inicio >= 0 and hora_inicio < 24 ),
    constraint chk_hora_fin check (hora_fin >= 0 and hora_fin < 25 ),
    constraint pk_actividad primary key (id),
    constraint fk_actividad__entity foreign key (owener_dni)
        references entity (dni),
    constraint fk_actividad__programacion foreign key (programacion_id)
        references programacion (id)
);
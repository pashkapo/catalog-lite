create table if not exists buildings
(
    id       serial primary key,
    country  varchar(128) not null,
    city     varchar(128) not null,
    street   varchar(128) not null,
    house    varchar(128) not null,
    location point        not null
);

create index idx_location on buildings using gist ((location::geometry::geography));

create table if not exists firms
(
    id          serial primary key,
    name        varchar(128) not null,
    building_id integer references buildings (id)
);

create index if not exists idx_building_id on firms (building_id);


create table if not exists rubrics
(
    id        serial primary key,
    parent_id integer references rubrics (id),
    name      varchar(128) not null,
    level     integer default 0
);

create index if not exists idx_parent_id on rubrics (parent_id);


create table if not exists firms_rubrics
(
    firm_id   integer references firms (id),
    rubric_id integer references rubrics (id)
);

create index if not exists idx_firm_id on firms_rubrics (firm_id);

create index if not exists idx_rubric_id on firms_rubrics (rubric_id);

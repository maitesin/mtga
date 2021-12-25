CREATE TABLE IF NOT EXISTS cards (
                                     id blob PRIMARY KEY,
                                     name       varchar,
                                     language   varchar,
                                     url        varchar,
                                     set_name    varchar,
                                     rarity     varchar,
                                     image      varchar,
                                     mana_cost   varchar,
                                     reprint    int2,
                                     price      varchar,
                                     released_at datetime,
                                     opts int
);

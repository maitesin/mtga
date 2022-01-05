CREATE TABLE IF NOT EXISTS cards (
                                     id blob,
                                     name       varchar,
                                     language   varchar,
                                     url        varchar,
                                     set_name    varchar,
                                     rarity     varchar,
                                     mana_cost   varchar,
                                     reprint    int2,
                                     price      varchar,
                                     released_at datetime,
                                     opts int,
                                     quantity int,
                                     condition varchar,
                                     PRIMARY KEY (id, language, opts, condition)
);

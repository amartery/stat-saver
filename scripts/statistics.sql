CREATE TABLE IF NOT EXISTS Statistic
(
    stat_id         bigserial not null primary key,
    event_date      date not null,
    views           bigint not null,
    clicks          bigint not null,
    cost            real not null,
    cpc             real not null,
    cpm             real not null
);


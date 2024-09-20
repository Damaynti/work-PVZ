-- +goose Up
-- +goose StatementBegin
create table pvz(
    id BIGSERIAL PRIMARY KEY NOT NULL ,
    title TEXT NOT NULL DEFAULT '' ,
    address TEXT NOT NULL DEFAULT '' ,
    contactInformation TEXT NOT NULL DEFAULT '' ,
    isDel BOOLEAN DEFAULT FALSE ,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
create table orders(
    id BIGSERIAL PRIMARY KEY NOT NULL ,
    pvz_id BIGINT REFERENCES pvz(id),
    fullName TEXT NOT NULL DEFAULT '' ,
    status TEXT NOT NULL DEFAULT '' ,
    orderCode TEXT NOT NULL DEFAULT '' ,
    isDel BOOLEAN DEFAULT FALSE ,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table pvz;
-- +goose StatementEnd

-- +goose StatementBegin
drop table orders;
-- +goose StatementEnd
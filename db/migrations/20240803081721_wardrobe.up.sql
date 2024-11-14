CREATE TABLE "wardrobe" (
    id uuid NOT NULL PRIMARY KEY,
    name varchar(255),
    color varchar(100),
    size varchar(10),
    price float DEFAULT 0,
    stock int DEFAULT 0,
    created_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now()
);
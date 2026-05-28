CREATE DATABASE fastcomp;

\c fastcomp

CREATE TABLE roles (
    id   SERIAL      PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

INSERT INTO roles (name) VALUES
    ('admin'),
    ('business');

CREATE TABLE users (
    id                SERIAL       PRIMARY KEY,
    first_name        VARCHAR(100) NOT NULL,
    last_name         VARCHAR(100) NOT NULL,
    business_name     VARCHAR(200) NOT NULL,
    email             VARCHAR(200) NOT NULL UNIQUE,
    password          TEXT         NOT NULL,
    phone             CHAR(10)     NOT NULL,
    website           TEXT,
    profile_photo     TEXT,
    registration_date TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    role_id           INTEGER      NOT NULL DEFAULT 2 REFERENCES roles(id)
);

CREATE INDEX idx_users_email ON users(email);

CREATE TYPE product_category AS ENUM (
    'audio',
    'computacion',
    'movil',
    'smarthome',
    'fotografia',
    'accesorios'
);

CREATE TYPE product_condition AS ENUM (
    'nuevo',
    'reacondicionado',
    'segunda_mano'
);

CREATE TABLE products (
    id          VARCHAR(20)       PRIMARY KEY,
    brand       VARCHAR(100)      NOT NULL,
    name        VARCHAR(200)      NOT NULL,
    category    product_category  NOT NULL,
    price       DECIMAL(10, 2)    NOT NULL CHECK (price > 0),
    stock       INTEGER           NOT NULL DEFAULT 0 CHECK (stock >= 0),
    condition   product_condition NOT NULL,
    featured    BOOLEAN           NOT NULL DEFAULT false,
    description TEXT,
    specs       TEXT[],
    image_url   TEXT,
    added_at    DATE              NOT NULL DEFAULT CURRENT_DATE
);

CREATE INDEX idx_products_category ON products(category);
CREATE INDEX idx_products_featured ON products(featured) WHERE featured = true;

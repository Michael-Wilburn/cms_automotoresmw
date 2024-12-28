CREATE TABLE IF NOT EXISTS cars (
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    online boolean NOT NULL DEFAULT false,
    car_type text NOT NULL,
    brand text NOT NULL,
    model text NOT NULL,
    year integer NOT NULL,
    kilometers bigint NOT NULL,
    car_domain text NOT NULL,
    price numeric(15,0) NOT NULL,
    info_price numeric(15,0) NOT NULL,
    currency text NOT NULL DEFAULT '$'::text,
    chasis_code text NOT NULL,
    motor_code text NOT NULL,
    images_urls text[] NOT NULL DEFAULT '{}',
    CONSTRAINT car_pkey PRIMARY KEY (id)
);

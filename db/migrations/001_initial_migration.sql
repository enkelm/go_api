-- Write your migrate up statements here
CREATE TABLE public.user(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    first_name VARCHAR(80) NOT NULL,
    last_name VARCHAR(80) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(60) NOT NULL
);

INSERT INTO public.user (first_name, last_name, email, password_hash) VALUES ('Enkel', 'Murati', 'enkel.murati33@gmail.com', '$2a$12$cupWkmlzPDRGz1XQRl6dI.uBs.6Joj/pz6xy90FJIBvIYHdhvlBfe')

---- create above / drop below ----
drop table user;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.

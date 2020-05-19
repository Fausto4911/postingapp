CREATE TABLE IF NOT EXISTS post (
	id SERIAL PRIMARY KEY NOT NULL,
	user_id INTEGER NOT NULL,
	title VARCHAR(100) NOT NULL,
	text  TEXT,
	category_id INTEGER NOT NULL,
	create_at TIMESTAMP DEFAULT NOW(),
	image VARCHAR(2048)
)

CREATE TABLE IF NOT EXISTS comment (
   id SERIAL PRIMARY KEY NOT NULL,
    user_id INTEGER NOT NULL,
	post_id INTEGER NOT NULL,
    text TEXT NOT NULL,
	create_at TIMESTAMP DEFAULT NOW(),
	parent INTEGER,
	child INTEGER
)

CREATE TABLE IF NOT EXISTS app_user (
    id SERIAL PRIMARY KEY NOT NULL,
	name VARCHAR(50) NOT NULL,
	password VARCHAR(100) NOT NULL,
	avatar VARCHAR(2048),
	create_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS category (
    id SERIAL PRIMARY KEY NOT NULL,
	name VARCHAR(50) NOT NULL,
	avatar VARCHAR(2048),
	create_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS likes (
    id SERIAL PRIMARY KEY NOT NULL,
	user_id INTEGER NOT NULL,
	type_id INTEGER,
	quantity INTEGER,
	entity_type INTEGER
);

ALTER TABLE post ADD CONSTRAINT post_app_user_id_fkey FOREIGN KEY (user_id)
REFERENCES app_user (id);
ALTER TABLE post ADD CONSTRAINT post_category_id_fkey FOREIGN KEY (category_id)
REFERENCES category (id);
ALTER TABLE comment ADD CONSTRAINT comment_app_user_id_fkey FOREIGN KEY (user_id)
REFERENCES app_user (id);
ALTER TABLE comment ADD CONSTRAINT comment_post_id_fkey FOREIGN KEY (post_id)
REFERENCES post (id);
ALTER TABLE likes ADD CONSTRAINT likes_app_user_id_fkey FOREIGN KEY (user_id)
REFERENCES app_user (id);




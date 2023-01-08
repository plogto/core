-- Types
CREATE TYPE background_color AS ENUM ('light', 'dim', 'dark');

CREATE TYPE credit_transaction_description_variable_type AS ENUM ('user', 'tag', 'ticket');

CREATE TYPE credit_transaction_status AS ENUM ('approved', 'pending', 'failed', 'canceled');

CREATE TYPE credit_transaction_template_name AS ENUM (
	'invite_user',
	'register_by_invitation_code',
	'approve_ticket'
);

CREATE TYPE credit_transaction_type AS ENUM ('order', 'transfer', 'commission', 'fund');

CREATE TYPE credit_transaction_type_name AS ENUM ('invite_user', 'register_by_invitation_code');

CREATE TYPE invited_user_status AS ENUM ('registered', 'posted');

CREATE TYPE notification_type_name AS ENUM (
	'welcome',
	'like_post',
	'reply_post',
	'like_reply',
	'accept_user',
	'follow_user',
	'mention_in_post'
);

CREATE TYPE post_status AS ENUM ('public', 'private');

CREATE TYPE primary_color AS ENUM ('blue', 'green', 'red', 'purple', 'orange', 'yellow');

CREATE TYPE ticket_status_type AS ENUM (
	'open',
	'closed',
	'approved',
	'solved',
	'accepted',
	'rejected'
);

CREATE TYPE user_role AS ENUM ('user', 'admin', 'super_admin');

-- Triggers
CREATE FUNCTION trigger_set_updated_at() RETURNS TRIGGER LANGUAGE plpgsql AS $$ BEGIN NEW.updated_at = now();
RETURN NEW;
END;
$$;

-- Connections
CREATE TABLE connections (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	follower_id uuid NOT NULL,
	following_id uuid NOT NULL,
	status INTEGER DEFAULT 1 NOT NULL,
	created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	deleted_at TIMESTAMP WITHOUT TIME ZONE
);

ALTER TABLE
	ONLY connections
ADD
	CONSTRAINT connections_pk PRIMARY KEY (id);

ALTER TABLE
	ONLY connections
ADD
	CONSTRAINT connections_fk0 FOREIGN KEY (following_id) REFERENCES users(id);

ALTER TABLE
	ONLY connections
ADD
	CONSTRAINT connections_fk1 FOREIGN KEY (follower_id) REFERENCES users(id);

-- Credit Transaction Description Variables
CREATE TABLE credit_transaction_description_variables (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	credit_transaction_info_id uuid NOT NULL,
	content_id uuid NOT NULL,
	KEY CHARACTER VARYING(150) NOT NULL,
	TYPE credit_transaction_description_variable_type NOT NULL,
	created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	deleted_at TIMESTAMP WITHOUT TIME ZONE
);

ALTER TABLE
	ONLY credit_transaction_description_variables
ADD
	CONSTRAINT credit_transaction_description_variables_pk PRIMARY KEY (id);

ALTER TABLE
	ONLY credit_transaction_description_variables
ADD
	CONSTRAINT credit_transaction_description_variables_fk0 FOREIGN KEY (credit_transaction_info_id) REFERENCES credit_transaction_infos(id);

-- Credit Transaction Infos
CREATE TABLE credit_transaction_infos (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	credit_transaction_template_id uuid,
	description TEXT,
	status credit_transaction_status DEFAULT 'pending' :: credit_transaction_status NOT NULL,
	created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	deleted_at TIMESTAMP WITHOUT TIME ZONE
);

ALTER TABLE
	ONLY credit_transaction_infos
ADD
	CONSTRAINT credit_transaction_infos_pk PRIMARY KEY (id);

ALTER TABLE
	ONLY credit_transaction_infos
ADD
	CONSTRAINT credit_transaction_infos_fk0 FOREIGN KEY (credit_transaction_template_id) REFERENCES credit_transaction_templates(id);

-- Credit Transaction Templates
CREATE TABLE credit_transaction_templates (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	NAME credit_transaction_template_name NOT NULL,
	amount REAL,
	CONTENT TEXT NOT NULL,
	deleted_at TIMESTAMP WITHOUT TIME ZONE
);

ALTER TABLE
	ONLY credit_transaction_templates
ADD
	CONSTRAINT credit_transaction_templates_pk PRIMARY KEY (id);

-- Credit Transactions
CREATE TABLE credit_transactions (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	user_id uuid NOT NULL,
	recipient_id uuid NOT NULL,
	relevant_credit_transaction_id uuid,
	credit_transaction_info_id uuid NOT NULL,
	amount REAL NOT NULL,
	url CHARACTER VARYING(24) NOT NULL,
	TYPE credit_transaction_type NOT NULL,
	created_at TIMESTAMP(6) WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	deleted_at TIMESTAMP(6) WITHOUT TIME ZONE
);

ALTER TABLE
	ONLY credit_transactions
ADD
	CONSTRAINT credit_transactions_pk PRIMARY KEY (id);

ALTER TABLE
	ONLY credit_transactions
ADD
	CONSTRAINT credit_transactions_fk0 FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE
	ONLY credit_transactions
ADD
	CONSTRAINT credit_transactions_fk1 FOREIGN KEY (recipient_id) REFERENCES users(id);

ALTER TABLE
	ONLY credit_transactions
ADD
	CONSTRAINT credit_transactions_fk2 FOREIGN KEY (relevant_credit_transaction_id) REFERENCES credit_transactions(id);

ALTER TABLE
	ONLY credit_transactions
ADD
	CONSTRAINT credit_transactions_fk3 FOREIGN KEY (credit_transaction_info_id) REFERENCES credit_transaction_infos(id);

-- Files
CREATE TABLE files (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	hash TEXT NOT NULL,
	NAME TEXT NOT NULL,
	width SMALLINT NOT NULL,
	height SMALLINT NOT NULL,
	created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	deleted_at TIMESTAMP WITHOUT TIME ZONE
);

ALTER TABLE
	ONLY files
ADD
	CONSTRAINT files_hash_key UNIQUE (hash);

ALTER TABLE
	ONLY files
ADD
	CONSTRAINT files_url_key UNIQUE (NAME);

ALTER TABLE
	ONLY files
ADD
	CONSTRAINT files_pk PRIMARY KEY (id);

-- Invited Users
CREATE TABLE invited_users (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	inviter_id uuid NOT NULL,
	invitee_id uuid NOT NULL,
	status invited_user_status DEFAULT 'registered' :: invited_user_status NOT NULL,
	created_at TIMESTAMP(6) WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	updated_at TIMESTAMP(6) WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	deleted_at TIMESTAMP(6) WITHOUT TIME ZONE DEFAULT now()
);

ALTER TABLE
	ONLY invited_users
ADD
	CONSTRAINT invited_users_pk PRIMARY KEY (id);

ALTER TABLE
	ONLY invited_users
ADD
	CONSTRAINT invited_users_fk0 FOREIGN KEY (inviter_id) REFERENCES users(id);

ALTER TABLE
	ONLY invited_users
ADD
	CONSTRAINT invited_users_fk1 FOREIGN KEY (invitee_id) REFERENCES users(id);

-- Liked Posts
CREATE TABLE liked_posts (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	user_id uuid NOT NULL,
	post_id uuid NOT NULL,
	created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	deleted_at TIMESTAMP WITHOUT TIME ZONE
);

ALTER TABLE
	ONLY liked_posts
ADD
	CONSTRAINT liked_posts_pk PRIMARY KEY (id);

ALTER TABLE
	ONLY liked_posts
ADD
	CONSTRAINT liked_posts_fk0 FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE
	ONLY liked_posts
ADD
	CONSTRAINT liked_posts_fk1 FOREIGN KEY (post_id) REFERENCES posts(id);

-- Notification Types
CREATE TABLE notification_types (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	NAME CHARACTER VARYING(100) NOT NULL,
	TEMPLATE TEXT NOT NULL,
	deleted_at TIMESTAMP WITHOUT TIME ZONE
);

ALTER TABLE
	ONLY notification_types
ADD
	CONSTRAINT notification_types_pk PRIMARY KEY (id);

-- Notifications
CREATE TABLE notifications (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	notification_type_id uuid NOT NULL,
	sender_id uuid NOT NULL,
	receiver_id uuid NOT NULL,
	post_id uuid,
	reply_id uuid,
	url TEXT NOT NULL,
	READ BOOLEAN DEFAULT FALSE NOT NULL,
	created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	deleted_at TIMESTAMP WITHOUT TIME ZONE
);

ALTER TABLE
	ONLY notifications
ADD
	CONSTRAINT notifications_pk PRIMARY KEY (id);

ALTER TABLE
	ONLY notifications
ADD
	CONSTRAINT notifications_fk0 FOREIGN KEY (notification_type_id) REFERENCES notification_types(id);

ALTER TABLE
	ONLY notifications
ADD
	CONSTRAINT notifications_fk1 FOREIGN KEY (sender_id) REFERENCES users(id);

ALTER TABLE
	ONLY notifications
ADD
	CONSTRAINT notifications_fk2 FOREIGN KEY (receiver_id) REFERENCES users(id);

ALTER TABLE
	ONLY notifications
ADD
	CONSTRAINT notifications_fk3 FOREIGN KEY (post_id) REFERENCES posts(id);

ALTER TABLE
	ONLY notifications
ADD
	CONSTRAINT notifications_fk4 FOREIGN KEY (reply_id) REFERENCES posts(id);

-- Online Users
CREATE TABLE online_users (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	user_id uuid NOT NULL,
	socket_id CHARACTER VARYING(20) NOT NULL,
	token TEXT NOT NULL,
	user_agent TEXT NOT NULL,
	created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	deleted_at TIMESTAMP WITHOUT TIME ZONE
);

ALTER TABLE
	ONLY online_users
ADD
	CONSTRAINT online_users_pk PRIMARY KEY (id);

ALTER TABLE
	ONLY online_users
ADD
	CONSTRAINT online_users_fk0 FOREIGN KEY (user_id) REFERENCES users(id);

-- Passwords
CREATE TABLE passwords (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	user_id uuid NOT NULL,
	PASSWORD TEXT NOT NULL,
	created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	deleted_at TIMESTAMP WITHOUT TIME ZONE
);

ALTER TABLE
	ONLY passwords
ADD
	CONSTRAINT passwords_pk PRIMARY KEY (id);

ALTER TABLE
	ONLY passwords
ADD
	CONSTRAINT passwords_fk0 FOREIGN KEY (user_id) REFERENCES users(id);

-- Post Attachments
CREATE TABLE post_attachments (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	post_id uuid NOT NULL,
	file_id uuid NOT NULL,
	created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	deleted_at TIMESTAMP WITHOUT TIME ZONE
);

ALTER TABLE
	ONLY post_attachments
ADD
	CONSTRAINT post_attachments_pk PRIMARY KEY (id);

ALTER TABLE
	ONLY post_attachments
ADD
	CONSTRAINT post_attachments_fk0 FOREIGN KEY (post_id) REFERENCES posts(id);

ALTER TABLE
	ONLY post_attachments
ADD
	CONSTRAINT post_attachments_fk1 FOREIGN KEY (file_id) REFERENCES files(id);

-- Post Tags
CREATE TABLE post_tags (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	tag_id uuid NOT NULL,
	post_id uuid NOT NULL,
	created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	deleted_at TIMESTAMP WITHOUT TIME ZONE
);

ALTER TABLE
	ONLY post_tags
ADD
	CONSTRAINT post_tags_pk PRIMARY KEY (id);

ALTER TABLE
	ONLY post_tags
ADD
	CONSTRAINT post_tags_fk0 FOREIGN KEY (tag_id) REFERENCES tags(id);

ALTER TABLE
	ONLY post_tags
ADD
	CONSTRAINT post_tags_fk1 FOREIGN KEY (post_id) REFERENCES posts(id);

-- Posts
CREATE TABLE posts (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	user_id uuid NOT NULL,
	parent_id uuid,
	child_id uuid,
	status post_status DEFAULT 'public' :: post_status NOT NULL,
	CONTENT TEXT,
	url CHARACTER VARYING(20) NOT NULL,
	created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	deleted_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE UNIQUE INDEX posts_url_key ON posts USING btree (url);

ALTER TABLE
	ONLY posts
ADD
	CONSTRAINT posts_pk PRIMARY KEY (id);

ALTER TABLE
	ONLY posts
ADD
	CONSTRAINT posts_fk0 FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE
	ONLY posts
ADD
	CONSTRAINT posts_fk1 FOREIGN KEY (parent_id) REFERENCES posts(id);

ALTER TABLE
	ONLY posts
ADD
	CONSTRAINT posts_fk2 FOREIGN KEY (child_id) REFERENCES posts(id);

-- Saved Posts
CREATE TABLE saved_posts (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	user_id uuid NOT NULL,
	post_id uuid NOT NULL,
	created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	deleted_at TIMESTAMP WITHOUT TIME ZONE
);

ALTER TABLE
	ONLY saved_posts
ADD
	CONSTRAINT saved_posts_pk PRIMARY KEY (id);

ALTER TABLE
	ONLY saved_posts
ADD
	CONSTRAINT saved_posts_fk0 FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE
	ONLY saved_posts
ADD
	CONSTRAINT saveed_posts_fk1 FOREIGN KEY (post_id) REFERENCES posts(id);

-- Tags
CREATE TABLE tags (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	NAME CHARACTER VARYING(100) NOT NULL,
	created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	deleted_at TIMESTAMP WITHOUT TIME ZONE
);

ALTER TABLE
	ONLY tags
ADD
	CONSTRAINT tags_name_key UNIQUE (NAME);

ALTER TABLE
	ONLY tags
ADD
	CONSTRAINT tags_pk PRIMARY KEY (id);

-- Ticket Message Attachments
CREATE TABLE ticket_message_attachments (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	ticket_message_id uuid NOT NULL,
	file_id uuid NOT NULL,
	created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	deleted_at TIMESTAMP WITHOUT TIME ZONE
);

ALTER TABLE
	ONLY ticket_message_attachments
ADD
	CONSTRAINT ticket_message_attachments_pk PRIMARY KEY (id);

ALTER TABLE
	ONLY ticket_message_attachments
ADD
	CONSTRAINT ticket_message_attachments_fk0 FOREIGN KEY (ticket_message_id) REFERENCES ticket_messages(id);

ALTER TABLE
	ONLY ticket_message_attachments
ADD
	CONSTRAINT ticket_message_attachments_fk1 FOREIGN KEY (file_id) REFERENCES files(id);

-- Ticket Messages
CREATE TABLE ticket_messages (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	ticket_id uuid NOT NULL,
	sender_id uuid NOT NULL,
	message TEXT NOT NULL,
	READ BOOLEAN DEFAULT FALSE NOT NULL,
	created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	deleted_at TIMESTAMP WITHOUT TIME ZONE
);

ALTER TABLE
	ONLY ticket_messages
ADD
	CONSTRAINT ticket_messages_pk PRIMARY KEY (id);

ALTER TABLE
	ONLY ticket_messages
ADD
	CONSTRAINT ticket_messages_fk0 FOREIGN KEY (ticket_id) REFERENCES tickets(id);

ALTER TABLE
	ONLY ticket_messages
ADD
	CONSTRAINT ticket_messages_fk1 FOREIGN KEY (sender_id) REFERENCES users(id);

-- Tickets
CREATE TABLE tickets (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	user_id uuid NOT NULL,
	subject TEXT NOT NULL,
	status ticket_status_type DEFAULT 'open' :: ticket_status_type NOT NULL,
	url CHARACTER VARYING(9),
	created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	deleted_at TIMESTAMP WITHOUT TIME ZONE
);

ALTER TABLE
	ONLY tickets
ADD
	CONSTRAINT tickets_pk PRIMARY KEY (id);

ALTER TABLE
	ONLY tickets
ADD
	CONSTRAINT tickets_fk0 FOREIGN KEY (user_id) REFERENCES users(id);

-- Users
CREATE TABLE users (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	username CHARACTER VARYING(32) NOT NULL,
	email CHARACTER VARYING(100) NOT NULL,
	full_name CHARACTER VARYING(64) NOT NULL,
	bio TEXT,
	ROLE user_role DEFAULT 'user' :: user_role NOT NULL,
	is_private BOOLEAN DEFAULT FALSE NOT NULL,
	avatar uuid,
	background uuid,
	primary_color primary_color DEFAULT 'blue' :: primary_color NOT NULL,
	background_color background_color DEFAULT 'light' :: background_color NOT NULL,
	is_verified BOOLEAN DEFAULT FALSE NOT NULL,
	invitation_code CHARACTER VARYING(7) NOT NULL,
	created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
	deleted_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE UNIQUE INDEX users_invitation_code_key ON users USING btree (invitation_code);

ALTER TABLE
	ONLY users
ADD
	CONSTRAINT users_email_key UNIQUE (email);

ALTER TABLE
	ONLY users
ADD
	CONSTRAINT users_username_key UNIQUE (username);

ALTER TABLE
	ONLY users
ADD
	CONSTRAINT users_pk PRIMARY KEY (id);

ALTER TABLE
	ONLY users
ADD
	CONSTRAINT users_fk0 FOREIGN KEY (avatar) REFERENCES files(id);

ALTER TABLE
	ONLY users
ADD
	CONSTRAINT users_fk1 FOREIGN KEY (background) REFERENCES files(id);

-- Triggers
CREATE TRIGGER update_connections BEFORE
UPDATE
	ON connections FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE TRIGGER update_credit_transaction_infos BEFORE
UPDATE
	ON credit_transaction_infos FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE TRIGGER update_invited_users BEFORE
UPDATE
	ON invited_users FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE TRIGGER update_liked_posts BEFORE
UPDATE
	ON liked_posts FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE TRIGGER update_notifications BEFORE
UPDATE
	ON notifications FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE TRIGGER update_online_users BEFORE
UPDATE
	ON online_users FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE TRIGGER update_passwords BEFORE
UPDATE
	ON passwords FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE TRIGGER update_post_tags BEFORE
UPDATE
	ON post_tags FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE TRIGGER update_posts BEFORE
UPDATE
	ON posts FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE TRIGGER update_saved_posts BEFORE
UPDATE
	ON saved_posts FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE TRIGGER update_tags BEFORE
UPDATE
	ON tags FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE TRIGGER update_ticket_messages BEFORE
UPDATE
	ON ticket_messages FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE TRIGGER update_tickets BEFORE
UPDATE
	ON tickets FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE TRIGGER update_users BEFORE
UPDATE
	ON users FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();
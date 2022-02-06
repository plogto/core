TRUNCATE TABLE IF EXISTS "notification_type";
INSERT INTO notification_type("name", "template")
VALUES
	('LIKE_POST','$$$___sender.username___$$$ liked your post.'),
	('REPLY_POST','$$$___sender.username___$$$ replied to your post: $$$___post.content___$$$.'),
	('LIKE_REPLY','$$$___sender.username___$$$ liked your reply.'),
	('FOLLOW_USER','$$$___sender.username___$$$ started following you.'),
	('ACCEPT_USER','$$$___sender.username___$$$ accepted your request.');
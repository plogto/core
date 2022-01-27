INSERT INTO notification_type("name", "template")
VALUES
	('LIKE_POST','$$$___sender.username___$$$ liked your post.'),
	('COMMENT_POST','$$$___sender.username___$$$ commented on your post: $$$___comment.content___$$$ '),
	('LIKE_COMMENT','$$$___sender.username___$$$ liked your comment.'),
	('FOLLOW_USER','$$$___sender.username___$$$ started following you.'),
	('ACCEPT_USER','$$$___sender.username___$$$ accepted your request.');
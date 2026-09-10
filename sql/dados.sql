INSERT INTO usuarios (nome, nick, email, senha) 
values 
("usuario 1", "usuario 1", "usuario1@gmail.com", "$2a$10$Lq.8Y2sN8L84HCYLzzCyZurQGafbghAwDtX1kFM3rU0jNK7m2viue"),
("Usuario 2", "usuario 2", "usuario2@gmail.com", "$2a$10$Lq.8Y2sN8L84HCYLzzCyZurQGafbghAwDtX1kFM3rU0jNK7m2viue"),
("Usuario 3", "usuario 3", "usuario3@gmail.com", "$2a$10$Lq.8Y2sN8L84HCYLzzCyZurQGafbghAwDtX1kFM3rU0jNK7m2viue");

INSERT INTO seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);

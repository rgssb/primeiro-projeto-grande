INSERT INTO usuarios (nome, nick, email, senha) 
values 
("usuario 1", "usuario 1", "usuario1@gmail.com", "$2a$10$tTHu0p5L7Hv0vpQNkz7HPu9QAheahu5kh0FjN1P9RyO0bhqceCW1e"),
("Usuario 2", "usuario 2", "usuario2@gmail.com", "$2a$10$tTHu0p5L7Hv0vpQNkz7HPu9QAheahu5kh0FjN1P9RyO0bhqceCW1e"),
("Usuario 3", "usuario 3", "usuario3@gmail.com", "$2a$10$tTHu0p5L7Hv0vpQNkz7HPu9QAheahu5kh0FjN1P9RyO0bhqceCW1e");

INSERT INTO seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicação do Usuaro 1", "Essa é a publicação do usuario 1, Iupi!!!", 1),
("Publicação do Usuaro 2", "Essa é a publicação do usuario 2, Iupi!!!", 2),
("Publicação do Usuaro 3", "Essa é a publicação do usuario 3, Iupi!!!", 3);
package repositorios

import (
	"API/src/modelos"
	"database/sql"
)

// Publicacoes representa um repositorio de publicacoes
type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um novo repositorio de publicacoes
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db: db}
}

// Criar insere uma publicacao no banco de dados
func (repositorio *Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into publicacoes (titulo, conteudo, autor_id) VALUES (?, ?, ?)",
	)

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoID, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoID), nil
}

// BuscarPorID busca uma publicacao por ID
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	var publicacao modelos.Publicacao
	erro := repositorio.db.QueryRow(`
		select p.*, u.nick from
		publicacoes p inner join usuarios u
		on u.id = p.autor_id where p.id = ?`,
		publicacaoID,
	).Scan(
		&publicacao.ID,
		&publicacao.Titulo,
		&publicacao.Conteudo,
		&publicacao.AutorID,
		&publicacao.Curtidas,
		&publicacao.CriadaEm,
		&publicacao.AutorNick,
	)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}

	return publicacao, nil

}

// Busca traz as publucacoes dos usuarios seguidos e tambem do proprio usuario que fez a requisicao
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT DISTINCT p.*, u.nick
		FROM publicacoes p
		INNER JOIN usuarios u ON u.id = p.autor_id
		LEFT JOIN seguidores s ON p.autor_id = s.usuario_id
		WHERE u.id = ? OR s.seguidor_id = ?
		order by 1 desc`,
		usuarioID, usuarioID,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	if erro = linhas.Err(); erro != nil {
		return nil, erro
	}

	return publicacoes, nil
}

//Atualizar altera a publicacao de um usuario
func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare(
		"update publicacoes set titulo = ?, conteudo = ? where id = ?",
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui uma publicacao do banco de dados
func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"delete from publicacoes where id = ?",
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

//BuscarPorUsuario busca todas as publicacoes de um usuario no banco de dados
func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		select p.*, u.nick from publicacoes p
		join usuarios u on u.id = p.autor_id
		where p.autor_id = ?`, usuarioID)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	if erro = linhas.Err(); erro != nil {
		return nil, erro
	}

	return publicacoes, nil
}

//Curtir adiciona uma curtida no poste de um usuario
func (repositorio Publicacoes) Curtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"update publicacoes set curtidas = curtidas + 1 where id = ?",
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}


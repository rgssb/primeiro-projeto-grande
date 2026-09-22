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

package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectandoComBanco() *sql.DB {
	stringDeConexao := "user=postgres dbname=Alura_Loja password=root host=localhost sslmode=disable" // Varivel que irá armazena a string de conecção com banco

	db, err := sql.Open("postgres", stringDeConexao) // aqui estou abrindo minha conexão, caso de erro  a mesangem vai para variavel err caso de certo vai para varival db

	if err != nil { // aqui estou verificando se deu erro
		panic(err.Error()) // retornando a mensagem do erro.
	}

	return db // caso de tudo certo eu retorno

}

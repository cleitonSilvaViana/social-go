/*
	Author: Cleiton Viana
*/

package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/cleitonSilvaViana/social-go/api/router"
	"github.com/cleitonSilvaViana/social-go/config"
)

var (
	Logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
)

func main() {
	slog.SetDefault(Logger)

	config.GetEnv()

		slog.LogAttrs(
		context.Background(),
		slog.LevelInfo, "start application",
		slog.String("port", "http://localhost:"+config.API_PORT),
	)


	router.InitRouter(config.API_PORT)
}

/*

log de requisição:
* rota acessada
* ip do cliente
* método http
* header http


Log de response:
* código de status
* cabeçalho da resposta
* Tempo de resposta: tempo decorrido desde o recebimento da requisição até o envio da resposta.


log de desempenho:
* Tempo de processamento: tempo decorrido para processar a lógica da API.
* Uso de memória: quantidade de memória alocada durante o processamento da requisição.
* Latência do banco de dados: tempo decorrido para realizar operações no banco de dados.


log de erro:
* stack trace
* mensagem de erro
* Tipo de erro: se é um erro de servidor, erro de validação de entrada, etc.


log banco de dados:
* incluir a consulta SQL
* o tempo de execução da consulta


Logs de segurança:

* Tentativas de login falhas: registrar tentativas de login com dados inválidos.
* Acessos a endpoints confidenciais: registrar quais usuários acessaram endpoints que exigem autorização especial.
* Atividades suspeitas: registrar qualquer atividade que possa indicar um ataque ou comportamento malicioso.


*/

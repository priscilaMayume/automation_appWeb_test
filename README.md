# automation_appWeb_test
automation_appWeb_test_golang

# Projeto de Testes Automatizados - API com Go

Este é um projeto de testes automatizados utilizando Go para testar uma API externa (ReqRes). O objetivo principal do projeto é validar as operações de login e criação de usuários, utilizando testes funcionais para garantir que a API está funcionando conforme o esperado.

## Tecnologias Utilizadas

- **Go**: Linguagem de programação utilizada para escrever os testes.
- **Resty**: Biblioteca para facilitar a comunicação HTTP.
- **Testify**: Framework para asserções em testes Go.
- **JSON**: Para a manipulação de respostas e solicitações da API.

## Objetivos do Projeto

- Testar a API ReqRes, realizando requisições para criar e obter usuários.
- Validar as respostas da API em termos de status HTTP e conteúdo da resposta.

## Estrutura do Projeto


├── internal
│ ├── api 
 │ │└── api_service.go # Funções de interação com a API 
 │ └── steps 
  │ └── login_steps.go # Definições dos testes e validações 
  ├── tests 
  │ └── login_test.go # Testes funcionais para login
   ├── go.mod # Dependências do Go
    └── README.md # Documentação do projeto


## Pré-requisitos

Antes de executar o projeto, você precisará configurar o ambiente de desenvolvimento e garantir que todas as dependências estejam instaladas.

### 1. Instalar o Go

Certifique-se de que o Go está instalado em sua máquina. Caso não tenha, siga as instruções abaixo:

brew install go

 Executando os Testes
O projeto foi configurado para realizar testes de integração com a API ReqRes. Para executar os testes, siga os passos abaixo:

Execute os testes utilizando o comando go test:

go test ./... -v

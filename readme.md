# API RESTful de Posts em Golang

Este projeto é uma API RESTful simples para gerenciar posts. Os posts são representados por um ID, um título e o conteúdo do corpo (body). A API é desenvolvida em Golang (Go) e utiliza a biblioteca "gorilla/mux" para lidar com roteamento de requisições HTTP.

## Funcionalidades

A API oferece as seguintes funcionalidades para interagir com os posts:

1. **Listar todos os posts**: Retorna todos os posts existentes.

2. **Obter um post específico**: Retorna as informações de um post específico com base em seu ID.

3. **Criar um novo post**: Cria um novo post com base nos dados fornecidos na requisição.

4. **Atualizar um post**: Atualiza um post existente com base no ID fornecido na requisição.

5. **Excluir um post**: Exclui um post existente com base no ID fornecido na requisição.

## Endpoints

A API possui os seguintes endpoints:

- `GET /posts`: Retorna todos os posts existentes.
- `POST /posts`: Cria um novo post.
- `GET /posts/{id}`: Retorna as informações de um post específico.
- `PUT /posts/{id}`: Atualiza um post existente.
- `DELETE /posts/{id}`: Exclui um post existente.

## Estrutura do Post

Um post é representado pelo seguinte JSON:

```json
{
    "id": "1",
    "title": "My first post",
    "body": "This is the content of my first post"
}
```

## Como Executar

Certifique-se de ter o Go instalado em seu sistema. Você pode obtê-lo em https://golang.org/.

Para executar a API, navegue até o diretório onde se encontra o arquivo `main.go` e execute o seguinte comando:

```bash
go run main.go
```

A API será executada localmente na porta 8000. Você pode acessar os endpoints utilizando ferramentas como cURL, Postman ou qualquer cliente HTTP.

## Aviso

Este projeto é apenas uma demonstração simples de como criar uma API RESTful básica em Golang. Não foi desenvolvido com recursos de segurança ou autenticação. Recomendamos adaptá-lo e melhorá-lo de acordo com as necessidades do seu projeto antes de colocá-lo em produção.

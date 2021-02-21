# Project Name
> carrinho-api-gorm

## Tecnologias

* [GoLang](https://golang.org/) - compilador da linguagem Go
* [Go Mod](https://github.com/golang/mod) - gerenciador de dependencias

## Instalação

Instalando dependências
```bash
$ go get
```

Removendo dependencias indesejadas
```bash
$ go mod tidy
```

Baixando as dependencias para a vendor local
```bash
$ go mod vendor
```

### Docker

Depois de instalado versão do `Docker` (windows, linux ou mac), irá executar o comando `docker --version` para verificar se o mesmo foi instalado.

Acessar a pasta raiz do projeto e executar o comando `docker-compose up -d`

Quando quiser parar o processo executar o comando `docker-compose down`

### Mysql

Conectar no `Mysql` via shell pelo comando: `docker exec -it mysql bash` e depois `mysql -h localhost -u root -p`

Informar senha do `root` para acessar o `Mysql` e aplicar o comando `create database carrinho;` para criar a base de dados.

## Iniciando

Buildando o projeto
```bash
# execute o comando abaixo para buildar a aplicação e garantir que nada está quebrado
$ go build
```

Executando o projeto
```bash
$ go run main.go or ./carrinho-api-gorm
```
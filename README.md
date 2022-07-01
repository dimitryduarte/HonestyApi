# HonestyApi

Este projeto desenvolvido em GoLang tem o intuito de fornecer uma API RESTFul de cadastro de produtos e usuários, produto projetado para atender demandas de compras pelo app de pequenas lojas internas dentro de empresas.

O intuito é o usuário consumir, e ele mesmo dar baixa do que comprou via app.

LIBRARIES
-----------------
Foram utilizados para este projeto os seguintes recursos:

- godotenv (Integração com arquivo .ENV)
- mux (Roteamento)
- gorm (ORM)
- driver/mysql (Comunicação com o banco)

PATTERN
------------------
O pattern utilizado para este projeto foi uma mistura entre Go Clean e o tão conhecido MVC

ENVIRONMENTS
------------------
Crie um arquivo na raíz do projeto com o nome ``.env`` e configure as variáveis de ambiente do projeto, se preferir você pode utilizar o arquivo ```.env.example``` como exemplo. As variáveis são:
```
API_SECRET=ApiS3cret
DB_HOST=127.0.0.1
DB_DRIVER=mysql
DB_USER=dbuser
DB_PASSWORD=123456
DB_NAME=honestyapp
DB_PORT=3306
PORT=8000
```
INSTRUÇÕES
---------------------
Após configurar as variáveis, para iniciar o projeto diretamente pelo VSCODE execute: 
```go run main.go```

Caso queira gerar um executável:
```
go build
```
E após isso, basta inicializar o ```.exe``` :
```
./HonestyApi.exe
```

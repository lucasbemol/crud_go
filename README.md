# Crud_Go
Esta aplicação se trata de um CRUD feito em Golang por mim, Lucas Martins Ramos. A Aplicação permite visualizar todos os Produto, Adicionar Produto, Editar Produto, Visualizar produtos individuais, Buscar por nome do Produto e Excluir um produto.

# Instruções
Para executar a aplicação, precisamos baixar as dependências, com o GO instalado e configurado em seu computador execue os comandos:
go get github.com/codegangsta/martini
go get github.com/codegangsta/martini-contrib/binding
go get github.com/codegangsta/martini-contrib/render

go get github.com/go-sql-driver/mysql

Para configurar o acesso ao Banco de dados MySQL precisamos editar o arquivo crud.go.
(linha 30), altere as configurações para seu database, a configuração defalt é: db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/crud")
altere o usuário e senha conforme as configurações de seu database, ex: sql.Open("mysql", "root:lucasbemol@tcp(localhost:3306)/crud")
A Aplicação vai criar o schema crud e tabela Product ao iniciar.

#Tecnologias

* Golang 1.5.1
* AngularJS 1.3.4
* Bootstrap 3
* Mysql

#Golang:
Para construção do REST utilizei a API 'Martini' para o Golang, com ela consigo mapear de forma simples os Endpoints, muito parecidos com os controllers Spring.
OBS: Poderia utilizar a API Gorila também, depois que construi a aplicação andei estudando as 2 API's e acredito que a Gorila possui mais vantagens de controle.
Os testes unitários ainda não estão utilizando Mock, porem se dividirmos melhor(em services) nossos método podemos implementar os mocks.
Para controle do banco de dados utilizamos o driver do MySQL e a própria API de DB do Golang.

#AngularJS:
Toda interação entre front-end e back-end é aessada pelo Angular.
* Utilizei Factory's para construir todo o acesso ao webservice, padronizando todas as chamadas e reaproveitando quando necessário.
* Utilizei as rotas($Router) do angular para mapear todas as paginas com seus endereços, assim temos o front-end e back-end independentes, as rotas do Angular funcionam muito bem e são simples para configurar.
* Para os furmulários utilizei o "Required" do Angular para obrigar o preenchimento dos campos, evitando assim erros desnecessários.
* Para as mascaras nos input's utilizei Angular UI Mask, com essa diretiva conseguimos colocar mascaras de forma muito simples em qualque parte do front-end
* Cada pagina possúi seu próprio controller Angular, com o Service de ações sendo injetado neles, dessa forma os controllers ficam pequenos, simples, e muito fácil para entender e dar manutenções futuras.

#Bootstrap:
* Utilizei formulários simples do Bootstrap, são simples para implementar, com um bom design e responsivo por padrão.
* Para inputs com Data, utilizei a Api Bootstrap-datapicker, com ela conseguimos configurar o formato da data e o idioma(pt-BR em nosso caso) do calendário.

#MySQL:
* Como dito nas instruções iniciais, precisamos apenas configurar nosso usuário e senha de acesso ao banco na aplicação. Ao inicar a aplicação irá construir o database e tabela para seu uso.

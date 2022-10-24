# Go-API-CRUD
Esta é uma API desenvolvida em Golang que se comunica com um banco de dados MySQL, além disso, também há a preocupação em separar em containers, servidor e banco de dados.
# Uso
## Banco MySQL
É necessario ter um banco/container pré-definido:
1. Instale uma imagem do container MySQL:
```
docker pull mysql
```
2. Crie um container e deixe em segundo plano:
```
docker container run --name go-api-mysql -e MYSQL_ROOT_PASSWORD=250721 --ip "172.17.0.2" -d mysql:8
```
A porta é a padrão: 3306

3. Inicie o banco 
```
docker container exec -it go-api-mysql bash
```
```
bash-4.4# mysql -h 172.17.0.2 -u root -p
```
```
Enter password: 250721
```
```
mysql> create database if not exists suzana_motorcycle_parts;
```
```
mysql> create table if not exists Part (
  id int not null auto_increment,
  name varchar(500) not null unique,
  brand varchar(50) not null,
  value float not null,
  primary key (id)
);
```
```
insert into Part (name, brand, value) values 
  ("Luva para Motociclista Dedo Longo Tam. P Material Emborrachado e Couro, Branco/ Preto", "Multilaser", 47.64),
  ("Capacete Moto R8 Pro Tork 56 Viseira Fume Preto/Vermelho", "Tork", 169.90),
  ("Lenço de cabeça, Romacci Máscara facial Fleece máscara facial cachecol para exterior à prova de vento à prova de frio equipamento de equitação para máscara de inverno", "Romacci", 99.19);
```
Para sair do terminal basta usar Ctrl + d
## Servidor
1. Build o Dockerfile
```
docker image build . -t go-api:1.0
```
2. Consulte a id da sua imagem de go-api
```
docker image ls
```
3. Crie um container go-api
```
docker container run --name go-api --ip "172.17.0.3" -d cc59a182cd15
```
Pronto, o server já está rodando.
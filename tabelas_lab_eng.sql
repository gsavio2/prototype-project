create table usuario (
  id                number primary key
 ,username          varchar2(80)
 ,senha             varchar2(30)
 ,tipo_usuario_id   number
 ,pessoa_id         number not null
);

create table tipo_usuario (
  id        number primary key
 ,descricao varchar2(40)
 ,status    varchar2(30)
  );
  
create table pessoa (
  id              number primary key
 ,nome            varchar2(200)
 ,email           varchar2(200)
 ,telefone        number
 ,celular         number
 ,cep             number
 ,cpf             number
 ,tipo_pessoa_id  number
);

create table evento (
  id                  number primary key
 ,tipo_evento_id      number
 ,descricao           varchar2(200)
 ,status              varchar2(30)
 ,data_execucao       date
 ,data_criacao        date
 ,data_final          date
 ,usuario_id_criacao  number
);

create table tipo_evento (
  id        number primary key
 ,descricao varchar2(50)
);

create table produto (
  id                number primary key
 ,descricao         varchar2(80)
 ,grupo_produto_id  number
 ,status            varchar2(30)
);

create table grupo_produto (
  id    number primary key
 ,descricao varchar2(40)
);
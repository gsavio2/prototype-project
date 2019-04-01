create or replace procedure insere_usuario(p_username     in varchar2
                                          ,p_senha        in varchar2
                                          ,p_tipo_usuario in varchar2
                                          ,p_pessoa_id    in number) is
                          
begin
  
  insert into usuario values (seq_usuario.nextval, p_username, p_senha, p_tipo_usuario, p_pessoa_id);
  
exception
  when others then
    raise_application_error(-20001, 'ERRO AO INSERIR USUÁRIO. ERRO: ' || sqlerrm);
end;


create or replace function recupera_senha_usuario(p_username  in varchar2
                                                 ,p_pessoa_id in number) return varchar2 is
  w_retorno  varchar2(100);
begin
  select senha
    into w_retorno
    from usuario
   where username   = p_username
     and pessoa_id  = p_pessoa_id;
     
  return w_retorno;
  
exception
  when no_data_found then
    raise_application_error(-20001, 'NÃO ENCONTRADO USUÁRIO COM ESTE USERNAME E PESSOA.');
  when others then
    raise_application_error(-20001, 'OTHERS NA RECUPERAR_SENHA_USUARIO. ERRO: ' || sqlerrm);
end;
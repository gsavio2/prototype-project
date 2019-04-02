<template>
<div>
<h1>Informe seus dados</h1>
<div class="form-register">
    <b-form @submit="cadastraUsuario" @reset="limpaForm" v-if="show">
      <b-form-group
        id="emailGroup"
        label="E-mail:"
        label-for="emailInput"
        description="Seu e-mail não será compartilhado com ninguém."
      >
        <b-form-input
          id="emailInput"
          type="email"
          v-model="user.email"
          required
          placeholder="exemplo@brejaria.com" />
      </b-form-group>

      <b-form-group id="nomeGroup" label="Nome:" label-for="nomeInput">
        <b-form-input
          id="nomeInput"
          type="text"
          v-model="user.nome"
          required
          placeholder="Digite seu nome" />
      </b-form-group>

      <b-form-group id="senhaGroup" label="Senha:" label-for="senha">
        <b-form-input
          id="senha"
          type="password"
          v-model="user.senha"
          required
          placeholder="Digite sua senha" />
      </b-form-group>

      <b-form-group id="interessesGroup" label="Interesses:">
        <b-form-checkbox-group v-model="user.perfil" id="interessesCheck">
          <b-form-checkbox value="1">Anunciar</b-form-checkbox>
          <b-form-checkbox value="2">Consumir</b-form-checkbox>
        </b-form-checkbox-group>
      </b-form-group>

      <b-button class="btn btn-md btn-primary btn-block" type="submit" variant="dark">Enviar</b-button>
      <b-button class="btn btn-md btn-primary btn-block" type="reset" variant="danger">Limpar</b-button>
 
    </b-form>
  </div>
  </div>
</template>

<script>

  import Usuario from "../services/user";

  export default {
    data() {
      return {
        user: {
          email: '',
          nome: '',
          perfil: [],
          senha: '',
        },
        show: true
      }
    },
    methods: {
      cadastraUsuario(evt) {
        evt.preventDefault()
        //alert(JSON.stringify(this.user))
        if((this.user.perfil).length >= 1) {
          Usuario.cadastraUsuario(this.user).then(resposta =>{
            this.user = {}
            alert("Salvo com sucesso")
            this.$router.push({name: 'perfil'})
          })
        } else {
          alert("Selecionar ao menos uma opção de Interesses")
        }
      },
      limpaForm(evt) {
        evt.preventDefault()
        /* limpa todos os campos */
        this.user.email = ''
        this.user.nome = ''
        this.user.senha = ''
        this.user.perfil = []
        /* limpa validações do navegador referente aos campos */
        this.show = false
        this.$nextTick(() => {
          this.show = true
        })
      }
    }
  }
</script>

<style scoped>
  .form-register {
    width: 100%;
    max-width: 400px;
    padding: 15px;
    margin: auto;
  }
  .form-register .checkbox {
    font-weight: 400;
  }
  .form-register .form-control {
    position: relative;
    box-sizing: border-box;
    height: auto;
    padding: 10px;
    font-size: 16px;
  }
  .form-signin .form-control:focus {
  z-index: 2;
  }
</style>




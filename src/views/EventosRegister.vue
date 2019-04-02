<template>
<div class="produtos-register">
    <h1>Informe os dados do seu evento</h1>
    <div class="form-register">
     <b-form @submit="cadastraEvento" @reset="limpaForm" v-if="show">
 <b-form-group
        id="nomeGroup"
        label="Nome: "
        label-for="nomeInput"
      >
        <b-form-input
          id="nomeInput"
          type="text"
          v-model="evento.nome"
          required
          placeholder="Digite o nome do seu evento" />
      </b-form-group>

      <b-form-group
        id="descricaoGroup"
        label="Descrição: "
        label-for="descricaoInput"
      >
        <b-form-textarea
          id="descricaoInput"
          type="text"
          v-model="evento.descricao"
          required
          placeholder="Faça as pessoas se interessarem pelo seu evento" />
      </b-form-group>

      <b-form-group id="criadorGroup" label="Criador:" label-for="criadorInput">
        <b-form-input
          id="criadorInput"
          type="text"
          v-model="evento.criador"
          required
          placeholder="Digite seu nome" />
      </b-form-group>

      <b-form-group id="localGroup" label="Local:" label-for="localInput">
        <b-form-input
          id="localInput"
          type="text"
          v-model="evento.local"
          required
          placeholder="Digite o local do evento" />
      </b-form-group>

      <b-form-group id="dataEventoGroup" label="Data do evento:" label-for="dataEventoInput">
        <b-form-input
          id="dataEventoInput"
          type="text"
          v-model="evento.dataEvento"
          required
          placeholder="Digite a data (dd/mm/aaaa hh:mm)" />
      </b-form-group>
    <div class="butao">
      <b-button class="btn btn-md btn-primary btn-block" type="submit" variant="dark">Enviar</b-button>
      <b-button class="btn btn-md btn-primary btn-block" type="reset" variant="danger">Limpar</b-button>
    </div>
    </b-form>
    </div>
</div>
</template>

<script>
import Evento from "../services/evento";

export default {
    name: 'eventoRegister',
    data() {
      return {
        evento: {
            nome: '',
            descricao: '',
            local: '',
            dataEvento: '',
            criador: '',
        },
        show: true
      }
    },
    methods: {
      cadastraEvento(evt) {
        evt.preventDefault()
        alert(JSON.stringify(this.evento))
          Evento.cadastraEvento(this.evento.criador, this.evento).then(resposta =>{
            this.evento = {}
            alert("Salvo com sucesso")
            this.$router.push({path: '/eventos'})
          })
      },
      limpaForm(evt) {
        evt.preventDefault()
        /* limpa todos os campos */
        this.evento.descricao = ''
        this.evento.dono = ''
        this.evento.nome = ''
        this.evento.local = ''
        this.evento.dataEvento = ''
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
  .butao {
    width: 100%;
    margin: auto;
  }
</style>

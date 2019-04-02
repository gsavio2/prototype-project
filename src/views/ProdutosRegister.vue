<template>
<div class="produtos-register">
    <h1>Informe os dados do seu produto</h1>
    <div class="form-register">
     <b-form @submit="cadastraProduto" @reset="limpaForm" v-if="show">
 <b-form-group
        id="nomeGroup"
        label="Nome: "
        label-for="nomeInput"
      >
        <b-form-input
          id="nomeInput"
          type="text"
          v-model="produto.nome"
          required
          placeholder="Digite o nome do seu produto" />
      </b-form-group>

      <b-form-group
        id="descricaoGroup"
        label="Descrição: "
        label-for="descricaoInput"
      >
        <b-form-textarea
          id="descricaoInput"
          type="text"
          v-model="produto.descricao"
          required
          placeholder="Faça as pessoas se interessarem pelo seu produto" />
      </b-form-group>

      <b-form-group id="donoGroup" label="Dono:" label-for="donoInput">
        <b-form-input
          id="donoInput"
          type="text"
          v-model="produto.dono"
          required
          placeholder="Digite seu nome" />
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
import Produto from "../services/produto";

export default {
    name: 'produtosRegister',
    data() {
      return {
        produto: {
            nome: '',
            descricao: '',
            dono: '',
        },
        show: true
      }
    },
    methods: {
      cadastraProduto(evt) {
        evt.preventDefault()
        alert(JSON.stringify(this.produto))
          Produto.cadastraProduto(this.produto.dono, this.produto).then(resposta =>{
            this.produto = {}
            alert("Salvo com sucesso")
            this.$router.push({name: 'produtos'})
          })
      },
      limpaForm(evt) {
        evt.preventDefault()
        /* limpa todos os campos */
        this.produto.descricao = ''
        this.produto.dono = ''
        this.produto.nome = ''
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

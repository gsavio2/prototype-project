<template>
  <div class="eventos">
    <h1>This is an eventos page</h1>
    <p>
      Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.
    </p>
    <b-button to="/eventos/register">Cadastre seu evento</b-button><!--vvvvvvvvvvv-->
    <b-table striped hover :items="eventos" :fields="fields" :sort-by.sync="sortBy">
           <template slot="acoes" slot-scope="row">
             <div class="acoes">  
               <b-button  >Excluir</b-button>
               <b-button  >Alterar</b-button>
               <b-button @click="adicionaParticipante(row.item, row.index)" v-model="row.evento">Participar</b-button>
               <b-button @click="removeParticipante(row.item, row.index)" v-model="row.evento">Sair</b-button>
              </div>
           </template>
        </b-table>
  </div>
</template>

<script>
import Evento from "../services/evento";

export default {
  data() {
    return {
      sortBy: 'ID',
      evento: {
        ID: '',
        nome: '',
        descricao: '',
        local: '',
        dataCriacao: '',
        criador: '',
        dataEvento: '',
        status: '',
        participantes: [],
      },
      eventos: [],
      fields: [{key: "ID", sortable: true}, 'Nome', 'Descricao', 'Local', 'Criador', 'DataCriacao', 'DataEvento', 'Status', 'Participantes', 'acoes'],
    }
  },
  mounted(){
    this.listaTodosEvento()
  },

  methods:{
    listaTodosEvento(){
      Evento.listaTodosEvento().then(resposta => {
        this.eventos = resposta.data
      })
    },

    adicionaParticipante(item, index){
      var idParticipante = prompt("Digite o id do participante")
      alert(JSON.stringify(item))
      alert(JSON.stringify(item.ID))
      alert(JSON.stringify(idParticipante))
      Evento.adicionaParticipante(item.ID, idParticipante, item).then(resposta =>{
        item = {}
        alert("Salvo com sucesso")
        this.listaTodosEvento()
      })
    },

    removeParticipante(item, index){
      var idParticipante = prompt("Digite o id do participante")
      alert(JSON.stringify(item))
      alert(JSON.stringify(item.ID))
      alert(JSON.stringify(idParticipante))
      Evento.removeParticipante(item.ID, idParticipante, item).then(resposta =>{
        item = {}
        alert("Removido com sucesso")
        this.listaTodosEvento()
      })
    }
  }
}
</script>

<style scoped>
</style>



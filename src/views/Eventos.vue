<template>
  <div class="eventos">
    <h1>This is an eventos page</h1>
    <p>
      Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.
    </p>
    <b-button to="/eventos/register">Cadastre seu evento</b-button><!--vvvvvvvvvvv-->
    <b-table striped hover :items="eventos" :fields="fields" v-bind="evento" :key="id">
           <template slot="acoes">
             <div class="acoes">
               <b-button  >Excluir</b-button>
               <b-button  >Alterar</b-button>
               <b-button v-on:click="adicionaParticipante" v-model="idParticipante">Participar</b-button>
              </div>
           </template>
           <!--<template v-for="(field, index) in fields" :slot="field.key" slot-scope="evento"></template>-->
        </b-table>
  </div>
</template>

<script>
import Evento from "../services/evento";

export default {
  data() {
    return {
      evento: {
        id: '',
        nome: '',
        descricao: '',
        local: '',
        dataCriacao: '',
        dataEvento: '',
        status: '',
        participantes: [],
      },
      idParticipante: '',
      eventos: [],
      fields: [{key: 'ID', colType: "id"}, {key: 'Nome'}, {key: 'Descricao'}, {key: 'Local'}, {key: 'DataCriacao'}, {key: 'DataEvento'}, {key: 'Status'}, {key: 'Criador'}, {key: 'Participantes'}, 'acoes'],
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
        adicionaParticipante: function(event){
            this.idParticipante = prompt("Digite o id do participante")
            Evento.adicionaParticipante(this.evento.ID, this.idParticipante, this.evento).then(resposta =>{
                this.evento = {}
                alert("Salvo com sucesso")
                this.listaTodosEvento()
            })
        }
    }


}
</script>

<style scoped>
</style>



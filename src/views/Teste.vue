<template>
    <div class="teste">
        {{user.name}}
        <b-form class="form-teste" @submit.prevent="postUser">
            <label>Name</label>
            <input type="text" class="form-control" placeholder="Name" v-model="user.name">
            <label>Email</label>
            <input type="text" class="form-control" placeholder="Email" v-model="user.email">
            <b-button type="submit">Enviar</b-button>
        </b-form>
        <b-table striped hover :items="users" :fields="fields">
           <template slot="acoes">
               <b-button>Excluir</b-button>
           </template>
        </b-table>
    </div>
</template>

<script>
import Usuario from "../services/teste";

export default {
    data() {
        return {
            
            user: {
                name: '',
                email: '',
            },
            users: [],
            fields: ['Name', 'Email', 'acoes'],
        }
    },
    mounted(){
        this.getUser()
    },

    methods:{
        getUser(){
            Usuario.getUser().then(resposta => {
                this.users = resposta.data
            })
        },
        postUser(){
            Usuario.postUser(this.user).then(resposta =>{
                this.user = {}
                alert("Salvo com sucesso")
                this.getUser()
            })
        }
    }
}
</script>

<style>
    .form-teste{
        width: 100%;
        max-width: 330px;
        padding: 15px;
        margin: auto;
    }
    .form-signin .form-control {
        position: relative;
        box-sizing: border-box;
        height: auto;
        padding: 10px;
        font-size: 16px;
    }
</style>

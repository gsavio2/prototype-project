<template>
    <div class="teste">
        {{user.name}}
        <b-form class="form-teste" @submit.prevent="post">
            <label>Name</label>
            <input type="text" class="form-control" placeholder="Name" v-model="user.name">
            <label>Email</label>
            <input type="text" class="form-control" placeholder="Email" v-model="user.email">
            <b-button type="submit">Enviar</b-button>
        </b-form>
        <tr v-for="user of users" :key="user.email">
            <td>a
                {{user.name}}
            </td>
            <td>
                {{user.email}}
            </td>
        </tr>
    </div>
</template>

<script>
import Usuario from "../services/teste";

export default {
    data(){
        return {
            user:{
                name: '',
                email: '',
            },
            users:[]
        }
    },
    mounted(){
        Usuario.getUser().then(resposta => {
            console.log(resposta)
        })
    },
    methods:{


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

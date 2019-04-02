<template>
<div class="login">
    <form class="form-signin">
        <img class="mb-4" src="../assets/logo.png" alt="" width="72" height="72">
        <h1 class="h3 mb-3 font-weight-normal">Dados de acesso</h1>
        <label for="inputEmail" class="sr-only">E-mail</label>
        <input type="text" id="inputEmail" v-model="login.email" class="form-control" placeholder="E-mail" required autofocus>
        <label for="inputPassword" class="sr-only">Senha</label>
        <input type="password" id="inputPassword" v-model="login.senha" class="form-control" placeholder="Senha" required>
        <!--<div class="checkbox mb-3">
            <label>
            <input type="checkbox" value="remember-me"> Remember me
            </label>
        </div>-->
        <b-button class="btn btn-lg btn-primary btn-block" variant="dark" type="submit" v-on:click="login">Entrar</b-button>
        <small>
          <router-link to="/register">Registre-se</router-link>
        </small>
    </form>
</div>
</template>

<script>
export default {
    name: 'login',
    data(){
        return{
            login: {
              email: '',
              senha: '',
            }
        }
    },
    methods: 
        {
            
        }
}
</script>

<style scoped>
.form-signin {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: auto;
}
.form-signin .checkbox {
  font-weight: 400;
}
.form-signin .form-control {
  position: relative;
  box-sizing: border-box;
  height: auto;
  padding: 10px;
  font-size: 16px;
}
.form-signin .form-control:focus {
  z-index: 2;
}
.form-signin input[type="text"] {
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}
.form-signin input[type="password"] {
  margin-bottom: 10px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}

</style>


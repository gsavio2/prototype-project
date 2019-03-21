<template>
<div class="form-register">
    <b-form class="" @submit="onSubmit" @reset="onReset" v-if="show">
      <b-form-group
        id="exampleInputGroup1"
        label="E-mail:"
        label-for="exampleInput1"
        description="Seu e-mail não será compartilhado com ninguém."
      >
        <b-form-input
          id="exampleInput1"
          type="email"
          v-model="form.email"
          required
          placeholder="email@exemplo.com" />
      </b-form-group>

      <b-form-group id="exampleInputGroup2" label="Nome:" label-for="exampleInput2">
        <b-form-input
          id="exampleInput2"
          type="text"
          v-model="form.name"
          required
          placeholder="Digite seu nome" />
      </b-form-group>

      <b-form-group id="exampleInputGroup2" label="Senha:" label-for="exampleInput3">
        <b-form-input
          id="exampleInput3"
          type="password"
          v-model="form.password"
          required
          placeholder="Digite sua senha" />
      </b-form-group>

      <!--<b-form-group id="exampleInputGroup4" label="Comida:" label-for="exampleInput4">
        <b-form-select id="exampleInput4" :options="foods" required v-model="form.food" />
      </b-form-group>-->

      <b-form-group id="exampleGroup4" label="Interesses:">
        <b-form-checkbox-group v-model="form.checked" id="exampleChecks">
          <b-form-checkbox value="me">Anunciar</b-form-checkbox>
          <b-form-checkbox value="that">Conhecer</b-form-checkbox>
        </b-form-checkbox-group>
      </b-form-group>

      <b-button class="btn btn-md btn-primary btn-block" type="submit" variant="dark">Enviar</b-button>
      <b-button class="btn btn-md btn-primary btn-block" type="reset" variant="danger">Limpar</b-button>
    </b-form>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        form: {
          email: '',
          name: '',
          food: null,
          checked: []
        },
        foods: [{ text: 'Select One', value: null }, 'Carrots', 'Beans', 'Tomatoes', 'Corn'],
        show: true
      }
    },
    methods: {
      onSubmit(evt) {
        evt.preventDefault()
        alert(JSON.stringify(this.form))
      },
      onReset(evt) {
        evt.preventDefault()
        /* Reset our form values */
        this.form.email = ''
        this.form.name = ''
        this.form.password = ''
        this.form.food = null
        this.form.checked = []
        /* Trick to reset/clear native browser form validation state */
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
    max-width: 330px;
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




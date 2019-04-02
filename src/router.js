import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/About.vue')
    },
    {
      path: '/perfil',
      name: 'perfil',
      component: () => import('./views/Perfil.vue')
    },
    {
      path: '/eventos',
      name: 'eventos',
      
      component: () => import('./views/Eventos.vue')
    },
    {
      path: '/eventos/register',
      name: 'eventosRegister',
    
      component: () => import('./views/EventosRegister.vue')
    },
    {
      path: '/login',
      name: 'login',
     
      component: () => import('./views/Login.vue')
    },
    {
      path: '/register',
      name: 'register',
     
      component: () => import('./views/Register.vue')
    },
    {
      path: '/produtos',
      name: 'produtos',
    
      component: () => import('./views/Produtos.vue')
    },
    {
      path: '/produtos/register',
      name: 'produtosRegister',
    
      component: () => import('./views/ProdutosRegister.vue')
    },
    {
      path: '/teste',
      name: 'teste',
    
      component: () => import('./views/Teste.vue')
    }
  ]
})

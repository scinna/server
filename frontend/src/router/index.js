import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Login from '../views/Login.vue'

import store from '../store';

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  }
];

const router = new VueRouter({
  routes
});

router.beforeEach((to, _, next) => {
    let token = store.state.User.Token;

    if (to.name !== 'Login' && (token.length === 0 || token === 'NO-TOKEN')) {
      next({name: 'Login'});
    } else {
      next();
    }
});

export default router;

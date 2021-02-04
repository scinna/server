import About from "@/views/About.vue";
import BrowseUser from '@/views/BrowseUser.vue';
import Home from '@/views/Home.vue'
import Login from "@/views/Login.vue";
import Logout from "@/views/Logout.vue";
import Media from "@/views/Media.vue";
import NotFound from "@/views/NotFound.vue";
import Register from "@/views/Register.vue";

import Vue from 'vue'
import VueRouter, {RouteConfig} from 'vue-router'


import store from '@/store';

Vue.use(VueRouter);

const redirectToLogin = (to, from, next) => {
    if (store.getters.isLoggedIn) {
        next();
        return;
    }

    next('/login');
}


const routes: Array<RouteConfig> = [
    {
        path: '/',
        name: 'Home',
        component: Home,
        beforeEnter: redirectToLogin,
    },
    {
        path: '/user/:username',
        name: 'Browse user',
        component: BrowseUser,
    },
    {
        path: '/login',
        name: 'Login',
        component: Login,
    },
    {
        path: '/logout',
        name: 'Logout',
        component: Logout,
    },
    {
        path: '/register',
        name: 'Register',
        component: Register,
    },
    {
        path: '/about',
        name: 'about',
        component: About,
    },
    {
        path: '/:id',
        name: 'Display media',
        component: Media,
    },
    {path: '*', component: NotFound}
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

export default router

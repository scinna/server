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

const routes: Array<RouteConfig> = [
    {
        path: '/',
        name: 'Home',
        component: Home,
        meta: { anonymous: true },
    },
    {
        path: '/browse/:username/:collection?',
        name: 'Browse user',
        component: BrowseUser,
        meta: { anonymous: true },
    },
    {
        path: '/login',
        name: 'Login',
        component: Login,
        meta: { anonymous: true },
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
        meta: { anonymous: true },
    },
    {
        path: '/about',
        name: 'about',
        component: About,
        meta: { anonymous: true },
    },
    {
        path: '/:id',
        name: 'Display media',
        component: Media,
        meta: { anonymous: true },
    },
    {path: '*', component: NotFound}
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

router.beforeEach((to, from, next) => {
    if (to.meta.anonymous || store.getters.isLoggedIn) {
        next();
        return;
    }

    next('/login');
})

export default router

import About from "@/views/About.vue";
import Home from '@/views/Home.vue'
import Login from "@/views/Login.vue";
import Media from "@/views/Media.vue";
import NotFound from "@/views/NotFound.vue";
import Register from "@/views/Register.vue";
import Vue from 'vue'
import VueRouter, {RouteConfig} from 'vue-router'

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
    {
        path: '/',
        name: 'Home',
        component: Home,
    },
    {
        path: '/login',
        name: 'Login',
        component: Login,
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

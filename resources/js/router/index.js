import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../pages/Login.vue'
import Welcome from '../pages/Welcome.vue'

Vue.use(VueRouter)

export default new VueRouter({
    mode: 'history',
    routes: [
        {path: '/', component: Welcome},
        {path: '/login', component: Login},
    ]
})
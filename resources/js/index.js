import './bootstap'
import Vue from 'vue'
import store from './store'
import router from './router'
import MainLayout from './layouts/MainLayout.vue'

new Vue({
    store,
    router,
    el: '#app',
    components: {
        MainLayout
    }
})
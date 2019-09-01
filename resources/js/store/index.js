import Vue from 'vue'
import Vuex from 'vuex'
import actions from './actions'
import mutations from './mutations'
import UserState from './modules/user'

Vue.use(Vuex)

export default new Vuex.Store({
    actions,
    mutations,
    strict: true,
    state: {},
    modules: {
        user: UserState,
    }
})
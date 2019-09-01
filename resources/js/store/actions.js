export default {
    addTodo ({ commit }, text) {
        commit('addTodo', {
        text,
        done: false
        })
    },

    removeTodo ({ commit }, todo) {
        commit('removeTodo', todo)
    },
}
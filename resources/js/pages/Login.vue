<template>
    <div class="login-container d-flex justify-content-center align-items-center">
        <form class="login-container__form">
            <div class="form-group text-center mb-4">
                <span class="display-4">Авторизация</span>
            </div>
            <div class="form-group">
                <input type="text" 
                    class="form-control form-control-lg" 
                    name="login" 
                    placeholder="Логин"
                    autocomplete="off"
                    v-model="login"
                >
            </div>
            <div class="form-group">
                <input 
                    type="password" 
                    class="form-control form-control-lg" 
                    name="password" 
                    placeholder="Пароль"
                    v-model="password"
                >
            </div>
            <div class="form-group">
                <button 
                    type="button" 
                    class="btn btn-primary btn-lg btn-block"
                    v-on:click="authorize"
                >Войти</button>
            </div>
        </form>
    </div>
</template>

<script>
    import api from '../api'
    import notifier from '../tools/notifier'

    export default {
        name: 'Login',
        data() {
            return {
                login: null,
                password: null,
            };
        },
        methods: {
            authorize: function() {
                let self = this;
                api.login(this.login, this.password).then(function(response) {

                    self.$router.push({name: 'dashboard'});
                }).catch(function(error) {                 
                    
                    let data = error.response.data;
                    notifier.error(data.content);
                });
            }
        },
        created() {

        } 
    }
</script>

<style>
   .login-container {
       width: 100%;
       min-height: 100vh;
   }
   .login-container__form {
       width: 430px;
   }
   .login-container__form .form-control-lg {
       height: 62px;
   }
   .login-container__form .btn-lg {
       height: 62px;
   }
</style>
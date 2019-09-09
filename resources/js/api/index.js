import Axios from "axios";

class Api {

    login(login, password) {

        return Axios.post('/api/v1/login/', {login: login, password: password});
    }
}

const api = new Api();

export default api
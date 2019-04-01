import { http } from "./config";

export default {

    listaTodosUsuario:() => {
        return http.get('user')
    },

    cadastraUsuario:(user) => {
        return http.post('user', user)
    }
}
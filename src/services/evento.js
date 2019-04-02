import { http } from "./config";

export default {

    listaTodosEvento:() => {
        return http.get('evento')
    },

    cadastraEvento:(idUsuario, evento) => {
        return http.post('user/' + idUsuario + '/evento', evento)
    }
}
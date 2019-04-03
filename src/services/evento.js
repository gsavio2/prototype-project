import { http } from "./config";

export default {

    listaTodosEvento:() => {
        return http.get('evento')
    },

    cadastraEvento:(idUsuario, evento) => {
        return http.post('user/' + idUsuario + '/evento', evento)
    },

    adicionaParticipante:(idEvento, idUsuario, evento) => {
        return http.post('evento/' + idEvento + '/user/' + idUsuario, evento)
    },

    removeParticipante:(idEvento, idUsuario, evento) => {
        return http.put('evento/' + idEvento + '/user/' + idUsuario, evento)
    }
}
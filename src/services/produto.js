import { http } from "./config";

export default {

    listaTodosProduto:() => {
        return http.get('produto')
    },

    cadastraProduto:(idUsuario, produto) => {
        return http.post('user/' + idUsuario + '/produto', produto)
    }
}
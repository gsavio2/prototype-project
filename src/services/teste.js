import { http } from "./config";

export default {

    getUser:() => {
        return http.get('user')
    },

    postUser:(user) => {
        return http.post('user', user)
    }


}
import { http } from "./config";

export default {
    getUser:() => {
        return http.get('user')
    }
}
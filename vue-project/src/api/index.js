import axios from "./instance";

const API_KEY = 'api/v2/';

const API_PATHS = {
    auth: 'users/auth/token/login/',
    user: 'users/employee/current/'
}

const api = {
    async auth(username, password) {
        const res = await axios.post(API_KEY + API_PATHS['auth'], {
            username,
            password
        });
        
        return res.data;
    },

    async currentUser(token) {
        const res = await axios.get(API_KEY + API_PATHS['user'], {
            headers: {
                "Authorization": "Token " + token
            }
        });

        return res.data;
    }
};

export default api;
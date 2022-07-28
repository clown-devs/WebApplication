import axios from "./instance";

const API_KEY = 'api/v2/';
const OLD_API_KEY = 'api/v1/';

const API_PATHS = {
    auth: 'users/auth/token/login/',
    user: 'users/employee/current/',
    meeting: 'meeting'
}

const api = {
    
    error: {},

    async auth(username, password) {
        
        try {
            const res = await axios.post(API_KEY + API_PATHS['auth'], {
                username,
                password
            });

            return res.data;

        } catch (error) {
            this.errorHandle(error);
            return null;
        }
        
    },

    async currentUser(token) {
        try {
            const res = await axios.get(API_KEY + API_PATHS['user'], {
                headers: {
                    "Authorization": "Token " + token
                }
            });
    
            return res.data;

        } catch(error) {
            this.errorHandle(error);
            return null;
        }
    },

    async getMeetings(past = false) {
       
        try {
            const res = await axios.get(
                OLD_API_KEY + API_PATHS['meeting'] + (past ? '?past=true' : '/')
            );
            return res.data;

        } catch (error) {
            this.errorHandle(error);
            return [];
        }

    },

    errorHandle(error) {
        
        if (error.response) {

            switch (error.response.status) {
                case 400:
                    this.error = "Неверный логин/пароль";
                    break;
                case 401:
                    this.error = 
                        error.response.statusText 
                        + "." 
                        + error.response.data.detail
                    ;
                    break;
                default:
                    this.error = error.message + "." + error.response.statusText;
            }
            
        } else if (error.request) {
            
            this.error = "Bad connection";

        } else {

            //

        }    
    }

};

export default api;
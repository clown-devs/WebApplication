import axios from 'axios'
import auth from '@/store/modules/auth';

const instance = axios.create({
    baseURL: 'http://sbermeeting.tk/',
    headers: {
        accept: 'application/json'
    }
});

export function setAuthHeaders() {
    instance.defaults.headers.common['Authorization'] = 'Token ' + auth.state.token;
};

export default instance;
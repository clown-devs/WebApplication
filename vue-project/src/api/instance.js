import axios from 'axios'

const instance = axios.create({
    baseURL: 'http://sbermeeting.tk/',
    headers: {
        accept: 'application/json'
    }
});

export default instance;
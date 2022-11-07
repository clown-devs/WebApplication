const axios = require('axios');
const url = 'http://sbermeeting.tk/';
const API_KEY = 'api/v2';
const API_METHOD = '/deploy';

console.log("deploy...");
axios.get(
    url + API_KEY + API_METHOD + '?password=123'
).then(function (response) {
    console.log(response.data);
}).catch(function (error) {
    console.log(error.code);
});
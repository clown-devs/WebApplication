import axios from "./instance";

const API_KEY = 'api/v2/';
const OLD_API_KEY = 'api/v1/';

const API_PATHS = {
    auth: 'users/auth/token/login/',
    user: 'users/employee/current/',
    meetings: 'meeting',
    client: 'client/',
    contact: 'contact/',
    place: 'place/',
    employeelist: 'employeelist/'
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

        } catch (error) {
            this.errorHandle(error);
            return null;
        }
    },

    async getMeetings(past = false) {

        try {
            const res = await axios.get(
                OLD_API_KEY + API_PATHS['meetings'] + (past ? '?past=true' : '/')
            );
            return res.data;

        } catch (error) {
            this.errorHandle(error);
            return [];
        }

    },

    async getClient(id) {
        try {
            const res = await axios.get(
                OLD_API_KEY + API_PATHS['client'] + String(id)
            );
            return res.data;

        } catch (error) {
            this.errorHandle(error);
            return {};
        }
    },

    async getClients() {
        try {
            let clients = await axios.get(
                OLD_API_KEY + API_PATHS['client']
            );
            return clients.data;

        } catch (error) {
            this.errorHandle(error);
            return [];
        }
    },

    async getContact(id) {
        try {
            let contact = await axios.get(
                OLD_API_KEY + API_PATHS['contact'] + String(id)
            );
            return contact.data;

        } catch (error) {
            this.errorHandle(error);
            return {};
        }
    },

    async getContacts() {
        try {
            let contacts = await axios.get(
                OLD_API_KEY + API_PATHS['contact']
            );
            return contacts.data;

        } catch (error) {
            this.errorHandle(error);
            return [];
        }
    },

    async getPlaces() {
        try {
            let places = await axios.get(
                OLD_API_KEY + API_PATHS['place']
            );
            return places.data;

        } catch (error) {
            this.errorHandle(error);
            return [];
        }
    },

    async getFreePlaces(start, end, date) {
        try {
            let places = await axios.get(
                OLD_API_KEY 
                + API_PATHS['place']
                + '?date=' + date
                + '&start=' + start
                + '&end=' + end
            );
            return places.data;

        } catch (error) {
            this.errorHandle(error);
            return [];
        }
    },

    async createMeeting(meeting) {
        try {
            const newMeeting = await axios.post(
                OLD_API_KEY + API_PATHS['meetings'] + '/', meeting
            );
            
            const newEmployee = await axios.post(
                OLD_API_KEY + API_PATHS['employeelist'], {
                "employee": newMeeting.data().creator,
                "meeting": newMeeting.data().id
            });
            
            return newEmployee.data();

        } catch(error) {
            this.errorHandle(error);
            return {};
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
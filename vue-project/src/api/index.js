import axios from "./instance";
import auth from '@/store/modules/auth'

const API_KEY = 'api/v2/';
const OLD_API_KEY = 'api/v1/';

const API_PATHS = {
    auth: 'users/auth/token/login/',
    user: 'users/employee/current/',
    meetings: 'meetings/',
    client: 'clients/',
    contact: 'clients/contacts/',
    place: 'meetings/place/',
    employeelist: 'employeelist/',
    users: 'users/employee/'
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

    async currentUser() {
        try {
            const res = await axios.get(API_KEY + API_PATHS['user'], {
                headers: {
                    "Authorization": "Token " + auth.state.token
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
                API_KEY + API_PATHS['meetings'] 
                + ("?employee=" + auth.state.user.id)
                + (past ? '&past=true' : '') 
            );
            return res.data;

        } catch (error) {
            this.errorHandle(error);
            return [];
        }

    },

    async getClients(isMy) {
        try {
            let clients = await axios.get(
                API_KEY + API_PATHS['client']
                + ( isMy ? "?employee=" + auth.state.user.id : "" )
            );
            return clients.data;

        } catch (error) {
            this.errorHandle(error);
            return [];
        }
    },

    async editClient(client) {
        try {
            let changedClient = await axios.patch(
                API_KEY + API_PATHS['client'] + client.id + "/", 
                client
            );
            return changedClient.data;

        } catch (error) {
            this.errorHandle(error);
            return {};
        }
    },

    async editContact(contact) {
        try {
            let changedContact = await axios.patch(
                API_KEY + API_PATHS['contact'] + contact.id + "/", 
                contact
            );
            return changedContact.data;

        } catch (error) {
            this.errorHandle(error);
            return {};
        }
    },

    async getClientContacts(clientId) {
        try {
            let contacts = await axios.get(
                API_KEY + API_PATHS['contact'] + "?client=" + clientId
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
                API_KEY + API_PATHS['place'] + '?my=true', {
                    headers: {
                        "Authorization": "Token " + auth.state.token
                    }
                }
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
                API_KEY
                + API_PATHS['place']
                + '?date=' + date
                + '&start=' + start
                + '&end=' + end
                + '&my=true'
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
                API_KEY + API_PATHS['meetings'], meeting
            );

            return newMeeting.data;

        } catch (error) {
            this.errorHandle(error);
            return {};
        }
    },

    async createClient(client) {
        try {
            let newClient = await axios.post(
                API_KEY + API_PATHS['client'], 
                client
            );
            return newClient.data;

        } catch (error) {
            this.errorHandle(error);
            return {};
        }
    },

    async createContact(contact) {
        try {
            let newContact = await axios.post(
                API_KEY + API_PATHS['contact'], 
                contact
            );
            return newContact.data;

        } catch (error) {
            this.errorHandle(error);
            return {};
        }
    },

    async isExistClient(client) {
        try {
            let res = await axios.get(
                API_KEY + API_PATHS['client'] 
                + "?inn=" + client.inn
            );
            return res.data.length ? true : false;

        } catch (error) {
            this.errorHandle(error);
            return false;
        }
    },

    async isExistContact(contact) {
        try {
            let res = await axios.get(
                API_KEY + API_PATHS['contact'] 
                + "?client=" + contact.client
                + "&first_name=" + contact.first_name 
                + "&second_name=" + contact.second_name
            );
            return res.data.length ? true : false;

        } catch (error) {
            this.errorHandle(error);
            return false;
        }
    },

    async getUsers() {
        try {
            let res = await axios.get(
                API_KEY + API_PATHS['users'] 
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
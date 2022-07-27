import api from '@/api'

export default {
    actions: {
        async logIn(context, { username, password, isSavedSession }) {

            const token = await api.auth(username, password);
            context.commit("setToken", token.auth_token);

            await context.dispatch('getUser');

            if (isSavedSession) {
                context.commit("saveSessionToLocalStorage");
            }

            context.commit("saveSessionToSessionStorage");

            return true;
        },

        async getUser(context) {
            const user = await api.currentUser(context.state.token);
            context.commit("updateUser", user);
        }
    },

    mutations: {
        setToken(state, token) {
            state.token = token;
        },

        saveSessionToLocalStorage(state) {
            localStorage.setItem("token", state.token);
            localStorage.setItem("user", JSON.stringify(state.user));
        },

        saveSessionToSessionStorage(state) {
            sessionStorage.setItem("token", state.token);
            sessionStorage.setItem("user", JSON.stringify(state.user));
        },

        updateUser(state, user) {
            state.user = user;
        },

        loadDataFromLocalStorage(state) {
            state.token = localStorage.getItem("token");
            state.user = JSON.parse(localStorage.getItem("user"));
        },

        loadDataFromSessionStorage(state) {
            state.token = sessionStorage.getItem("token");
            state.user = JSON.parse(sessionStorage.getItem("user"));
        },

        logOut(state) {
            state.token = "";
            state.user = {};
            localStorage.removeItem('token');
            localStorage.removeItem("user");
        }
    },

    state: {
        token: null,
        user: {},
        baseURL: "http://sbermeeting.tk/api/v2/",
    },

    getters: {}
}
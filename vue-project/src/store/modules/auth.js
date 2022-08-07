import api from '@/api'
import {setAuthHeaders} from '@/api/instance'

export default {
    actions: {
        async logIn(context, { username, password, isSavedSession }) {
            
            const token = await api.auth(username, password);
            if (token === null) {
                context.commit("setErrorMessage", api.error);
                api.error = '';
                return false;
            }

            context.commit("setToken", token.auth_token);
            
            let isAuth = true;

            const user = await context.dispatch('getUser');
            if (user === null) {
                isAuth = false;    
            }

            if (isSavedSession) {
                context.commit("saveSessionToLocalStorage");
            }

            context.commit("saveSessionToSessionStorage");

            if (isAuth) {
                context.commit("setErrorMessage", "");
            }
            
            return isAuth;
            
        },

        async getUser(context) {
            const user = await api.currentUser(context.state.token);
            if (user === null) {
                context.commit("setErrorMessage", api.error);
                api.error = '';
                return null;
            }

            context.commit("updateUser", user);
            return user;
        }
    },

    mutations: {
        setToken(state, token) {
            state.token = token;
            setAuthHeaders();
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
            setAuthHeaders();
        },

        loadDataFromSessionStorage(state) {
            state.token = sessionStorage.getItem("token");
            state.user = JSON.parse(sessionStorage.getItem("user"));
            setAuthHeaders();
        },

        logOut(state) {
            state.token = "";
            state.user = {};
            localStorage.removeItem('token');
            localStorage.removeItem("user");
            sessionStorage.removeItem('token');
            sessionStorage.removeItem("user");
        },

        setErrorMessage(state, error) {
            state.errorMessage = error; 
        }
    },

    state: {
        token: null,
        user: {},
        errorMessage: ''
    },

    getters: {}
}
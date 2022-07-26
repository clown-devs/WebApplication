export default {
    actions: {
        async logIn(context, {username, password, isSavedSession}) {
            const res = await fetch(
                context.state.baseURL + "users/auth/token/login/",
                {
                  method: "POST",
                  headers: {
                    "Content-Type": "application/json",
                  },
                  body: JSON.stringify({
                    username,
                    password
                  }),
                }
              );
        
              const token = await res.json();
              context.commit("setToken", token.auth_token);
              
              await context.dispatch('getUser');

              if (isSavedSession) {
                context.commit("saveTokenToLocalStorage");
                context.commit("saveUserToLocalStorage");
              }

              return true;
        },

        async getUser(context) {
            const res = await fetch(context.state.baseURL + "users/employee/current/", {
                method: "GET",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": "Token " + context.state.token
                }
              });
        
              const user = await res.json();
              context.commit("updateUser", user);
        }
    },
    
    mutations: {
        setToken(state, token) {
            state.token = token;
        },

        saveTokenToLocalStorage(state) {
            localStorage.setItem("token", state.token);
        },

        saveUserToLocalStorage(state) {
            localStorage.setItem("user", JSON.stringify(state.user));
        },

        updateUser(state, user) {
            state.user = user;
        },

        loadDataFromLocalStorage(state) {
            state.token = localStorage.getItem("token");
            state.user = JSON.parse(localStorage.getItem("user"));
        },

        logOut(state) {
            state.token = "";
            state.user = {};
            localStorage.removeItem('token');
            localStorage.removeItem("user");
        }
    },
    
    state: {
        token: "",
        user: {},
        baseURL: "http://sbermeeting.tk/api/v2/",
    },
    
    getters: {}
}
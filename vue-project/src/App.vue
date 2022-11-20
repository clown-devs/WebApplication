<template>
  <div class="app" :class="{ loginColor: isSignIn }">
    <router-view></router-view>
  </div>
</template>

<script>
import { mapMutations } from "vuex";
import api from "@/api";

export default {
  computed: {
    isSignIn() {
      return this.$route.path === "/login";
    },
  },

  methods: {
    ...mapMutations([
      "loadDataFromLocalStorage",
      "loadDataFromSessionStorage"
      ]),
  },

  mounted() {
    if(sessionStorage.getItem("token") !== null) {
      this.loadDataFromSessionStorage();
      return;
    }

    if (localStorage.getItem("token") !== null) {
      this.loadDataFromLocalStorage();
      return;
    }

    api.auth();

    this.$router.push("/login");
  },
};
</script>

<style>
body {
  margin: 0;
}

.app {
  height: 100vh;
  width: 100vw;
  background-color: #e2eee3;
}

.loginColor {
  background-color: #eaeaea;
}


::-webkit-scrollbar {
  width: 5px;
}

::-webkit-scrollbar-track {
  box-shadow: inset 0 0 1px rgba(255, 255, 255, 0);
  border-radius: 10px;
}

::-webkit-scrollbar-thumb {
  border-radius: 10px;
  background: #D3D3D3;
  box-shadow: 0 3px 2.85625px rgba(0, 0, 0, 0.25);
}

::-webkit-scrollbar-thumb:hover {
  background: #D3D3D3;
}
</style>
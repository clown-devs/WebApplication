<template>
  <div class="app" :class="{ loginColor: isSignIn }">
    <router-view></router-view>
  </div>
</template>

<script>
import { mapMutations } from "vuex";

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
</style>
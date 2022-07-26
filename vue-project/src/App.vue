<template>
  <div class="app" :class="{ loginColor: isSignIn }">
    <router-view></router-view>
    <button @click="logOutWrapper">Logout</button>
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
    ...mapMutations(["loadDataFromLocalStorage", "logOut"]),

    logOutWrapper() {
      this.$router.push("/login");
      this.logOut();
    },
  },

  mounted() {
    if (localStorage.getItem("token") === null) {
      this.$router.push("/login");
      return;
    }

    this.loadDataFromLocalStorage();
  },
};
</script>

<style>
.app {
  height: 100vh;
  width: 100vw;
  background-color: #e2eee3;
}

.loginColor {
  background-color: #eaeaea;
}
</style>
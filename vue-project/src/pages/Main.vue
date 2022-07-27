<template>
  <div class="wrapper">
    <nav-bar></nav-bar>

    <mains></mains>
  <div>
    <button @click="logOutWrapper">Logout</button>
  </div>
</template>

<script>
import { mapMutations } from "vuex";
import NavBar from "@/components/NavBar.vue";
import mains from "@/components/main.vue";

export default {
  components: { NavBar, mains },
  
  methods: {
    ...mapMutations(["logOut"]),

    logOutWrapper() {
      this.$router.push("/login");
      this.logOut();
    },
  },
};
</script>

<style>
body {
  margin: 0;
}
</style>
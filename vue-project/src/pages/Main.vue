<template>
  <div>
    <h1>Main page</h1>
    <button @click="logOutWrapper">Logout</button>
  </div>
</template>

<script>
import { mapMutations } from "vuex";
export default {
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
</style>
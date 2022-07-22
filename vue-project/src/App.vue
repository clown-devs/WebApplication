<template>
  <div class="app" :class="{ loginColor: isSignIn }">
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  computed: {
    isSignIn() {
      return this.$route.path === "/login";
    },
  },
};
</script>

<style>
.app {
  height: 100vh;
  width: 100vw;
  background-color: #E2EEE3;
}

.loginColor {
  background-color: #eaeaea;
}
</style>
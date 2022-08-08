<template>
  <div>
    <nav-bar></nav-bar>
    <loading-indicate></loading-indicate>
  </div>
</template>

<script>
import NavBar from "@/components/NavBar.vue";
import LoadingIndicate from "@/components/UI/LoadingIndicate.vue";

export default {
  components: { NavBar, LoadingIndicate },

  data() {
    return {
      clients: []
    };
  }
};
</script>

<style scoped>
</style>
<template>
  <div>
    <nav-bar></nav-bar>
    <h1>Meeting results page</h1>
    <loading-indicate></loading-indicate>
  </div>
</template>

<script>
import NavBar from "@/components/NavBar.vue";
import LoadingIndicate from "@/components/UI/LoadingIndicate.vue";

export default {
  components: { NavBar, LoadingIndicate},
};
</script>

<style>
</style>
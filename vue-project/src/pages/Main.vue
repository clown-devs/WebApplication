<template>
  <div class="wrapper">
    <nav-bar></nav-bar>

    <mains></mains>
  </div>
</template>

<script>
import NavBar from "@/components/NavBar.vue";
import mains from "@/components/main.vue";
export default {  
  components: { NavBar, mains },
};
</script>

<style>
body {
  margin: 0;
}
</style>
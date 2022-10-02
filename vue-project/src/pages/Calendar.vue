<template>
  <div class="calendar-container">
    <nav-bar></nav-bar>
    <main>
      <h1>Calendar page</h1>
      <loading-indicate></loading-indicate>
      <user-directions></user-directions>
    </main>
  </div>
</template>

<script>
import NavBar from "@/components/NavBar.vue";
import LoadingIndicate from "@/components/UI/LoadingIndicate.vue";
import UserDirections from "@/components/UserDirections.vue";

export default {
  components: {
    NavBar,
    LoadingIndicate,
    UserDirections
  },
};
</script>

<style scoped>
main {
  height: calc(100vh - 110px);
  padding-top: 110px;
}

.calendar-container {
  height: 100vh;
}
</style>
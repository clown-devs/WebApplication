<template>
  <div class="wrapper">
    <nav-bar></nav-bar>
    <meetings></meetings>
  </div>
</template>

<script>

import NavBar from "@/components/NavBar.vue";
import meetings from "@/components/Meetings.vue";

export default {
  components: { NavBar, meetings },
};
</script>

<style scoped>
body {
  margin: 0;
}

.wrapper {
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: column;
}
</style>
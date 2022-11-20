<template>
  <div class="wrapper">
    <nav-bar></nav-bar>
    <h1>Meeting results page</h1>
<!--    <loading-indicate></loading-indicate>-->
    <event-tooltip :is-show="true"></event-tooltip>
  </div>
</template>

<script>
import NavBar from "@/components/NavBar.vue";
import LoadingIndicate from "@/components/UI/LoadingIndicate.vue";
import EventTooltip from "@/components/Calendar/DaySlice/EventTooltip";

export default {
  components: {EventTooltip, NavBar, LoadingIndicate},
};
</script>

<style scoped>
.wrapper {
  height: calc(100vh - 110px);
  padding-top: 110px;
}
</style>
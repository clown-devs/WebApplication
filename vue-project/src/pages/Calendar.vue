<template>
  <div class="calendar-container">
    <nav-bar></nav-bar>
    <main>
      <calendar-container></calendar-container>
<!--      <loading-indicate></loading-indicate>-->
    </main>
  </div>
</template>

<script>
import NavBar from "@/components/NavBar.vue";
import LoadingIndicate from "@/components/UI/LoadingIndicate.vue";
import MultiCheckbox from "@/components/UI/MultiCheckbox.vue";
import CalendarContainer from "@/components/Calendar/CalendarContainer";

export default {
  components: {
    NavBar,
    LoadingIndicate,
    MultiCheckbox,
    CalendarContainer,
  },

  data() {
    return {};
  },

  async mounted() {
  },
};
</script>

<style scoped>
main {
  height: calc(100vh - 110px);
  padding-top: 80px;
  overflow-x: hidden;
  overflow-y: hidden;
}

.calendar-container {
  height: 100vh;
}
</style>
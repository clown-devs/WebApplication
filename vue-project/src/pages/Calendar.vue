<template>
  <div class="calendar-container">
    <nav-bar></nav-bar>
    <main>
      <h1>Calendar page</h1>
      <loading-indicate></loading-indicate>
      <multi-checkbox :dataOptions="test"></multi-checkbox>
    </main>
  </div>
</template>

<script>
import NavBar from "@/components/NavBar.vue";
import LoadingIndicate from "@/components/UI/LoadingIndicate.vue";
import UserDirections from "@/components/UserDirections.vue";
import MultiCheckbox from "@/components/UI/MultiCheckbox.vue";

export default {
  components: {
    NavBar,
    LoadingIndicate,
    UserDirections,
    MultiCheckbox,
  },

  data() {
    return {
      test: [
        {
          name: "option 1",
          id: 1,
          isSelected: false,
        },
        {
          name: "option 2",
          id: 2,
          isSelected: false,
        },
        {
          name: "option 3",
          id: 3,
          isSelected: false,
        },
        {
          name: "option 4",
          id: 4,
          isSelected: false,
        },
        {
          name: "option 5",
          id: 5,
          isSelected: false,
        },
        {
          name: "option 6",
          id: 6,
          isSelected: false,
        },
        {
          name: "option 6",
          id: 6,
          isSelected: false,
        },
        {
          name: "option 6",
          id: 6,
          isSelected: false,
        },
        {
          name: "option 6",
          id: 6,
          isSelected: false,
        },
        {
          name: "option 6",
          id: 6,
          isSelected: false,
        },
        {
          name: "option 6",
          id: 6,
          isSelected: false,
        },
      ],
    };
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
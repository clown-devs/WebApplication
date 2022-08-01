<template>
  <div>
    <nav-bar></nav-bar>
    <h1>Clients page</h1>
    <ul v-if="isLoadedPlaces">
      <li v-for="place in places" :key="place">{{ place }}</li>
    </ul>
    <loading-indicate v-else></loading-indicate>
  </div>
</template>

<script>
import NavBar from "@/components/NavBar.vue";
import LoadingIndicate from "@/components/UI/LoadingIndicate.vue";
import api from "@/api";

export default {
  components: { NavBar, LoadingIndicate },

  data() {
    return {
      isLoadedPlaces: false,
      places: [],
    };
  },

  async mounted() {
    const places = await api.getPlaces();
    this.places = places;
    this.isLoadedPlaces = true;
  },
};
</script>

<style scoped>
</style>
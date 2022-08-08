<template>
  <div>
    <nav-bar></nav-bar>
    <h1>Список клиентов</h1>
    <main v-if="isLoadedClientsFromApi">
      <clients :clientsInput="clients"></clients>
    </main>
    <loading-indicate v-else></loading-indicate>
  </div>
</template>

<script>
import NavBar from "@/components/NavBar.vue";
import LoadingIndicate from "@/components/UI/LoadingIndicate.vue";
import Clients from "@/components/Client.vue";
import api from "@/api";

export default {
  components: { NavBar, LoadingIndicate, Clients},

  data() {
    return {
      clients: [],
      isLoadedClientsFromApi: false
    };
  },

  async mounted() {
    this.clients = await api.getClients();
    this.isLoadedClientsFromApi = true;
  }
};
</script>

<style scoped>
h1 {
  margin: 53px 10.14% 21px 10.14%;
  font-family: "Exo 2";
  font-weight: 700;
  font-size: 1.43rem;
  transition-duration: 0.5s;
}

main {
  background-color: #e2eee3;
}

h1:hover {
  color: #00b268;
  transition-duration: 0.5s;
}
</style>
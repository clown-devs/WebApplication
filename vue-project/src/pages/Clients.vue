<template>
  <div>
    <nav-bar></nav-bar>
    <h1>Список клиентов</h1>
    <main v-if="isLoadedClientsFromApi">
      <div class="client-container">
        <ul class="clients-list" v-if="clients.length">
          <li v-for="client in clients" :key="client">
            <client :clientData=client></client>
          </li>
        </ul>
        <ul class="clients-list empty" v-else>
          <li class="clients-item clients-item-empty">
            <span class="empty-message">Нет клиентов</span>
          </li>
        </ul>
      </div>
    </main>
    <loading-indicate v-else></loading-indicate>
  </div>
</template>

<script>
import NavBar from "@/components/NavBar.vue";
import LoadingIndicate from "@/components/UI/LoadingIndicate.vue";
import Client from "@/components/Client.vue";
import api from "@/api";

export default {
  components: { NavBar, LoadingIndicate, Client },

  data() {
    return {
      clients: [],
      isLoadedClientsFromApi: false,
    };
  },

  async mounted() {
    this.clients = await api.getClients();
    this.isLoadedClientsFromApi = true;
  },
};
</script>

<style scoped>
/* Main styles */

h1 {
  margin: 53px 10.14% 21px 10.14%;
  font-family: "Exo 2";
  font-weight: 700;
  font-size: 1.43rem;
}

main {
  background-color: #e2eee3;
}

.clients-list {
  margin: 0 10.14% 0 10.14%;
  list-style: none;
  padding: 0;
}

.clients-item-empty {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 94px;
  background-color: #F1FBF2;
  border: 0.71px solid #f1fbf2;
  box-shadow: 0px 2.85625px 2.85625px rgba(0, 0, 0, 0.25);
  border-radius: 10px;
}

.empty-message {
  font-family: "Exo 2";
  font-weight: 700;
  font-size: 1rem;
}

/* Animations and hovers */

h1 {
  transition-duration: 0.5s;
}

h1:hover {
  color: #00b268;
  transition-duration: 0.5s;
}
</style>
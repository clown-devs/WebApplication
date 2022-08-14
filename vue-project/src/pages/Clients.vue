<template>
  <div class="wrapper">
    <nav-bar></nav-bar>
    <h1>Список клиентов</h1>
    <input
      type="text"
      class="search-input"
      v-model="searchQuery"
      placeholder="Поиск..."
    />
    <main v-if="isLoadedClientsFromApi">
      <div class="client-container">
        <ul class="clients-list" v-if="clients.length">
          <li v-for="client in searchedClients" :key="client">
            <client 
              :clientData="client"
              @edit="editedClient"
            ></client>
          </li>
        </ul>
        <ul class="clients-list empty" v-else>
          <li class="clients-item clients-item-empty">
            <span class="empty-message">Нет клиентов</span>
          </li>
        </ul>
      </div>
      <add-button class="add-button" @click="createClient">Добавить</add-button>
    </main>
    <loading-indicate v-else></loading-indicate>
  </div>
</template>

<script>
import NavBar from "@/components/NavBar.vue";
import LoadingIndicate from "@/components/UI/LoadingIndicate.vue";
import Client from "@/components/Client.vue";
import api from "@/api";
import AddButton from "@/components/UI/AddButton.vue";

export default {
  components: { 
    NavBar,
    LoadingIndicate,
    Client,
    AddButton 
  },

  data() {
    return {
      clients: [],
      isLoadedClientsFromApi: false,
      searchQuery: "",
    };
  },

  async mounted() {
    this.clients = await api.getClients();
    this.isLoadedClientsFromApi = true;
  },

  computed: {
    searchedClients() {
      
      return this.clients.filter(client => 
        (client.name.toLowerCase().includes(this.searchQuery.toLowerCase()) 
          || client.inn.includes(this.searchQuery))
      );

    },
  },

  methods: {
    editedClient(client) {
      
      this.clients = this.clients.map(item => {
        if (item.id === client.id) {
          return client;
        }

        return item;
      });
      
    },

    async createClient() {
      // const newClient = await api.createClient({
      //   name: "TestClient",
      //   inn: "1234567890"
      // });

      // this.clients.push(newClient);
    }
  }
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
  display:flex;
  flex-direction: column;
}
  
.wrapper {
  padding-top: 100px;
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
  background-color: #f1fbf2;
  border: 0.71px solid #f1fbf2;
  box-shadow: 0px 2.85625px 2.85625px rgba(0, 0, 0, 0.25);
  border-radius: 10px;
}

.empty-message {
  font-family: "Exo 2";
  font-weight: 700;
  font-size: 1rem;
}

.search-input {
  margin: 0 10.14% 21px 10.14%;
}

.add-button {
  align-self: center;
  margin-top: 5rem;
}

/* Animations and hovers */

h1 {
  transition-duration: 0.5s;
}

h1:hover {
  color: #00b268;
  transition-duration: 0.5s;
}

@media (max-width: 1200px) {
  .wrapper {
    padding-top: 80px;
  }
}


@media (max-width: 992px) {
  .wrapper {
    padding-top: 80px;
  }
}

@media (max-width: 768px) {
  .wrapper {
    padding-top: 50px;
  }
}
</style>
<template>
  <div class="wrapper">
    <nav-bar></nav-bar>
    <div class="list-checkbox">
      <h1>Список клиентов</h1>

      <div class="checkbox">
        <h3 class="display">Показать всех клиентов</h3>
        <input type="checkbox" name="highload1" id="highload1" />
        <label
          for="highload1"
          data-onlabel=""
          data-offlabel=""
          class="lb1"
        ></label>
      </div>
    </div>
    <form action="">
      <input
        type="text"
        class="search-input"
        v-model="searchQuery"
        placeholder=""
      />
      <img src="/svg/search.svg" alt="" class="search-icon" />
    </form>
    <main v-if="isLoadedClientsFromApi">
      <div class="client-container">
        <ul class="clients-list" v-if="clients.length">
          <li v-for="client in searchedClients" :key="client">
            <client :clientData="client" @edit="editedClient"></client>
          </li>
        </ul>
        <ul class="clients-list empty" v-else>
          <li class="clients-item clients-item-empty">
            <span class="empty-message">Нет клиентов</span>
          </li>
        </ul>
      </div>
      <add-button class="add-button" @click="createClient"
        >Добавить нового клиента</add-button
      >
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
    AddButton,
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
      return this.clients.filter(
        (client) =>
          client.name.toLowerCase().includes(this.searchQuery.toLowerCase()) ||
          client.inn.includes(this.searchQuery)
      );
    },
  },

  methods: {
    editedClient(client) {
      this.clients = this.clients.map((item) => {
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
    },
  },
};
</script>

<style scoped>
/* Main styles */

h1 {
  margin: 53px 10.14% 21px 10.14%;
  font-family: "Exo 2";
  font-weight: 700;
  font-size: 2rem;
}

form {
  position: relative;
  margin: 0 10.14% 0 10.14%;
}

main {
  background-color: #e2eee3;
  display: flex;
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
  width: 100%;
  height: 42px;
  background: #f5f5f5;
  border: 1px solid rgba(122, 116, 116, 0.64);
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  border-radius: 20px;
  padding: 0;
  margin-bottom: 18px;
  outline: none;
  padding-left: 10px;
  font-size:24px;
}

.search-input:focus {
border-color: #7A7474;
}

.search-icon {
  position: absolute; 
  top: 3px;
  right: 21px;
  border: none;
  cursor: pointer;
}

/*checkbox*/
.list-checkbox {
  display: flex;
  align-items: center;
}

.lb1 {
  margin: 2em 0 0 2em;
}

#highload1 {
  display: none;
}

#highload1 + .lb1,
#highload1 + .lb1::before,
#highload1 + .lb1::after {
  transition: all 0.3s;
}

#highload1 + .lb1 {
  display: inline-block;
  position: relative;
  width: 47px;
  height: 30px;
  border-radius: 30px;
  cursor: pointer;
}

#highload1 + .lb1::before {
  display: block;
  content: attr(data-offlabel);
  position: absolute;
  top: 18px;
  right: 10px;
  color: black;
  font-family: "Open Sans", sans-serif;
  font-size: 19px;
}

#highload1 + .lb1::after {
  border-radius: 50%;
  content: "";
  position: absolute;
  top: 3px;
  left: 3px;
  width: 25px;
  height: 25px;
  background-color: white;
}

#highload1:checked + .lb1::before {
  content: attr(data-onlabel);
  left: 16px;
  right: auto;
  color: #fff;
}

#highload1:checked + .lb1::after {
  left: 19px;
  background-color: #f7f7f7;
}

#highload1 + .lb1 {
  background-color: #ccc;
}

#highload1:checked + .lb1 {
  background-color: green;
}

#highload1:checked + .lb1::before {
  color: #fff;
}
.checkbox {
  display: flex;
  margin-left: auto;
  margin-right: 10.14%;
  align-items: flex-end;
}

.display {
  margin: 0;
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
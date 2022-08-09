<template>
  <div class="client-container">
    <ul class="clients-list" v-if="clients.length">
      <li class="clients-item" v-for="client in clients" :key="client">
        <div class="client-item-name">
          <p class="client-item-name-content">{{ client.name }}</p>
        </div>
        <div class="client-item-inn">
          <p class="client-item-inn-content">ИНН: {{ client.inn }}</p>
        </div>
        <div class="client-info">
          <a
            class="client-info-content"
            @click="touchClientInfoButton(client.id)"
            >Подробнее</a
          >
        </div>
      </li>
    </ul>
    <ul class="clients-list empty" v-else>
      <li class="clients-item clients-item-empty">
        <span class="empty-message">Нет клиентов</span>
      </li>
    </ul>
  </div>
</template>

<script>
import api from "@/api";

export default {
  data() {
    return {
      clients: this.clientsInput,
      contacts: new Map(),
    };
  },

  props: {
    clientsInput: {
      type: Array,
      default: [],
    },
  },

  methods: {
    async touchClientInfoButton(clientId) {
      const contacts = await api.getClientContacts(clientId);
      this.contacts.set(clientId, contacts);
    },
  },
};
</script>

<style scoped>
/* Main styles */

.clients-list {
  margin: 0 10.14% 0 10.14%;
  list-style: none;
  padding: 0;
}

.clients-item {
  display: grid;
  background: #ffffff;
  border: 0.71px solid #f1fbf2;
  box-shadow: 0px 2.85625px 2.85625px rgba(0, 0, 0, 0.25);
  border-radius: 10px;
  margin: 1rem 0 1rem 0;
  height: 94px;
  grid-template-rows: [start] 1fr [row2] 0.5fr [row3] 1fr [row-end];
  grid-template-columns: [start] 1fr [col2] 1fr [col-end];
  font-family: "Exo 2";
  font-weight: 500;
}

.client-item-name {
  margin-left: 2rem;
  grid-row: start / row2;
  grid-column: start / col2;
  font-size: 1.25rem;
  display: flex;
  align-items: flex-end;
  justify-content: flex-start;
  transition-duration: 1s;
}

.client-item-inn {
  margin-left: 2rem;
  grid-row: row3 / row-end;
  grid-column: start / col2;
  font-size: 1.25rem;
  transition-duration: 1s;
  display: flex;
  align-items: flex-start;
  justify-content: flex-start;
}

.client-info {
  margin-right: 2rem;
  grid-row: row2 / row3;
  grid-column: col2 / col-end;
  font-size: 1.25rem;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  font-weight: 700;
  color: #7a7474;
  transition-duration: 1s;
}

.clients-item-empty {
  display: flex;
  align-items: center;
  justify-content: center;
}

.empty-message {
  font-family: "Exo 2";
  font-weight: 700;
  font-size: 1rem;
}

/* Animations and hovers */

.clients-item {
  transition-duration: 1s;
}

.client-info-content,
.client-item-inn-content,
.client-item-name-content {
  transition-duration: 1s;
  margin: 0;
}

.client-info-content:hover,
.client-item-inn-content:hover,
.client-item-name-content:hover {
  transition-duration: 1s;
  color: #00b268;
}

.clients-item:hover {
  border: 3px solid green;
  transition-duration: 1s;
}
</style>
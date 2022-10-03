<template>
  <ul class="segments-container">
    <li class="segment-item" v-for="direction in directions" :key="direction">
      <!-- <select
        class="direction-select-item"
      > -->
      <!-- <select
        class="client-input"
        v-model="selectedClient"
        @change="selectClientHandler"
      >
        <option
          v-for="client in clients"
          :value="client"
          :disabled="client.id === 0"
        >
          {{ client.name }}
        </option>
      </select> --> 
      {{ direction.name }}
    </li>
  </ul>
</template>

<script>
import api from "@/api/index";

export default {
  data() {
    return {
      directions: [],
      users: new Map(),
    };
  },

  async mounted() {
    this.directions = await api.getDirections();
    const users = await api.getUsers();
    this.setUsers(users);
  },

  methods: {
    setUsers(users) {
      users.forEach((item) => {
        this.users.set(item.direction, new Array());
      });

      users.forEach((item) => {
        this.users.get(item.direction).push(item);
      });
    },
  },
};
</script>

<style scoped>
.segments-container {
  width: 509px;
  height: 359px;
  background: #f5f5f5;
  border-bottom: 1px solid #7a7474;
  border-radius: 10px;
}
</style>
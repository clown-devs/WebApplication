<template>
  <div class="client-history-container">
    <div class="close-container">
      <button class="close-btn" @click="closeHistory">
        <p>x</p>
      </button>
    </div>
    <div class="history-list-container" v-if="isLoadedData">
      <ul class="client-history-list" v-if="historyList.length">
        <li class="history-list-item" v-for="item in historyList" :key="item">
          <span class="history-item-number"
            >{{ this.historyList.indexOf(item) + 1 }}.</span
          >
          <span class="history-item-date">{{ item.date }}</span>
          <span class="history-item-meeting-topic">{{ item.topic }}</span>
        </li>
      </ul>
      <ul class="client-history-list empty" v-else>
        <li class="history-list-item empty-item">История пуста...</li>
      </ul>
    </div>
    <loading-indicate class="loading-indicate" v-else></loading-indicate>
  </div>
</template>

<script>
import api from "@/api";
import LoadingIndicate from "@/components/UI/LoadingIndicate.vue";

export default {
  components: { LoadingIndicate },

  data() {
    return {
      historyList: [],
      isLoadedData: false,
    };
  },

  async mounted() {
    // this.historyList = await api.getMeetings();
    this.isLoadedData = true;
  },

  methods: {
    closeHistory() {
      this.$emit("close");
    },
  },

  props: {
    client: {
      type: Object,
    },
  },
};
</script>

<style scoped>
.client-history-container {
  border: 1px solid black;
  background: #f1fbf2;
  border: 0.714062px solid #f1fbf2;
  box-shadow: 0px 2.85625px 2.85625px rgba(0, 0, 0, 0.25);
  border-radius: 10px;
  margin-top: 10px;
}

.client-history-list {
  max-height: 238px;
  font-family: "Exo 2";
  font-size: 1.25rem;
  font-weight: 500;
  color: black;
  list-style: none;
  padding: 0;
  margin-left: 3rem;
  margin-right: 3rem;
  margin-bottom: 3rem;
  overflow: auto;
}

.history-list-item {
  display: flex;
  justify-content: flex-start;
  gap: 1rem;
  margin-top: 1rem;
}

.history-item-meeting-topic {
  word-break: break-all;
}

.empty {
  display: flex;
  align-items: center;
  overflow: hidden;
}

.close-container {
  display: flex;
  justify-content: flex-end;
  background-color: inherit;
  margin-right: 0.7rem;
}

.close-btn {
  border: 0;
  cursor: pointer;
  font-family: "Exo 2";
  font-size: 1.5rem;
  background-color: inherit;
  color: #7a7474;
}

p {
  margin: 0;
}

.loading-indicate {
    margin: 0;   
}

/* Animations and hovers */
.close-btn {
  transition: 0.5s;
}

.close-btn:focus,
.close-btn:hover {
  color: #00b268;
  transition: 0.5s;
}
</style>
<template>
  <div class="client-container">
    <div class="client">
      <div class="client-name">
        <p class="client-name-content">{{ clientData.name }}</p>
      </div>

      <div class="client-inn">
        <p class="client-inn-content">ИНН: {{ clientData.inn }}</p>
      </div>

      <div class="client-info">
        <a
          class="client-info-content"
          @click="touchClientInfoButton()"
          >{{ buttonTitle }}</a
        >
      </div>
    </div>

    <ul class="contact-list" v-if="isShowContactList">
      <li class="contact-item" v-for="contact in contacts" :key="contact">
        <contact :contact="contact"></contact>
      </li>
    </ul>
  </div>
</template>

<script>
import api from "@/api";
import Contact from "@/components/Contact.vue";

export default {
  components: { Contact },

  data() {
    return {
      isShowContactList: false,
      contacts: [],
      buttonTitle: "Подробнее",
    };
  },

  props: {
    clientData: {
      type: Object,
      default: {},
    },
  },

  methods: {
    async touchClientInfoButton() {
  
      if (!this.isShowContactList) {
        this.contacts = await api.getClientContacts(this.clientData.id);
        this.buttonTitle = "Скрыть";
      } else {
        this.buttonTitle = "Подробнее";
      }

      this.isShowContactList = !this.isShowContactList;
    },
  },
};
</script>

<style scoped>
/* Main styles */

.client-container {
  margin: 1rem 0 1rem 0;
}

.client {
  display: grid;
  background-color: #f1fbf2;
  border: 0.71px solid #f1fbf2;
  box-shadow: 0px 2.85625px 2.85625px rgba(0, 0, 0, 0.25);
  border-radius: 10px;
  height: 94px;
  grid-template-rows: [start] 1fr [row2] 0.5fr [row3] 1fr [row-end];
  grid-template-columns: [start] 1fr [col2] 1fr [col-end];
  font-family: "Exo 2";
  font-weight: 500;
}

.client-name {
  margin-left: 2rem;
  grid-row: start / row2;
  grid-column: start / col2;
  font-size: 1.5rem;
  display: flex;
  align-items: flex-end;
  justify-content: flex-start;
  transition-duration: 1s;
  font-weight: 700;
}

.client-inn {
  margin-left: 2rem;
  grid-row: row3 / row-end;
  grid-column: start / col2;
  font-size: 1rem;
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

.contact-list {
  list-style: none;
  padding: 0;
}

.client-info-content {
  cursor: pointer;
}

/* Animations and hovers */

.client {
  transition-duration: 0.5s;
}

.client:hover {
  border: 3px solid green;
  transition-duration: 0.5s;
}

.client-info-content,
.client-inn-content,
.client-name-content {
  transition-duration: 0.5s;
  margin: 0;
}

.client-info-content:hover,
.client-inn-content:hover,
.client-name-content:hover {
  transition-duration: 0.5s;
  color: #00b268;
}
</style>
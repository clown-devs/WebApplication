<template>
  <div class="client-container">
    <div class="client">
      <div class="client-icon-container">
        <img class="client-icon" src="/svg/client-icon.svg" />
      </div>
      <div class="client-name">
        <p class="client-name-content">{{ client.name }}</p>
      </div>

      <div class="client-inn">
        <p class="client-inn-content">ИНН: {{ client.inn }}</p>
      </div>

      <div class="header-buttons">
        <edit-button class="edit-btn" @click="editClient"></edit-button>
        <show-info-btn
          class="show-info-btn"
          @click="touchClientInfoButton"
          :showInfo="isShowContactList"
        ></show-info-btn>
      </div>
    </div>

    <ul class="contact-list" v-if="isShowContactList">
      <li class="contact-item" v-for="contact in contacts" :key="contact">
        <contact 
          :contactData="contact"
          @edit="editedContact"
        ></contact>
      </li>
    </ul>
  </div>
</template>

<script>
import api from "@/api";
import Contact from "@/components/Contact.vue";
import ShowInfoBtn from "@/components/UI/ShowInfoButton.vue";
import EditButton from "@/components/UI/EditButton.vue";

export default {
  components: { Contact, ShowInfoBtn, EditButton },

  data() {
    return {
      isShowContactList: false,
      contacts: [],
      client: this.clientData
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
      }

      this.isShowContactList = !this.isShowContactList;
    },

    async editClient() {
      // this.client = await api.editClient({
      //   id: this.clientData.id,
      //   name: this.clientData.name,
      //   inn: "77771231423"
      // });
      
      // this.$emit("edit", this.client);
    },

    editedContact(contact) {

      this.contacts = this.contacts.map(item => {
        if (item.id === contact.id) {
          return contact;
        }

        return item;
      });

    }
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
  grid-template-columns: [start] 0.185fr [col2] 1fr [col3] 1fr [col-end];
  font-family: "Exo 2";
  font-weight: 500;
}

.client-name {
  grid-row: start / row2;
  grid-column: col2 / col3;
  font-size: 1.5rem;
  display: flex;
  align-items: flex-end;
  justify-content: flex-start;
  transition-duration: 0.5s;
  font-weight: 700;
}

.client-inn {
  grid-row: row3 / row-end;
  grid-column: col2 / col3;
  font-size: 1rem;
  transition-duration: 0.5s;
  display: flex;
  align-items: flex-start;
  justify-content: flex-start;
}

.show-info-btn {
  grid-row: row2 / row3;
  grid-column: col3 / col-end;
}

.edit-btn {
  grid-row: row2 / row3;
  grid-column: col3 / col-end;
}

.client-icon-container {
  grid-row: start / row-end;
  grid-column: start / col2;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 23px;
  margin-left: 19px;
}

.client-icon {
  height: auto;
  min-width: 40px;
}

.contact-list {
  list-style: none;
  padding: 0;
}

.header-buttons {
  margin-right: 2rem;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  gap: 0.5rem;
  grid-row: row2 / row3;
  grid-column: col3 / col-end;
}

/* Animations and hovers */

.client {
  transition-duration: 0.5s;
}

.client:hover {
  border: 3px solid green;
  transition-duration: 0.5s;
}

.client-inn-content,
.client-name-content {
  transition-duration: 0.5s;
  margin: 0;
}

.client-inn-content:hover,
.client-name-content:hover {
  transition-duration: 0.5s;
  color: #00b268;
}
</style>
<template>
  <div class="client-container">
    <div class="client"
    :id="String('client') + String(index)"
          
          >
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

        <button class="create-contact-btn" @click="touchCreateContact"></button>

        <button class="show-client-history" @click="touchShowCLientHistory">
          Встречи
        </button>

        <show-info-btn class="show-info-btn" @click="touchClientInfoButton" :showInfo="isShowContactList">
        </show-info-btn>
      </div>
    </div>

    <client-history v-if="isShowClientHistory" @close="closeHistory"></client-history>

    <ul class="direction-list">
      <li class="direction-item" v-for="(direction, index) in directions" :key="direction">
        <button class="direction-segment" :id="String('direction') + String(index)"
          @click="showUsersByDirection(direction, index)">
          {{ direction.name }}
        </button>
      </li>
    </ul>

    <!--v-bind:class="{
            'active-direction-button': isActiveDirectionButton,
          }"
          <button
          class="direction-segment"
        >
          {{ direction.name }}
        </button>
        -->


    <div class="contact-list-container" v-if="isShowContactList">
      <ul class="contact-list" v-if="allContacts.length">
        <li class="contact-item" v-for="contact in allContacts" :key="contact">
          <contact :contactData="contact" @edit="touchEditContact(contact)"></contact>
        </li>
      </ul>
      <ul contact-list v-else>
        <contact></contact>
      </ul>
    </div>
  </div>
</template>

<script>
import api from "@/api";
import Contact from "@/components/Contact.vue";
import ShowInfoBtn from "@/components/UI/ShowInfoButton.vue";
import EditButton from "@/components/UI/EditButton.vue";
import ClientHistory from "@/components/ClientHistory.vue";

export default {
  components: {
    Contact,
    ShowInfoBtn,
    EditButton,
    ClientHistory,
  },

  data() {
    return {
      isShowContactList: false,
      touchedDirectionButtonIndex: -1,
      directions: [
        {
          id: 1,
          name: "Платежи и переводы",
        },
        {
          id: 1,
          name: "Комната 2",
        },
        {
          id: 1,
          name: "Комната 3",
        },
        {
          id: 1,
          name: "Комната 4",
        },
        {
          id: 1,
          name: "Комната 5",
        },
        {
          id: 1,
          name: "Комната 6",
        },
        {
          id: 1,
          name: "Комната 7",
        },
        {
          id: 1,
          name: "Комната 8",
        },
        {
          id: 1,
          name: "Комната 9",
        },
        {
          id: 1,
          name: "Комната 10",
        },
        {
          id: 1,
          name: "Комната 11",
        },
        {
          id: 1,
          name: "Комната 12",
        },
        {
          id: 1,
          name: "К",
        },
        {
          id: 1,
          name: "Комната 14",
        },
        {
          id: 1,
          name: "Комната 15",
        },
        {
          id: 1,
          name: "Комната 6",
        },
        {
          id: 1,
          name: "Комната 7",
        },
        {
          id: 1,
          name: "Комната 8",
        },
        {
          id: 1,
          name: "Комната 9",
        },
        {
          id: 1,
          name: "Комната 10",
        },
      ],
      isShowClientHistory: false,
      contacts: [],
      client: this.clientData,
      isCreateContactMode: true,
    };
  },

  props: {
    clientData: {
      type: Object,
      default: {},
    },

    newContact: {
      type: Object,
      default: undefined,
    },
  },

  methods: {
    async touchClientInfoButton() {
      if (!this.isShowContactList) {
        this.contacts = await api.getClientContacts(this.clientData.id);
      }

      this.isShowContactList = !this.isShowContactList;
    },

    touchCreateContact() {
      this.$emit("create", this.clientData);
      this.isCreateContactMode = true;
    },

    editClient() {
      this.$emit("edit", this.clientData);
    },

    touchEditContact(contact) {
      this.$emit("editContact", contact);
      this.isCreateContactMode = false;
    },
    allContacts() {
      if (this.newContact === undefined) {
        return this.contacts;
      }

      const isNewContact = true;
      this.contacts = this.contacts.map((item, isNewContact) => {
        if (item.id === this.newContact.id) {
          isNewContact = false;
          return this.newContact;
        }

        return item;
      });

      if (this.isCreateContactMode) {
        this.contacts.push(this.newContact);
      }

      return this.contacts;
    },
    touchShowCLientHistory() {
      this.isShowClientHistory = true;
    },

    closeHistory() {
      this.isShowClientHistory = false;
    },

    showUsersByDirection(direction, index) {
      let indexPress = 'direction' + String(this.touchedDirectionButtonIndex)
      let elem = document.getElementById(indexPress)

      if (this.touchedDirectionButtonIndex !== -1) {
        elem.style.background = 'white';
        elem.style.color = '#7a7474';

      }

      this.touchedDirectionButtonIndex = index

      indexPress = 'direction' + String(index)
      elem = document.getElementById(indexPress)
      elem.style.background = '#00b268';
      elem.style.color = 'white';
    },
  },

  computed: {
    allContacts() {
      if (this.newContact === undefined) {
        return this.contacts;
      }

      const isNewContact = true;
      this.contacts = this.contacts.map((item, isNewContact) => {
        if (item.id === this.newContact.id) {
          isNewContact = false;
          return this.newContact;
        }

        return item;
      });

      if (this.isCreateContactMode) {
        this.contacts.push(this.newContact);
      }

      return this.contacts;
    },
  },
};
</script>

<style scoped>
/* Main styles */

* {
  font-family: "Exo 2";
}

.create-contact-btn {
  height: 20px;
  width: 20px;
  background-image: url("../../public/svg/AddButtonIcon.svg");
  background-color: inherit;
  border: 0;
  cursor: pointer;
  margin-left: 30px;
}

.client-container {
  margin: 1rem 0 1rem 0;
}

.direction-list {
  display: flex;
  list-style-type: none;
  padding-left: 0;
  overflow-x: scroll;
}

.direction-item {
  margin-right: 10px;
  margin-top: 22px;
  margin-bottom: 16px;
}

.direction-segment {
  background: #ffffff;
  padding: 0;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  border-radius: 15px;
  border: none;
  white-space: nowrap;
  cursor: pointer;
  height: 40px;
  padding: 8px 22px 8px 22px;
  color: #7a7474;
  font-size: 20px;
}

.client {
  display: grid;
  background-color: #fff;
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
  margin-right: 2 rem;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  gap: 0.5rem;
  grid-row: row2 / row3;
  grid-column: col3 / col-end;
}

.show-client-history {
  font-family: "Exo 2";
  font-weight: 700;
  font-size: 1.25rem;
  color: #7a7474;
  border: 0;
  background-color: inherit;
  cursor: pointer;
  margin-left: 30px;
}

.client-inn-content,
.client-name-content {
  margin: 0;
  cursor: default;
}

/* Animations and hovers */

.client,
.show-client-history {
  transition-duration: 0.5s;
}

.client:hover {
  transition-duration: 0.5s;
  background: #e5e4e4;
}

.direction-segment:hover {
  background: #00b268;
  color: white;
  transition-duration: 0.5s;
}

.active-direction-button {
  background: #00b268;
  color: white;
  transition-duration: 0.5s;
}

.direction-segment {
  transition-duration: 0.5s;
}

.create-contact-btn {
  transition-duration: 0.5s;
}

.show-client-history:hover,
.create-contact-btn:hover {
  transition-duration: 0.5s;
  color: #00b268;
}

.direction-list::-webkit-scrollbar {
  height: 6px;
  /* ширина scrollbar */
}

.direction-list::-webkit-scrollbar-thumb {
  background-color: #d3d3d3;
  /* цвет плашки */
  box-shadow: none;
}

@media (max-width: 768px) {
  .client-icon-container {
    display: none;
  }

  .client {
    display: grid;
    background-color: #f1fbf2;
    border: 0.71px solid #f1fbf2;
    box-shadow: 0px 2.85625px 2.85625px rgba(0, 0, 0, 0.25);
    border-radius: 10px;
    height: 94px;
    grid-template-rows: [start] 1.2fr [row2] 0.5fr [row3] 1fr [row-end];
    grid-template-columns: [start] 0.1fr [col2] 1fr [col3] 1fr [col-end];
    font-family: "Exo 2";
    font-weight: 500;
  }

  .client-inn-content {
    font-size: 13px;
  }

  .client-name-content {
    font-size: 19px;
  }

  .header-buttons {
    margin-right: 0.5rem;
    display: flex;
    align-items: center;
    justify-content: flex-end;
    gap: 0.5rem;
    grid-row: row2 / row3;
    grid-column: col3 / col-end;
  }

  .show-client-history {
    font-family: "Exo 2";
    font-weight: 700;
    font-size: 1.25rem;
    color: #7a7474;
    border: 0;
    background-color: inherit;
    cursor: pointer;
    margin-left: 0px;
  }
}
</style>
<template>
  <div class="client-container">
    <div class="client" @click="touchClientInfoButton">
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

        <show-info-btn
          class="show-info-btn"
          :showInfo="isShowContactList"
        ></show-info-btn>
      </div>
    </div>

    <client-history
      v-if="isShowClientHistory"
      @close="closeHistory"
    ></client-history>

    <ul class="direction-list" v-show="isShowContactList">
      <li
        v-for="direction in directions"
        :key="direction"
        class="direction-item"
      >
        <button
          class="direction-segment"
          @click="showUsersByDirection(direction)"
        >
          {{ direction.name }}
        </button>
      </li>
    </ul>

    <div class="direction-users-list-container" v-if="selectedDirection !== -1">
      <ul class="direction-users-list" v-if="directionUsers.length">
        <li
          class="direction-users-item"
          v-for="directionUser in directionUsers"
          :key="directionUser"
        >
          <direction-user :directionUserData="directionUser"></direction-user>
        </li>
        <div class="divider"></div>
      </ul>
      <ul class="direction-users-list" v-else>
        <direction-user></direction-user>
        <div class="divider"></div>
      </ul>
    </div>

    <div class="contact-list-container" v-if="isShowContactList">
      <ul class="contact-list" v-if="allContacts.length">
        <li class="contact-item" v-for="contact in allContacts" :key="contact">
          <contact
            :contactData="contact"
            @edit="touchEditContact(contact)"
          ></contact>
        </li>
      </ul>
      <ul class="contact-list" v-else>
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
import DirectionUser from "@/components/DirectionUser.vue";

export default {
  components: {
    Contact,
    ShowInfoBtn,
    EditButton,
    ClientHistory,
    DirectionUser,
  },

  data() {
    return {
      isShowContactList: false,
      directions: [],
      isShowClientHistory: false,
      contacts: [],
      client: this.clientData,
      isCreateContactMode: true,
      directionUsers: [],
      selectedDirection: -1,
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
      this.selectedDirection = -1;
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

    touchShowCLientHistory() {
      this.isShowClientHistory = true;
    },

    closeHistory() {
      this.isShowClientHistory = false;
    },

    async showUsersByDirection(direction) {
      if (!this.isShowContactList) {
        return;
      }

      if (this.selectedDirection === direction.id) {
        this.selectedDirection = -1;
        return;
      }

      this.directionUsers = await api.getClientLinkedUsersByDirection(
        this.clientData.id,
        direction.id
      );

      this.selectedDirection = direction.id;
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

  async mounted() {
    this.directions = await api.getDirections();
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

.client {
  cursor: pointer;
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

.contact-list,
.direction-users-list {
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

.divider {
  background-color: #d3d3d3;
  height: 6px;
  border-radius: 64.2656px;
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
  height: 6px; /* ширина scrollbar */
}

.direction-list::-webkit-scrollbar-thumb {
  background-color: #d3d3d3; /* цвет плашки */
  box-shadow: none;
}

@media (max-width: 768px) {
  .client-icon-container {
    display: none;
  }

  .client {
    display: grid;
    background-color: white;
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
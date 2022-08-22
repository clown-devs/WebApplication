<template>
  <div class="wrapper">
    <nav-bar></nav-bar>
    <main>
      <div class="list-checkbox">
        <h1>Список клиентов</h1>

        <div class="checkbox">
          <h3 class="display">{{ checkboxTitle }}</h3>
          <input
            type="checkbox"
            name="highload1"
            id="highload1"
            @click="touchCheckbox"
            checked
          />
          <label
            for="highload1"
            data-onlabel=""
            data-offlabel=""
            class="lb1"
          ></label>
        </div>
      </div>

      <form action="" class="search-form">
        <input
          type="text"
          class="search-input"
          v-model="searchQuery"
          placeholder=""
        />
        <img src="/svg/search.svg" alt="" class="search-icon" />
      </form>

      <div class="client-list-container" v-if="isLoadedClientsFromApi">
        <div class="clients-container">
          <ul class="clients-list" v-if="clients.length">
            <li v-for="client in searchedClients" :key="client">
              <client
                :clientData="client"
                @edit="touchEditClient"
                @create="touchCreateContact"
                @editContact="touchEditContact"
                :newContact="prepareNewContactForChildComponent(client.id)"
              ></client>
            </li>
          </ul>
          <ul class="clients-list empty" v-else>
            <li class="clients-item clients-item-empty">
              <span class="empty-message">Нет клиентов</span>
            </li>
          </ul>
        </div>
        <add-button class="add-button" @click="touchCreateClient"
          >Добавить нового клиента</add-button
        >
      </div>
      <loading-indicate v-else></loading-indicate>

      <popup
        v-if="displayClientPopup"
        @closePopup="closePopupForClient"
        class="client__modal-window"
      >
        <template v-slot:header>
          <span class="popup-title">{{ clientPopupTitle }}</span>
        </template>
        <template v-slot:content>
          <div class="company-name-container">
            <input
              type="text"
              class="company-name-input"
              placeholder="Имя клиента"
              v-model="newClientName"
              :class="{ invalid: v$.newClientName.$error }"
            />
            <small
              v-if="v$.newClientName.$error"
              class="validate-error-message"
            >
              Имя клиента не должно быть пустым!
            </small>
          </div>

          <div class="company-inn-container">
            <input
              type="text"
              class="company-inn-input"
              placeholder="ИНН"
              v-model="newClientInn"
              :class="{ invalid: v$.newClientInn.$error }"
            />
            <small
              v-if="v$.newClientInn.$dirty && v$.newClientInn.required.$invalid"
              class="validate-error-message"
            >
              ИНН является обязательным полем!
            </small>
            <small
              v-if="
                v$.newClientInn.minLength.$invalid ||
                v$.newClientInn.maxLength.$invalid
              "
              class="validate-error-message"
            >
              Длина ИНН должна быть длиной 12 символов!
            </small>
          </div>
        </template>
        <template v-slot:footer>
          <add-button
            class="popup-footer-btn"
            @click="clientPopupActionHandler"
          >
            {{ popupButtonTitle }}
          </add-button>
          <small v-if="isExistClient" class="validate-error-message">
            Клиент с таким ИНН уже существует!
          </small>
        </template>
      </popup>

      <popup
        v-if="displayContactPopup"
        @closePopup="closePopupForContact"
        class="contact__modal-window"
      >
        <template v-slot:header>
          <span class="popup-title">{{ contactPopupTitle }}</span>
        </template>

        <template v-slot:content>
          <div class="contact-name-container first-name">
            <input
              type="text"
              class="contact-name-input"
              placeholder="Имя"
              v-model="newContactFirstName"
              :class="{ invalid: v$.newContactFirstName.$error }"
            />
            <small
              v-if="v$.newContactFirstName.$error"
              class="validate-error-message"
            >
              Имя контакта не должно быть пустым!
            </small>
          </div>

          <div class="contact-name-container second-name">
            <input
              type="text"
              class="contact-name-input"
              placeholder="Фамилия"
              v-model="newContactSecondName"
              :class="{ invalid: v$.newContactSecondName.$error }"
            />
            <small
              v-if="v$.newContactSecondName.$error"
              class="validate-error-message"
            >
              Фамилия контакта не должна быть пустой!
            </small>
          </div>

          <div class="contact-name-container third-name">
            <input
              type="text"
              class="contact-name-input"
              placeholder="Отчество"
              v-model="newContactThirdName"
            />
          </div>

          <div class="contact-phone-container">
            <input
              type="text"
              class="contact-phone-input"
              placeholder="Телефон"
              v-model="newContactPhone"
            />
          </div>

          <div class="contact-position-container">
            <input
              type="text"
              class="contact-position-input"
              placeholder="Должность"
              v-model="newContactPosition"
            />
          </div>

          <div class="contact-email-container">
            <input
              type="text"
              class="contact-email-input"
              placeholder="Почта"
              v-model="newContactEmail"
            />
          </div>
        </template>

        <template v-slot:footer>
          <add-button
            class="popup-footer-btn"
            @click="contactPopupActionHandler"
          >
            {{ popupButtonTitle }}
          </add-button>
          <small v-if="isExistContact" class="validate-error-message">
            Контакт с таким именем и клиентом уже существует!
          </small>
        </template>
      </popup>
    </main>
  </div>
</template>

<script>
import NavBar from "@/components/NavBar.vue";
import LoadingIndicate from "@/components/UI/LoadingIndicate.vue";
import Client from "@/components/Client.vue";
import api from "@/api";
import AddButton from "@/components/UI/AddButton.vue";
import Popup from "@/components/UI/Popup.vue";
import auth from "@/store/modules/auth";
import { required, minLength, maxLength } from "@vuelidate/validators";
import useVuelidate from "@vuelidate/core";

export default {
  components: {
    NavBar,
    LoadingIndicate,
    Client,
    AddButton,
    Popup,
  },

  setup() {
    return {
      v$: useVuelidate(),
    };
  },

  validations() {
    return {
      newClientName: { required },
      newClientInn: {
        required,
        minLength: minLength(12),
        maxLength: maxLength(12),
      },
      newContactFirstName: { required },
      newContactSecondName: { required },
    };
  },

  data() {
    return {
      clients: [],
      isLoadedClientsFromApi: false,
      searchQuery: "",
      isMyClients: true,
      displayClientPopup: false,
      displayContactPopup: false,
      newClientName: "",
      newClientInn: "",
      newContactFirstName: "",
      newContactThirdName: "",
      newContactSecondName: "",
      newContactPhone: "",
      newContactPosition: "",
      newContactEmail: "",
      isExistClient: false,
      isExistContact: false,
      isCreateClientMode: false,
      isCreateContactMode: false,
      editClientObject: {},
      editContactObject: {},
      clientIdForNewContact: "",
      newContactForClient: new Map(),
      editClientObject: {
        name: "",
        inn: "",
      },
    };
  },

  async mounted() {
    this.clients = await api.getClients(true);
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

    checkboxTitle() {
      return this.isMyClients ? "Мои клиенты" : "Все клиенты";
    },

    employee_list() {
      let res = new Array();
      res.push(auth.state.user.id);
      return res;
    },

    popupButtonTitle() {
      return this.isCreateClientMode || this.isCreateContactMode
        ? "Создать"
        : "Сохранить";
    },

    clientPopupTitle() {
      return this.isCreateClientMode
        ? "Создание клиента"
        : "Редактирование клиента";
    },

    contactPopupTitle() {
      return this.isCreateContactMode
        ? "Создание контакта"
        : "Редактирование контакта";
    }
  },

  methods: {

    prepareNewContactForChildComponent(clientId) {
      const newContact = this.newContactForClient.get(clientId);
      this.newContactForClient.delete(clientId);
      return newContact;
    },

    async editClient() {
      const isCorrectForm = await this.isCorrectClientForm();
      if (!isCorrectForm) {
        return;
      }

      this.editClientObject.name = this.newClientName;
      this.editClientObject.inn = this.newClientInn;

      const editedClient = await api.editClient(this.editClientObject);
      this.closePopupForClient();
    },

    touchEditClient(client) {
      this.editClientObject = client;
      this.newClientName = client.name;
      this.newClientInn = client.inn;

      this.isCreateClientMode = false;
      this.showPopupForClient();
    },

    touchCreateClient() {
      this.isCreateClientMode = true;
      this.showPopupForClient();
    },

    touchCreateContact(client) {
      this.isCreateContactMode = true;
      this.showPopupForContact();
      this.clientIdForNewContact = client.id;
    },

    touchEditContact(contact) {
      this.isCreateContactMode = false;
      this.showPopupForContact();
      this.editContactObject = contact;
    
      this.newContactFirstName = contact.first_name;
      this.newContactSecondName = contact.second_name;
      this.newContactThirdName = contact.third_name;
      this.newContactPosition = contact.position;
      this.newContactPhone = contact.phone;
      this.newContactEmail = contact.email;
    },

    async createClient() {
      this.isExistClient = false;

      const isCorrectForm = await this.isCorrectClientForm();
      if (!isCorrectForm) {
        return;
      }

      const createdClient = {
        name: this.newClientName,
        inn: this.newClientInn,
        employee_list: this.employee_list,
      };

      this.isExistClient = await api.isExistClient(createdClient);

      if (this.isExistClient) {
        return;
      }

      const newClient = await api.createClient(createdClient);

      this.clients.push(newClient);
      this.closePopupForClient();
    },

    async touchCheckbox() {
      this.isLoadedClientsFromApi = false;
      this.isMyClients = !this.isMyClients;
      this.clients = await api.getClients(this.isMyClients);
      this.isLoadedClientsFromApi = true;
    },

    closePopupForClient() {
      this.displayClientPopup = false;
      this.cleareCreateClientPopup();
      this.isExistClient = false;
      this.isCreateClientMode = false;
    },

    closePopupForContact() {
      this.displayContactPopup = false;
      this.cleareCreateContactPopup();
      this.isExistContact = false;
    },

    showPopupForClient() {
      this.displayClientPopup = true;
    },

    showPopupForContact() {
      this.displayContactPopup = true;
    },

    cleareCreateClientPopup() {
      this.newClientName = "";
      this.newClientInn = "";
      this.v$.$reset();
    },

    cleareCreateContactPopup() {
      this.newContactFirstName = "";
      this.newContactSecondName = "";
      this.newContactThirdName = "";
      this.newContactPhone = "";
      this.newContactPosition = "";
      this.newContactEmail = "";
      this.v$.$reset();
    },

    async clientPopupActionHandler() {
      if (this.isCreateClientMode) {
        await this.createClient();
        return;
      }

      await this.editClient();
    },

    async contactPopupActionHandler() {
      if (this.isCreateContactMode) {
        await this.createContact();
        return;
      }

      await this.editContact();
    },

    async createContact() {
      this.isExistContact = false;

      const isCorrect = await this.isCorrectContactForm();
      if (!isCorrect) {
        return;
      }

      const createdContact = {
        client: this.clientIdForNewContact,
        first_name: this.newContactFirstName,
        second_name: this.newContactSecondName,
        third_name: this.newContactThirdName,
        phone: this.newContactPhone,
        position: this.newContactPosition,
        email: this.newContactEmail,
      };

      this.isExistContact = await api.isExistContact(createdContact);
      if (this.isExistContact) {
        return;
      }

      const newContact = await api.createContact(createdContact);
      this.newContactForClient.set(createdContact.client, newContact);

      this.closePopupForContact();
    },

    async editContact() {
      const isCorrect = await this.isCorrectContactForm();
      if (!isCorrect) {
        return;
      }

      const changedContact = await api.editContact({
        first_name: this.newContactFirstName,
        second_name: this.newContactSecondName,
        third_name: this.newContactThirdName,
        phone: this.newContactPhone,
        position: this.newContactPosition,
        email: this.newContactEmail,
        client: this.editContactObject.client,
        id: this.editContactObject.id
      });
      
      this.newContactForClient.set(
        this.editContactObject.client,
        changedContact
      );

      this.closePopupForContact();
    },

    async isCorrectContactForm() {
      const isCorrectFirstName = await this.v$.newContactFirstName.$validate();
      const isCorrectSecondName =
        await this.v$.newContactSecondName.$validate();
      if (isCorrectFirstName && isCorrectSecondName) {
        return true;
      }

      return false;
    },

    async isCorrectClientForm() {
      const isCorrectClientName = await this.v$.newClientName.$validate();
      const isCorrectClientInn = await this.v$.newClientInn.$validate();

      if (isCorrectClientName && isCorrectClientInn) {
        return true;
      }

      return false;
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
  height: 100%;
}

.client-list-container {
  background-color: #e2eee3;
  display: flex;
  flex-direction: column;
  max-height: 70%;
}

.wrapper {
  padding-top: 100px;
  height: calc(100% - 100px);
}

.clients-list {
  margin: 0 10.14% 0 10.14%;
  list-style: none;
  padding: 0;
}

.clients-container {
  overflow-y: auto;
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
  font-size: 24px;
}

.search-icon {
  position: absolute;
  top: 3px;
  right: 21px;
  border: none;
  cursor: pointer;
}

.search-form {
  margin-right: calc(10.14% + 10px);
}

.add-button {
  align-self: center;
  margin-top: 1.5rem;
  justify-content: flex-end;
  min-height: 40px;
}

/* Checkbox */

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

/* Popup styles */

.popup-title {
  flex: 2;
  display: flex;
  align-items: center;
  justify-content: center;
}

.popup-footer-btn {
  width: 80%;
  height: 40px;
}

.company-inn-input,
.company-name-input,
.contact-name-input,
.contact-phone-input,
.contact-position-input,
.contact-email-input {
  width: 100%;
  height: 30px;
  background: #f5f5f5;
  border-top: 0;
  border-left: 0;
  border-right: 0;
  border-bottom: 1px solid #7a7474;
  border-radius: 10px;
  padding-left: 13px;
  font-family: "Exo 2";
  font-weight: 700;
  font-size: 1rem;
}

.company-name-container,
.company-inn-container,
.contact-name-container,
.contact-phone-container,
.contact-position-container,
.contact-email-container {
  width: 80%;
  display: flex;
  flex-direction: column;
  align-items: center;
}

input[type="text"].invalid {
  border-bottom-color: red;
  margin-bottom: 5px;
}

.validate-error-message {
  margin-top: 5px;
  color: red;
  font-weight: bolder;
}

/* Animations and hovers */

h1 {
  transition-duration: 0.5s;
}

h1:hover {
  color: #00b268;
  transition-duration: 0.5s;
}

.search-input:focus {
  border-color: #7a7474;
}

@media (max-width: 1200px) {
  .wrapper {
    padding-top: 80px;
    margin-right: 40px;
    margin-left: 40px;
  }

  .search-form {
    margin-right: calc(10px);
  }

  .checkbox {
    margin-right: 0;
  }

  h1 {
    margin: 53px 0 21px 0;
    font-size: 1.7rem;
  }

  .clients-list {
    margin: 0;
  }

  form {
    margin: 0;
  }
}

@media (max-width: 992px) {
  .wrapper {
    padding-top: 80px;
    margin-right: 10px;
    margin-left: 10px;
  }

  .search-form {
    margin-right: calc(10px);
  }

  .checkbox {
    margin-right: 0;
  }

  h1 {
    margin: 53px 0 21px 0;
    font-size: 1.4rem;
  }

  .clients-list {
    margin: 0;
  }

  form {
    margin: 0;
  }
}

@media (max-width: 768px) {
  .wrapper {
    padding-top: 50px;
    margin-right: 10px;
    margin-left: 10px;
  }

  .search-form {
    margin-right: calc(10px);
  }

  .checkbox {
    margin-right: 0;
  }

  h1 {
    margin: 53px 0 21px 0;
    font-size: 1.3rem;
  }

  .clients-list {
    margin: 0;
  }

  form {
    margin: 0;
  }

  .lb1 {
    margin: 2em 0 0 1em;
  }

  #highload1 + .lb1 {
    display: inline-block;
    position: relative;
    width: 43px;
    height: 25px;
    border-radius: 30px;
    cursor: pointer;
  }

  #highload1 + .lb1::after {
    border-radius: 50%;
    content: "";
    position: absolute;
    top: 3px;
    left: 3px;
    width: 20px;
    height: 20px;
    background-color: white;
  }
}
</style>
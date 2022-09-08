<template>
  <div class="popup_wrapper" ref="popup_wrapper">
    <form class="v-popup animation" @submit.prevent>
      <div class="v-popup__header">
        <span class="popup-title"> {{ popupTitle }}</span>
        <span>
          <i class="material-icons" @click="closePopup"> close </i>
        </span>
      </div>

      <div class="v-popup__content">
        <div class="info-meeting">
          <div class="info-meet-main">
            <div class="client-container">
              <select
                class="client-input"
                v-model="selectedClient"
                @change="selectClientHandler"
              >
                <!-- <option value="" disabled selected>
                  Клиент
                </option> -->
                <option
                  v-for="client in clients"
                  :value="client"
                  :disabled="client.id === 0"
                >
                  {{ client.name }}
                </option>
              </select>
            </div>

            <div class="contact-container">
              <select class="contact-input" v-model="selectedContact">
                <option
                  v-for="contact in contacts"
                  :value="contact"
                  :disabled="contact.id === 0"
                >
                  {{ contact.name }}
                </option>
              </select>
            </div>

            <div class="theme-container">
              <input
                type="text"
                class="theme-input"
                placeholder="Тема"
                v-model="theme"
                :class="{ invalid: v$.theme.$error }"
              />
              <small v-if="v$.theme.$error" class="validate-error-message">
                Тема встречи является обязательным полем!
              </small>
            </div>

            <div class="date-icon-all">
              <div class="date-container">
                <input
                  type="date"
                  class="date-input"
                  placeholder="Дата"
                  v-model="date"
                  :class="{ invalid: v$.date.$error || !this.isCorrectDate }"
                  @change="changeTime"
                />
                <small v-if="v$.date.$error" class="validate-error-message">
                  Поле даты - обязательное поле!
                </small>
                <small class="validate-error-message" v-if="!isCorrectDate">
                  Дата должна быть >= текущей даты!
                </small>
              </div>

              <!-- <div class="icon-date">
                <button class="date-button">
                  <img src="/svg/date.svg" alt="" class="date-icon" />
                </button>
              </div> -->
            </div>

            <div class="time-start-end">
              <div class="time-start">
                <input
                  type="time"
                  class="date-start-input"
                  placeholder="Время начала"
                  v-model="start"
                  :class="{
                    invalid: v$.start.$error || !this.isCorrectSelectedTime,
                  }"
                  @change="changeTime"
                />
                <small v-if="v$.start.$error" class="validate-error-message">
                  Время начала встречи - обязательное поле!
                </small>
                <small
                  class="validate-error-message"
                  v-if="!isCorrectSelectedTime"
                >
                  Время начало должно быть меньше время конца встречи!
                </small>
              </div>

              <div class="time-end">
                <input
                  type="time"
                  class="date-end-input"
                  placeholder="Время окончания"
                  v-model="end"
                  :class="{ invalid: v$.end.$error }"
                  @change="changeTime"
                />
                <small v-if="v$.end.$error" class="validate-error-message">
                  Время окончания встречи - обязательное поле!
                </small>
              </div>
            </div>

            <div class="place-container">
              <select
                class="place-input"
                v-model="place"
                :class="{ invalid: v$.end.$error }"
              >
                <option
                  v-for="place in freePlaces"
                  :value="place"
                  :disabled="place.id === 0"
                >
                  {{ place.name }}
                </option>
              </select>
              <small v-if="v$.place.$error" class="validate-error-message">
                Место встречи - обязательное поле!
              </small>
            </div>
          </div>

          <div class="info-worker">
            <div class="worker-container">
              <select
                class="worker-input"
                v-model="selectedUsersInMeeting"
                multiple
              >
                <option value="" disabled selected>Сотрудники</option>
                <option v-for="user in users" v-bind:value="user">
                  {{ user.first_name }} {{ user.third_name }}
                  {{ user.second_name }}
                </option>
              </select>
            </div>
          </div>
        </div>

        <div class="info-note">
          <div class="note-container">
            <textarea
              type="text"
              class="note-input"
              placeholder="Заметки"
              v-model="note"
            ></textarea>
          </div>
        </div>
      </div>

      <div class="footer_btns">
        <add-button class="popup-footer-btn" @click="touchSaveButtonHandler"
          >Сохранить</add-button
        >
        <add-button class="popup-footer-btn" v-if="!isCreatePopup"
          >Отменить</add-button
        >
      </div>
    </form>
  </div>
</template>

<script>
import api from "@/api";
import { required, minLength, maxLength } from "@vuelidate/validators";
import useVuelidate from "@vuelidate/core";
import AddButton from "@/components/UI/AddButton.vue";
import auth from "@/store/modules/auth";

export default {
  components: { AddButton },

  props: {
    isCreatePopup: {
      type: Boolean,
      default: true,
    },

    editingMeet: {
      type: Object,
      default: null,
    },
  },

  setup() {
    return {
      v$: useVuelidate(),
    };
  },

  validations() {
    return {
      theme: { required },
      start: { required },
      end: { required },
      date: { required },
      place: { required },
    };
  },

  data() {
    return {
      selectedClient: {
        name: "Клиент",
        id: 0,
        inn: "",
      },

      selectedContact: {
        name: "Контактное лицо",
        id: 0,
      },

      selectedUsersInMeeting: [],

      clients: [
        {
          name: "Клиент",
          id: 0,
          inn: "",
        },
      ],

      contacts: [
        {
          name: "Контактное лицо",
          id: 0,
        },
      ],

      freePlaces: [
        {
          name: "Свободные переговорки",
          id: 0,
          is_private: false,
        },
      ],

      start: "",
      end: "",
      date: "",
      theme: "",

      place: {
        name: "Свободные переговорки",
        id: 0,
        is_private: false,
      },

      note: "",
      users: [],
    };
  },

  methods: {
    async touchSaveButtonHandler() {
      const isValidate = await this.isCorrectForm();
      if (!isValidate) {
        return;
      }

      let newMeeting = {
        date: this.date,
        start: this.start,
        end: this.end,
        place_name: this.place.name,
        place_id: this.place.id !== null ? this.place.id : "",
        topic: this.theme,
        note: this.note,
        creator_id: auth.state.user.id,
        contact_id: this.selectedContact.id,
        client_id: this.selectedClient.id !== 0 ? this.selectedClient.id : "",
        employee_list: this.getEmployeeList(),
      };

      newMeeting = await api.createMeeting(newMeeting);
      this.$emit("createMeeting", newMeeting);
      this.closePopup();
    },

    closePopup() {
      this.$emit("closePopup");
      this.clearePopup();
    },

    async isCorrectForm() {
      const isCorrectMeetingTheme = await this.v$.theme.$validate();
      const isCorrectStartMeeting = await this.v$.start.$validate();
      const isCorrectEndMeeting = await this.v$.end.$validate();
      const isCorrectDateMeeting = await this.v$.date.$validate();
      const isCorrectPlace = await this.v$.place.$validate();

      if (
        isCorrectMeetingTheme &&
        isCorrectStartMeeting &&
        isCorrectEndMeeting &&
        isCorrectDateMeeting &&
        isCorrectPlace &&
        this.isCorrectSelectedTime &&
        this.isCorrectDate
      ) {
        return true;
      }

      return false;
    },

    clearePopup() {
      this.selectedClient = "0";
      this.selectedContact = "";
      this.selectedPlace = "";
      this.clients = [];
      this.contacts = [];
      this.freePlaces = [];
      this.start = "";
      this.end = "";
      this.date = "";
      this.theme = "";
      this.place = "";
      this.v$.$reset();
    },

    async selectClientHandler() {
      const contacts = await api.getClientContacts(this.selectedClient.id);
      this.fillContactsForSelect(contacts);
    },

    contactPosition(contact) {
      if (contact.position === null || contact.position === "") {
        return "";
      }

      return "(" + contact.position + ")";
    },

    async changeTime() {
      if (this.start !== "" && this.end !== "" && this.date !== "") {
        this.fillFreePlaces();
      }
    },

    getEmployeeList() {
      const ls = [];
      for (let index in this.selectedUsersInMeeting) {
        ls.push(this.selectedUsersInMeeting[index].id);
      }

      return ls;
    },

    async fillDataInputs() {
      this.selectedClient = {
        name: this.editingMeet.client_name,
        id: this.editingMeet.client_id,
        inn: this.editingMeet.client_inn,
      };

      await this.selectClientHandler();

      this.selectedContact = {
        name: this.editingMeet.contact_name,
        id: this.editingMeet.contact_id,
      };

      this.theme = this.editingMeet.topic;
      this.note = this.editingMeet.note;
      this.date = this.editingMeet.date;
      this.start = this.editingMeet.start;
      this.end = this.editingMeet.end;

      this.fillFreePlaces();
      this.place = {
        name: this.editingMeet.place_name,
        id: this.editingMeet.place_id,
        is_private: false,
      };

      if (this.freePlaces.indexOf(this.place) === -1) {
        this.freePlaces.push(this.place);
      }
    },

    fillContactsForSelect(contacts) {
      this.contacts = [
        {
          name: "Контактное лицо",
          id: 0,
        },
      ];

      for (let i in contacts) {
        const fullName =
          contacts[i].second_name +
          " " +
          contacts[i].first_name +
          " " +
          contacts[i].third_name;
        this.contacts.push({
          name: fullName,
          id: contacts[i].id,
        });
      }
    },

    async fillFreePlaces() {
      const freePlaces = await api.getFreePlaces(
        this.start,
        this.end,
        this.date
      );

      this.freePlaces = [];

      this.freePlaces.push({
        name: "Свободные переговорки",
        id: 0,
        is_private: false,
      });

      this.freePlaces.push({
        name: "У клиента",
        id: null,
        is_private: false,
      });

      for (let i in freePlaces) {
        this.freePlaces.push(freePlaces[i]);
      }
    },
  },

  async mounted() {
    let vm = this;

    document.addEventListener("click", function (item) {
      if (item.target === vm.$refs["popup_wrapper"]) {
        vm.closePopup();
      }
    });

    const clients = await api.getClients(true);
    for (let index in clients) {
      this.clients.push(clients[index]);
    }

    if (!this.isCreatePopup) {
      this.fillDataInputs();
    }

    this.users = await api.getUsers();
  },

  computed: {
    isCorrectSelectedTime() {
      if (this.start === "" || this.end === "") {
        return true;
      }

      const start = new Date();
      const startHours = this.start.split(":")[0];
      const startMinutes = this.start.split(":")[1];
      start.setHours(startHours, startMinutes);

      const end = new Date();
      const endHours = this.end.split(":")[0];
      const endMinutes = this.end.split(":")[1];
      end.setHours(endHours, endMinutes);

      return start < end;
    },

    isCorrectDate() {
      if (this.date === "") {
        return true;
      }

      if (!this.isCreatePopup) {
        return true;
      }

      const selectedDate = new Date(this.date);
      const currentDate = new Date();
      currentDate.setHours(0, 0, 0, 0);

      return selectedDate >= currentDate;
    },

    popupTitle() {
      return this.isCreatePopup ? "Новая встреча" : "Редактирование встречи";
    },
  },
};
</script>

<style scoped lang="scss">
.material-icons {
  cursor: pointer;
}

.popup_wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  position: fixed;
  right: 0;
  left: 0;
  top: 0;
  background: rgba(0, 0, 0, 0.5);
  bottom: 0;
  z-index: 100;
}

.popup-title {
  flex: 2;
  display: flex;
  align-items: center;
  justify-content: center;
}

.info-meeting,
.time-start-end {
  display: flex;
}

.worker-input {
  width: 350px;
  padding-bottom: 200px;
  background: #f5f5f5;
  border-bottom: 1px solid #7a7474;
  border-radius: 10px;
}

.info-worker {
  margin-left: auto;
}

.client-input,
.contact-input,
.place-input {
  background: #f5f5f5;
  border-bottom: 1px solid #7a7474;
  border-radius: 7px;
  padding-right: 56%;
  height: 20px;
  width: 400px;
  margin-bottom: 10px;
}

.contact-container,
.client-container,
.theme-container,
.time-start,
.time-end,
.date-container,
.place-container {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.theme-input {
  background: #f5f5f5;
  border-bottom: 1px solid #7a7474;
  border-radius: 7px;
  height: 30px;
  width: 400px;
  margin-bottom: 10px;
}

.date-input {
  background: #f5f5f5;
  border-bottom: 1px solid #7a7474;
  border-radius: 7px;
  height: 20px;
  width: 400px;
  margin-bottom: 10px;
}

.date-start-input {
  background: #f5f5f5;
  border-bottom: 1px solid #7a7474;
  border-radius: 7px;
  height: 20px;
  margin-bottom: 10px;
  margin-right: 160px;
  width: 100px;
}

.note-input {
  background: #f5f5f5;
  border-bottom: 1px solid #7a7474;
  border-radius: 7px;
  padding-right: 56%;
  height: 60px;
  width: 448px;
  margin-bottom: 10px;
}

.date-end-input {
  background: #f5f5f5;
  border-bottom: 1px solid #7a7474;
  border-radius: 7px;
  height: 20px;
  margin-bottom: 10px;
  margin-left: auto;
}

.v-popup {
  border-radius: 25px;
  position: fixed;
  background: #ffffff;
  box-shadow: 0 0 17px 0 #e7e7e7;
  z-index: 10;
  width: 74%;

  &__header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin: 2rem;
    font-family: "Exo 2";
    font-weight: 700;
    font-size: 1.5rem;
  }

  &__footer {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    margin: 2rem;
  }

  &__content {
    display: flex;
    flex-direction: column;
    gap: 2rem;
    margin: 2rem;
  }

  .close_modal {
    color: #ffffff;
    background: #00b268;
  }
}

.animation {
  animation-duration: 0.55s;
  animation-fill-mode: both;
  animation-name: fadeInDown;
}

/* Validations */

.validate-error-message {
  color: red;
  font-weight: bolder;
  margin-bottom: 5px;
}

.theme-input.invalid,
.date-start-input.invalid,
.date-end-input.invalid,
.date-input.invalid,
.place-input.invalid {
  border-color: red;
  margin-bottom: 5px;
}

/* Responsive */

@keyframes fadeInDown {
  0% {
    opacity: 0;
    transform: translate3d(0px, -100%, 0px);
  }
  100% {
    opacity: 1;
    transform: none;
  }
}

@media (max-width: 1200px) {
  .v-popup {
    top: 150px;
  }
}

@media (max-width: 768px) {
  .v-popup {
    top: 0;
    width: 100vw;
    height: 100vh;
    border-radius: 0;
  }
}

/* Animations and hovers */

.material-icons {
  transition-duration: 0.5s;
}

.material-icons:hover {
  transition-duration: 0.5s;
  color: #00b268;
}
</style>
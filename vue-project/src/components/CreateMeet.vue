<template>
  <div class="popup_wrapper" ref="popup_wrapper">
    <form class="v-popup animation" @submit.prevent>
      <div class="v-popup__header">
        <slot name="header"></slot>
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
                <option value="" disabled selected>Клиент</option>
                <option v-for="client in clients" v-bind:value="client">
                  {{ client.name }}
                </option>
              </select>
            </div>

            <div class="contact-container">
              <select class="contact-input" v-model="selectedContact">
                <option value="" disabled selected>Контактное лицо</option>
                <option v-for="contact in contacts" v-bind:value="contact">
                  {{ contact.first_name }} {{ contact.third_name }}
                  {{ contact.second_name }} {{ contactPosition(contact) }}
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
                  :class="{ invalid: v$.start.$error || !this.isCorrectSelectedTime}"
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
                <option value="" disabled selected>Место встречи</option>
                <option v-for="place in freePlaces" v-bind:value="place">
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
                <option
                  v-for="user in selectedUsersInMeeting"
                  v-bind:value="user"
                >
                  {{ user.first_name }} {{ user.third_name }}
                  {{ user.second_name }}
                </option>
              </select>
            </div>
          </div>
        </div>

        <div class="info-note">
          <div class="note-container">
            <textarea type="text" class="note-input" placeholder="Заметки"></textarea>
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

export default {
  components: { AddButton },

  props: {
    isCreatePopup: {
      type: Boolean,
      default: true,
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
      selectedClient: "",
      selectedContact: "",
      selectedUsersInMeeting: "",
      clients: [],
      contacts: [],
      freePlaces: [],
      start: "",
      end: "",
      date: "",
      theme: "",
      place: "",
    };
  },

  methods: {
    async touchSaveButtonHandler() {
      if (await !this.isCorrectForm()) {
        return;
      }
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
      this.selectedClient = "";
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
      this.contacts = await api.getClientContacts(this.selectedClient.id);
    },

    contactPosition(contact) {
      if (contact.position === null || contact.position === "") {
        return "";
      }

      return "(" + contact.position + ")";
    },

    async changeTime() {
      if (this.start !== "" || this.end !== "" || this.date !== "") {
        this.freePlaces = await api.getFreePlaces(
          this.start,
          this.end,
          this.date
        );

        this.freePlaces.push({
          name: "У клиента"
        });
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

    this.clients = await api.getClients(true);
    this.selectedUsersInMeeting = await api.getUsers();
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
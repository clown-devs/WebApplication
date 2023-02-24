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


            <div class="date-container">
              <input
                  type="date"
                  class="date-input"
                  placeholder="Дата"
                  id="datepicker"
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


            <div class="time-start-end">
              <div class="time-start">
                <input
                    type="time"
                    class="date-start-input"
                    placeholder="Время начала"
                    id="start-picker"
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
                    id="end-picker"
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

          <div class="worker-container">
            <select
                class="worker-input"
                v-model="selectedUsersInMeeting"
                multiple
                style="padding: 5px 0 0 5px"
            >
              <option
                  v-for="user in users"
                  v-bind:value="user"
                  :disabled="user.id === 0"
              >
                {{ prepareUserForDisplay(user) }}
              </option>
            </select>
          </div>
        </div>

        <div class="info-tags">
          <ul class="tag-input">
            <li v-for="(tag, index) in tags"
                v-bind:value="tag"
                class="tag-list">
              <label class="checkbox">
                <input class="checkbox__input" type="checkbox" @click="tag.isSelected = !tag.isSelected" v-bind:id="'checkbox' + index">
                <svg class="checkbox__icon" viewBox="0 0 22 22">
                  <rect width="21" height="21" x=".5" y=".5" fill="#FFF" stroke="#7a7474" rx="3"/>
                  <path class="tick" stroke="#7a7474" fill="none" stroke-linecap="round" stroke-width="4"
                        d="M4 10l5 5 9-9"/>
                </svg>
                <span class="checkbox__label">{{ prepareTagForDisplay(tag) }}</span>
              </label>
            </li>
          </ul>
        </div>

        <div class="info-note">
          <div class="note-container">
            <textarea
                type="text"
                class="note-input"
                placeholder="Заметки"
                v-model="note"
                style="padding: 5px 0 0 5px"
                maxlength="4096"
            ></textarea>
          </div>
        </div>
      </div>

      <div class="footer-btns">
        <add-button
            class="popup-footer-btn popup-edit"
            v-if="!isCreatePopup"
            @click="touchCancelMeetingButtonHandler"
        >Удалить
        </add-button>
        <add-button
            class="popup-footer-btn popup-save" @click="touchSaveButtonHandler"
        >Сохранить
        </add-button
        >
      </div>
    </form>
  </div>
</template>

<script>
import api from "@/api";
import {required} from "@vuelidate/validators";
import useVuelidate from "@vuelidate/core";
import AddButton from "@/components/UI/AddButton.vue";
import auth from "@/store/modules/auth";
import Popup from "@/components/UI/Popup";
import login from "@/pages/Login";

export default {
  components: {
    Popup,
    AddButton
  },

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
      theme: {required},
      start: {required},
      end: {required},
      date: {required},
      place: {required},
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

      selectedUsersInMeeting: [
        {
          name: "Сотрудники",
          id: 0,
          direction: ""
        },
      ],

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

      users: [
        {
          name: "Сотрудники",
          id: 0,
          direction: ""
        },
      ],

      tags: []
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
        tags_list: this.getTagsList()
      };

      if (this.isCreatePopup) {
        newMeeting = await api.createMeeting(newMeeting);
        this.$emit("createMeeting", newMeeting);
      } else {
        newMeeting["id"] = this.editingMeet.id;
        const editMeeting = await api.editMeeting(newMeeting);
        this.$emit("editMeeting", editMeeting);
      }

      this.closePopup();
    },

    closePopup() {
      this.$emit("closePopup");
      this.clearPopup();
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

    clearPopup() {
      this.selectedClient = "0";
      this.selectedContact = "";
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
      if (this.selectedClient.id !== null) {
        const contacts = await api.getClientContacts(this.selectedClient.id);
        this.fillContactsForSelect(contacts);
      }
    },

    async changeTime() {
      if (this.start !== "" && this.end !== "" && this.date !== "") {
        await this.fillFreePlaces();
      }
    },

    getEmployeeList() {
      const ls = [];
      for (let index in this.selectedUsersInMeeting) {
        if (this.selectedUsersInMeeting[index].id === 0) {
          continue;
        }

        ls.push(this.selectedUsersInMeeting[index].id);
      }

      return ls;
    },

    getTagsList() {
      const ls = [];
      this.tags.forEach(tag => {
        if (tag.isSelected) ls.push(tag.id);
      });
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

      await this.fillFreePlaces();
      this.place = {
        name: this.editingMeet.place_name,
        id: this.editingMeet.place_id,
        is_private: false,
      };

      if (this.freePlaces.indexOf(this.place) === -1) {
        this.freePlaces.push(this.place);
      }

      await this.fillUsers();
      await this.fillTags();
      await this.fillSelectedUsers();
      await this.fillSelectedTags();


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

      let isAccessOldSelectedPlace = false;
      for (let i in freePlaces) {
        if (freePlaces[i].id === this.place.id) {
          isAccessOldSelectedPlace = true;
        }
        this.freePlaces.push(freePlaces[i]);
      }

      if (!isAccessOldSelectedPlace) {
        this.place = {};
      }
    },

    async fillUsers() {
      const users = await api.getUsers();
      this.users = [];
      for (let i in this.selectedUsersInMeeting) {
        this.users.push(this.selectedUsersInMeeting[i]);
      }
      for (let i in users) {
        const fullName =
            users[i].second_name +
            " " +
            users[i].first_name[0] + '.' +
            users[i].third_name[0];

        this.users.push({
          name: fullName,
          id: users[i].id,
          direction: users[i].direction_name === null ? "" : users[i].direction_name
        });
      }
    },

    async fillTags() {
      const tags = await api.getTags();
      this.tags = [];

      for (let i in tags) {
        this.tags.push({
          id: tags[i].id,
          name: tags[i].name,
          isSelected: false
        });
      }
    },

    async fillSelectedUsers() {
      const selectedEmployers = new Map();
      for (let i in this.editingMeet.employee_list) {
        selectedEmployers.set(this.editingMeet.employee_list[i], null);
      }

      const allUsers = this.users;
      for (let i in allUsers) {
        const fullName =
            allUsers[i].second_name +
            " " +
            allUsers[i].first_name +
            " " +
            allUsers[i].third_name;
        if (selectedEmployers.has(allUsers[i].id)) {
          selectedEmployers.set(allUsers[i].id, fullName);
        }
      }

      this.selectedUsersInMeeting = [];
      for (let i in this.users) {
        if (selectedEmployers.get(this.users[i].id) !== undefined) {
          this.selectedUsersInMeeting.push(this.users[i]);
        }
      }
    },

    async fillSelectedTags() {
      for (let i in this.editingMeet.tags_list) {
        for (let j = 0; j < this.tags.length; ++j) {
          if (this.editingMeet.tags_list[i] === this.tags[j].id) {
            this.tags[j].isSelected = true;
            let qwe = document.getElementById('checkbox' + j)
            qwe.checked = true
          }
        }
      }
    },

    async touchCancelMeetingButtonHandler() {
      await api.deleteMeeting(this.editingMeet.id);
      this.$emit("cancelMeeting", this.editingMeet.id);
      this.closePopup();
    },

    prepareUserForDisplay(user) {
      return user.name + ((user.id !== 0) ? "(" + user.direction + ")" : "");
    },

    prepareTagForDisplay(tag) {
      return tag.name;
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
      await this.fillDataInputs();
    } else {
      await this.fillUsers();
      await this.fillTags();
    }
  },

  computed: {
    isCorrectSelectedTime() {
      if (this.start === "" || this.end === "") {
        return true;
      }

      const start = new Date();
      const startHours = this.start.split(":")[0];
      const startMinutes = this.start.split(":")[1];
      start.setHours(Number(startHours), Number(startMinutes));

      const end = new Date();
      const endHours = this.end.split(":")[0];
      const endMinutes = this.end.split(":")[1];
      end.setHours(Number(endHours), Number(endMinutes));

      return start < end;
    },

    selectedNames() {
      return this.tags
          .filter(tag => tag.isSelected)
          .map(tag => tag.name);
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
  overflow: auto;
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

.info-meeting {
  width: 100%;
}

.info-meet-main {
  margin-right: 10px;
  flex: 1.5;
}

.tag-input {
  display: flex;
  list-style-type: none;
  border: 0;
  border-bottom: 1px solid #7a7474;
  border-radius: 10px;
  background: #f5f5f5;
  height: 50px;
  margin-bottom: 0;
  padding-left: 15px;
}

.worker-input {
  background: #f5f5f5;
  width: 100%;
  border: 0;
  border-bottom: 1px solid #7a7474;
  border-radius: 10px;
}

.worker-container {
  display: flex;
  width: 350px;
  justify-content: flex-end;
}

.client-input {
  background: #f5f5f5;
  border: 0;
  border-bottom: 1px solid #7a7474;
  border-radius: 7px;
  height: 30px;
  width: 100%;
  margin-bottom: 10px;
}

.contact-input {
  background: #f5f5f5;
  border: 0;
  border-bottom: 1px solid #7a7474;
  border-radius: 7px;
  height: 30px;
  width: 100%;
  margin-bottom: 10px;
}

.place-input {
  background: #f5f5f5;
  border: 0;
  border-bottom: 1px solid #7a7474;
  border-radius: 7px;
  height: 50px;
  width: 100%;
}


.time-start {
  margin-right: 20px;
  flex: 1 0 auto;
}

.time-end {
  flex: 1 0 auto;
}

.client-container {
  width: 100%;
}

.popup-save,
.popup-edit {
  width: 400px;
  margin-top: 30px;
  margin-bottom: 30px;
}

.popup-edit {
  background-color: #CD5C5C;
}

.theme-input {
  background: #f5f5f5;
  border: 0;
  border-bottom: 1px solid #7a7474;
  border-radius: 7px;
  height: 45px;
  width: 99.5%;
  margin-bottom: 10px;
  padding: 0 0 0 5px;
}

.footer-btns {
  display: flex;
  justify-content: space-around;
}

.date-input {
  background: #f5f5f5;
  border: 0;
  border-bottom: 1px solid #7a7474;
  border-radius: 7px;
  height: 30px;
  width: 99.3%;
  margin-bottom: 10px;
  padding: 0 0 0 5px
}

#datepicker:before {
  content: "Дата: ";
  margin-right: 0.6em;
  padding-left: 5px;
  color: #9d9d9d;
}


#start-picker:before {
  content: "Время начала: ";
  margin-right: 0.6em;
  padding-left: 5px;
  color: #9d9d9d;
}

#end-picker:before {
  content: "Время окончания: ";
  margin-right: 0.6em;
  padding-left: 5px;
  color: #9d9d9d;
}

.date-start-input {
  background: #f5f5f5;
  border: 0;
  border-bottom: 1px solid #7a7474;
  border-radius: 7px;
  height: 30px;
  margin-bottom: 10px;
  margin-right: 30px;
  width: 100%;
}

.note-input {
  background: #f5f5f5;
  border: 0;
  border-bottom: 1px solid #7a7474;
  border-radius: 7px;
  height: 100px;
  width: 100%;
  margin-top: 20px;
  resize: none;
}

.date-end-input {
  background: #f5f5f5;
  border: 0;
  border-bottom: 1px solid #7a7474;
  border-radius: 7px;
  height: 30px;
  width: 100%;
  margin-bottom: 10px;
}

.v-popup {
  border-radius: 25px;
  position: fixed;
  background: #ffffff;
  box-shadow: 0 0 17px 0 #e7e7e7;
  z-index: 10;
  width: 74%;
  overflow: auto;

  &__header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin: 2rem;
    font-family: "Exo 2", serif;
    font-weight: 700;
    font-size: 1.5rem;
  }

  &__footer {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    margin: 4rem;
  }

  &__content {
    flex-direction: column;
    gap: 2rem;
    margin: 2rem 3rem 0 3rem;
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
  .theme-input {
    width: 99%;
  }

  .date-input {
    width: 99%;
  }

  .v-popup {
    top: 150px;
  }

  .time-start-end {
    display: block;
  }

  .time-start {
    margin-right: 0;
  }

  .date-start-input {
    margin-right: 0;
  }

  .popup-save,
  .popup-edit {
    width: 300px;
  }

  .worker-container {
  }
}

@media (max-width: 992px) {
  .info-meeting {
    display: block;
  }

  .info-meet-main {
    margin-right: 0;
  }

  .time-start-end {
    display: flex;
    justify-content: space-between;
  }

  .time-start {
    width: 200px;
    margin-right: 20px;
  }

  .time-end {
    width: 200px;
  }

  .place-input {
    margin-bottom: 10px;
  }

  .worker-container {
    width: 100%;
  }

  .popup-save,
  .popup-edit {
    width: 250px;
  }

}

@media (max-width: 768px) {
  .v-popup {
    top: 0;
    width: 100vw;
    height: 100vh;
    border-radius: 0;
  }

  .time-start-end {
    display: block;
  }

  .time-start {
    margin-right: 0;
    width: 100%;
  }

  .time-end {
    width: 100%;
  }

  .popup-save,
  .popup-edit {
    width: 180px;
  }
  .client-input {
    height: 40px;
  }

  .contact-input {
    height: 40px;
  }

  .theme-input {
    height: 65px;
    width: 99%;
  }

  .date-input {
    height: 40px;
  }

  .date-start-input {
    height: 40px;
    margin-bottom: 10px;
    margin-right: 0;
  }
  .date-end-input {
    height: 40px;
    margin-bottom: 10px;
  }

  .place-input {
    height: 40px;
  }

  .note-input {
    height: 120px;
    margin-top: 10px;
  }

  .worker-input {
    height: 40px;
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


/* Я ЭТОТ КАЛ ПОФИКШУ ЧИСТО НА ВРЕМЯ */

.checkbox__input:checked + .checkbox__icon .tick {
  /* убираем смещение для отрезков, чтобы включить анимацию галочки */
  stroke-dashoffset: 0;
}

.checkbox {
  /* меняем внешний вид курсора */
  cursor: pointer;
  /* выравниваем элементы по центру */
  display: flex;
  align-items: center;
}

/* отдельные настройки для самого чекбокса */
.checkbox__input {
  /* устанавливаем абсолютное позиционирование */
  position: absolute;
  /* задаём высоту и ширину */
  width: 15px;
  height: 15px;
  /* делаем чекбокс непрозрачным, чтобы скрыть исходный элемент и заменить его потом нарисованным */
  opacity: 0;
  /* меняем внешний вид курсора */
  cursor: pointer;
  margin: 0;
}

/* настройки для SVG-иконки */
.checkbox__icon {
  /* размеры совпадают с размерами скрытого чекбокса */
  width: 15px;
  height: 15px;
  /* убираем ограничение по наименьшей ширине блока */
  flex-shrink: 0;
  /* разрешаем отображать содержимое за пределами блока */
  overflow: visible;
}

/* общие настройки для нового чекбокса и галочки */
.checkbox__icon .tick {
  /* рисовать будем всё отрезками по 20 пикселей */
  stroke-dasharray: 20px;
  /* но сместим следующие отрезки тоже на 20 пикселей, чтобы получить сплошные линии */
  stroke-dashoffset: 20px;
  /* это даст нам плавную анимацию отрисовки галочки */
  transition: stroke-dashoffset 0.2s ease-out;
}

/* настройки для подписи чекбокса */
.checkbox__label {
  /* добавляем отступ слева */
  margin-left: 5px;
  font-size: 13px;
}

.tag-list:not(:last-child) {
  margin-right: 15px;
}

.tag-list {
  align-self: center;
}
</style>
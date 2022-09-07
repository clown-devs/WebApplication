<template>
  <main class="main-container">
    <div v-if="isLoadedMeetings" class="near-meet">
      <h2 class="near-meet-text">Ближайшая встреча:</h2>

      <div class="container-window" v-if="nearMeeting !== undefined">
        <button class="pencil-meet">
          <img src="/svg/pencil.svg" alt="" class="pencil-icon" />
        </button>
        <p class="client-name">{{ nearMeeting.client_name }}</p>
        <div class="window-tel-and-contact">
          <p class="contact-window">
            {{ nearMeeting.contact_name }}
          </p>
          <p class="tel-window">{{ nearMeeting.contact_phone }}</p>
        </div>
        <p class="theme-window">{{ nearMeeting.topic }}</p>
        <div class="data-and-time-window">
          <p class="data-window">{{ nearMeeting.date }}</p>
          <p class="time-window">
            {{ nearMeeting.start }} - {{ nearMeeting.end }}
          </p>
        </div>
        <p class="place-window">{{ nearMeeting.place_name }}</p>
      </div>

      <div class="container-window empty" v-else>
        <p>Нет ближайшей встречи</p>
      </div>

      <div class="green-button-add">
        <add-button @click="touchCreateMeet">Добавить встречу</add-button>
      </div>
    </div>

    <create-meet
      v-if="displayMeetPopup"
      @closePopup="closePopupForMeet"
      class="meet__modal-window"
    >
      <template v-slot:header>
        <span class="popup-title"> {{ meetPopupTitle }}</span>
      </template>

      
      <template v-slot:footer>
        <add-button class="popup-footer-btn" @click="clientPopupActionHandler">
          Сохранить
        </add-button>
      </template>
    </create-meet>

    <div v-if="isLoadedMeetings" class="list-meet">
      <h2 class="list-meet-text">Список встреч:</h2>

      <segmented-control
        class="segmented-control"
        @selected="selectedSegmentedControl"
      >
        <template v-slot:item1>
          <span>Текущие</span>
        </template>
        <template v-slot:item2>
          <span>Прошедшие</span>
        </template>
      </segmented-control>

      <ul class="list-all-meet" v-if="meetings.length">
        <li v-for="meeting in meetings" :key="meeting" class="all-meet-item">
          <div class="data-time-place-item">
            <p class="data-item">{{ meeting.date }}</p>
            <p class="time-item">{{ meeting.start }} - {{ meeting.end }}</p>
            <p class="place-item">{{ meeting.place_name }}</p>
            <edit-button class="edit-btn"></edit-button>
          </div>
          <p class="theme-item">{{ meeting.topic }}</p>
          <p class="client-item">{{ meeting.client_name }}</p>
        </li>
      </ul>
      <ul class="list-all-meet" v-else>
        <li class="all-meet-item empty">
          <p>Нет встреч</p>
        </li>
      </ul>
    </div>

    <loading-indicate v-else class="indicator"></loading-indicate>
  </main>
</template>

<script>
import AddButton from "@/components/UI/AddButton.vue";
import api from "@/api";
import LoadingIndicate from "@/components/UI/LoadingIndicate.vue";
import SegmentedControl from "@/components/UI/SegmentedControl.vue";
import EditButton from "@/components/UI/EditButton.vue";
import CreateMeet from "@/components/CreateMeet.vue";

export default {
  components: {
    AddButton,
    LoadingIndicate,
    SegmentedControl,
    EditButton,
    CreateMeet,
  },

  data() {
    return {
      meetings: [],
      nearMeeting: undefined,
      isLoadedMeetings: false,
      isCreateMeetMode: false,
      displayMeetPopup: false,
    };
  },

  async mounted() {
    this.meetings = await api.getMeetings();
    if (this.meetings.length) {
      this.nearMeeting = this.meetings[0];
      this.prepareMeetingsForDisplay();
    }

    // this.meetings.push(
    //     {
    //       id: 1,
    //       date: '12.05.2022',
    //       start: '12:30',
    //       end: '13:00',
    //       place_name: 'Peregovorka 1',
    //       topic: 'Hahahahaha'
    //     }
    // );

    this.isLoadedMeetings = true;
  },

  computed: {
    meetPopupTitle() {
      return this.isCreateMeetMode ? "Новая встреча" : "Редактирование встречи";
    },
  },

  methods: {
    async prepareMeetingsForDisplay() {
      this.meetings.forEach((meeting) => {
        meeting.start = meeting.start.substr(0, 5);
        meeting.end = meeting.end.substr(0, 5);
      });
    },

    async selectedSegmentedControl(selectedFirstControl) {
      if (selectedFirstControl) {
        this.meetings = await api.getMeetings();
        return;
      }

      this.meetings = await api.getMeetings(true);
    },

    touchCreateMeet() {
      this.isCreateMeetMode = true;
      this.showPopupForMeet();
    },

    showPopupForMeet() {
      this.displayMeetPopup = true;
    },

    closePopupForMeet() {
      this.displayMeetPopup = false;
      this.isCreateMeetMode = false;
    },
  },
};
</script>

<style scoped>
@import url("../assests/fonts/Exo_2/stylesheet.css");

* {
  font-family: "Exo 2";
}

ul {
  list-style: none;
  padding: 0;
}

main {
  display: flex;
  justify-content: space-around;
  width: 100vw;
  height: calc(100vh - 110px);
  margin-top: 1rem;
}
.empty {
  display: flex;
  align-items: center;
  justify-content: center;
}

button {
  background-color: white;
  border: none;
  margin: 0;
  padding: 0;
}

.popup-title {
  flex: 2;
  display: flex;
  align-items: center;
  justify-content: center;
}

p {
  font-size: 20px;
  margin: 0;
}

li p {
  font-family: "Exo 2";
  font-weight: 700;
  font-size: 0.75rem;
  margin-left: 17px;
  margin-right: 17px;
}

.main-container {
  padding-top: 110px;
}
.theme-item {
  font-size: 0.88rem;
  overflow: auto;
  word-break: break-all;
}

.near-meet-text {
  margin-top: 32px;
  margin-bottom: 40px;
  font-size: 2rem;
  font-family: "Exo 2";
  font-weight: 700;
}

.list-meet-text {
  margin-top: 32px;
  margin-bottom: 40px;
  font-size: 2rem;
  font-family: "Exo 2";
  font-weight: 700;
}

.near-meet {
  margin-right: 40px;
  margin-left: 10.14%;
  flex: 2.6;
}

.green-button-add {
  display: flex;
  justify-content: center;
}

.theme-item {
  white-space: normal;
}
.data-time-place-item {
  display: flex;
  justify-content: space-between;
  white-space: normal;
}

.window-tel-and-contact {
  display: flex;
}

.pencil-meet {
  margin-top: 38px;
  margin-right: 23px;
  margin-left: auto;
  display: flex;
  background: #ffffff;
  border: 0;
  padding: 0;
}

.contact-window {
  flex: 1;
}

.tel-window {
  flex: 1;
}

.all-meet-item {
  height: 127px;
  background: #fff;
  border: 0.714062px solid #f1fbf2;
  box-shadow: 0px 2.85625px 2.85625px rgba(0, 0, 0, 0.25);
  border-radius: 10px;
  margin-bottom: 9px;
  margin-top: 14px;
}

.data-and-time-window {
  display: flex;
}

.near-meet {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.container-window {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  background: #ffffff;
  border: 1px solid #bdbdbd;
  border-radius: 30px;
  margin-bottom: 2rem;
  justify-content: space-around;
  height: 63.5%;
}

.client-name,
.contact-window,
.theme-window,
.data-window,
.place-window {
  margin-left: 40px;
}

.place-window {
  margin-bottom: 60px;
}

.theme-window {
  margin-bottom: 123px;
  font-weight: 700;
  font-size: 2rem;
  word-break: break-all;
  overflow: auto;
}

.list-all-meet {
  margin-top: 0;
  margin-bottom: 0;
  height: 72%;
  background: #00b268;
  border-radius: 0px 0px 20px 20px;
  overflow: auto;
  min-width: 406px;
}

.list-all-meet li {
  font-weight: 700;
  margin-left: 16px;
  margin-right: 16px;
}

.data-window {
  flex: 1;
  font-size: 1.125rem;
  font-weight: 700;
}

.time-window {
  flex: 1;
  font-weight: 700;
  font-size: 1.25rem;
}

.client-name {
  margin-bottom: 26px;
}

.list-meet {
  flex: 1;
  margin-right: 10.14%;
  display: flex;
  flex-direction: column;
  height: 100%;
}

.data-window,
.time-window,
.place-window {
  font-weight: bolder;
}

.pencil-icon {
  max-width: 100%;
}

.segmented-control {
  height: 40px;
}

.edit-btn {
  background-color: #ffffff;
  margin-right: 10px;
}

/* Responsive layout */

@media (max-width: 1200px) {
  main {
    flex-direction: column;
    align-content: center;
    background-color: #e2eee3;
    height: auto;
    gap: 1rem;
    margin-top: 2rem;
  }

  .main-container {
    padding-top: 90px;
  }

  .nav-bar-add-meet-main {
    display: block;
  }

  .list-meet {
    margin-left: 2rem;
    margin-right: 40px;
    margin-left: 40px;
    height: auto;
  }

  .list-all-meet {
    overflow: hidden;
    height: auto;
    min-width: 100%;
  }

  .near-meet {
    margin-left: 40px;
    flex: 1;
  }

  .container-window {
    height: 300px;
    width: 100%;
  }

  .green-button-add {
    display: none;
  }

  .theme-window,
  .client-name {
    margin-bottom: 0;
  }

  .pencil-meet {
    margin: 1rem 1rem 1rem 1rem;
    justify-content: flex-end;
  }

  .pencil-icon {
    max-width: 30px;
    max-height: 30px;
  }

  .theme-window {
    margin-right: 1rem;
  }

  .near-meet-text,
  .list-meet-text {
    margin-top: 1rem;
    margin-bottom: 1rem;
  }
}

@media (max-width: 992px) {
  .container-window {
    width: 100%;
    margin-bottom: 0;
    height: 250px;
  }

  .main-container {
    padding-top: 90px;
    margin-top: 1rem;
  }

  .near-meet {
    margin-left: 10px;
    margin-right: 10px;
  }

  .list-meet {
    margin-left: 10px;
    margin-right: 10px;
  }

  .near-meet-text,
  .pencil-meet,
  .list-meet-text {
    display: none;
  }

  .place-window {
    margin-bottom: 0;
    font-size: 1.25rem;
    font-weight: 700;
  }
}

@media (max-width: 768px) {
  .container-window {
    height: 200px;
    gap: 0;
    justify-content: space-around;
    margin-top: 0;
  }

  main p {
    font-size: 1rem;
  }

  .main-container {
    padding-top: 70px;
    margin-top: 1rem;
  }

  .data-window,
  .time-window,
  .place-window {
    font-size: 1rem;
  }

  .theme-window {
    font-size: 1.1rem;
  }

  .list-all-meet li p,
  .data-time-place-item p {
    font-size: 0.7rem;
  }
}

@media (max-height: 851px) {
  .container-window {
    justify-content: center;
    gap: 1rem;
  }

  .client-name,
  .contact-window,
  .theme-window,
  .data-window,
  .place-window {
    margin-bottom: 0;
    margin-top: 0;
  }

  .pencil-meet {
    display: none;
  }
}

@media (max-height: 415px) {
  .container-window {
    height: 180px;
    justify-content: center;
    gap: 0.5rem;
  }

  .client-name,
  .contact-window,
  .tel-window,
  .theme-window,
  .data-window,
  .time-window,
  .place-window {
    font-size: 1.1rem;
  }

  .near-meet-text,
  .pencil-meet,
  .list-meet-text {
    display: none;
  }

  .place-window {
    margin-bottom: 0;
  }
}

/* Hovers and animations */

.contact-window {
  transition-duration: 0.5s;
}

.container-window:hover {
  transform: translateY(-3px);
  transition: 0.5s;
}

.near-meet-text,
.client-name,
.contact-window,
.tel-window,
.theme-window,
.data-window,
.time-window,
.place-window,
.data-item,
.time-item,
.place-item,
.theme-item,
.client-item,
.list-meet-text {
  transition: 0.3s;
}

.near-meet-text:hover,
.client-name:hover,
.contact-window:hover,
.tel-window:hover,
.theme-window:hover,
.data-window:hover,
.time-window:hover,
.place-window:hover,
.data-item:hover,
.time-item:hover,
.place-item:hover,
.theme-item:hover,
.client-item:hover,
.list-meet-text:hover {
  color: #00b268;
  transition: 0.3s;
}

.all-meet-item {
  transition: 0.5s;
}

.all-meet-item:hover {
  transform: translateY(-3px);
  transition: 0.5s;
}
</style>
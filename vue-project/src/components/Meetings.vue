<template>
  <main>
    <div v-if="isLoadedMeetings" class="near-meet">
      <h4 class="near-meet-text">Ближайшая встреча:</h4>

      <div class="container-window">
        <button class="pencil-meet">
          <img src="/svg/pencil.svg" alt="" class="pencil-icon" />
        </button>
        <p class="client-wiindow">{{ nearMeeting.client.name }}</p>
        <div class="window-tel-and-contact">
          <p class="contact-window">
            {{ nearMeeting.contact.first_name }}
            {{ nearMeeting.contact.second_name }}
            {{ nearMeeting.contact.third_name }}
          </p>
          <p class="tel-window">{{ nearMeeting.contact.phone }}</p>
        </div>
        <p class="theme-window">{{ nearMeeting.topic }}</p>
        <div class="data-and-time-window">
          <p class="data-window">{{ nearMeeting.date }}</p>
          <p class="time-window">
            {{ nearMeeting.start }} - {{ nearMeeting.end }}
          </p>
        </div>
        <p class="place-window">{{ nearMeeting.place.name }}</p>
      </div>

      <div class="green-button-add">
        <buttons></buttons>
      </div>
    </div>

    <div v-if="isLoadedMeetings" class="list-meet">
      <h4 class="list-meet-text">Список встреч</h4>

      <ul class="list-all-meet">
        <li v-for="meeting in meetings" :key="meeting" class="all-meet-item">
          <div class="data-time-place-item">
            <p class="data-item">{{ meeting.date }}</p>
            <p class="time-item">{{ meeting.start }} - {{ meeting.end }}</p>
            <p class="place-item">{{ meeting.place.name }}</p>
          </div>
          <p class="theme-item">{{ meeting.topic }}</p>
          <p class="client-item">{{ meeting.client.name }}</p>
        </li>
      </ul>
    </div>

    <loading-indicate v-else class="indicator"></loading-indicate>
  </main>
</template>

<script>
import buttons from "@/components/UI/button-add.vue";
import api from "@/api";
import LoadingIndicate from "@/components/UI/LoadingIndicate.vue" ;

export default {
  components: { buttons, LoadingIndicate},

  data() {
    return {
      meetings: [
        {
          client: {},
          place: {},
        },
      ],
      nearMeeting: {
        contact: {},
        client: {},
        place: {},
      },
      isLoadedMeetings: false
    };
  },

  async mounted() {
    this.meetings = await api.getMeetings();
    this.prepareMeetingsForDisplay();
    this.nearMeeting = this.meetings[0];
    this.isLoadedMeetings = true
  },

  methods: {
    async prepareMeetingsForDisplay() {
      const clients = await api.getClients();
      const mathchIdToClient = new Map();
      clients.forEach((item) => {
        mathchIdToClient.set(item.id, item);
      });

      const places = await api.getPlaces();
      const matchIdToPlace = new Map();
      places.forEach((item) => {
        matchIdToPlace.set(item.id, item);
      });

      this.meetings.forEach((meeting) => {
        meeting.client = mathchIdToClient.get(meeting.client);
        meeting.start = meeting.start.substr(0, 5);
        meeting.end = meeting.end.substr(0, 5);
        meeting.place = matchIdToPlace.get(meeting.place);
      });

      const contact = await api.getContact(this.meetings[0].contact);
      this.meetings[0].contact = contact;
    },
  },
};
</script>

<style scoped>
@import url('../assests/fonts/Exo_2/stylesheet.css');

* {
  font-family: 'Exo 2';
}

ul {
  list-style: none;
  padding: 0;
}

button {
  background-color: white;
  border: none;
  margin: 0;
  padding: 0;
}

p {
  font-size: 20px;
  margin: 0;
}

li p {
  font-family: 'Exo 2';
  font-weight: 700;
  font-size: 0.75rem;
  margin-left: 17px;
  margin-right: 17px;
}

.theme-item {
  font-size: 0.88rem;
  overflow: auto;
  word-break: break-all;
}

.near-meet-text {
  margin-top: 31px;
  margin-bottom: 31px;
}

.near-meet {
  margin-right: 40px;
  margin-left: 10.14%;
  flex: 2;
}

.green-button-add {
  display: flex;
  justify-content: center;
}

.data-time-place-item {
  display: flex;
  justify-content: space-between;
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
  background: #f1fbf2;
  border: 0.714062px solid #f1fbf2;
  box-shadow: 0px 2.85625px 2.85625px rgba(0, 0, 0, 0.25);
  border-radius: 10px;
  margin-bottom: 9px;
  margin-top: 9px;
}

.data-and-time-window {
  display: flex;
}

.near-meet {
  display: flex;
  flex-direction: column;
  height: 80vh;
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
  height: 80vh;
}

.client-wiindow,
.contact-window,
.theme-window,
.data-window,
.place-window {
  margin-left: 40px;
}

.place-window {
  margin-bottom: 60px;
}

main {
  display: flex;
  justify-content: space-around;
  width: 100vw;
}

.theme-window {
  margin-bottom: 123px;
  font-weight: 700;
  font-size: 2rem;
  word-break: break-all;
  overflow: auto;
}

.list-all-meet {
  height: 86%;
  background: #00b268;
  border-radius: 20px;
  overflow: auto;
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

.client-wiindow {
  margin-bottom: 26px;
}

.list-meet {
  flex: 1;
  margin-right: 10.14%;
  height: 87vh;
}

.data-window,
.time-window,
.place-window {
  font-weight: bolder;
}

.pencil-icon {
  max-width: 100%;
}

/* Responsive layout */

@media (max-width: 1200px) {
  main {
    flex-direction: column;
    align-content: center;
    background-color: #e2eee3;
  }

  .list-meet {
    margin-left: 2rem;
    margin-right: 40px;
    margin-left: 40px;
    height: auto;
  }

  .list-all-meet {
    overflow: none;
    width: 100%;
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
  .client-wiindow {
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
}

@media (max-width: 992px) {
  .container-window {
    width: 100%;
    margin-bottom: 0;
    margin-top: 1rem;
    height: 250px;
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
    margin-top: 1rem;
  }

  main p {
    font-size: 1rem;
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

  .client-wiindow,
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
    margin-top: 0.5rem;
    gap: 0;
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

</style>
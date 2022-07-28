<template>
  <main>
    <div class="container">
      <div class="near-meet">
        <h4 class="near-meet-text">Ближайшая встреча:</h4>

        <div class="window-near-meet">
          <div class="container-window">
            <button class="pencil-meet">
              <img src="/svg/pencil.svg" alt="" class="pencil-icon" />
            </button>
            <p class="client-wiindow">{{nearMeeting.client.name}}</p>
            <div class="window-tel-and-contact">
              <p class="contact-window">
                {{nearMeeting.contact.first_name}} {{nearMeeting.contact.second_name}} {{nearMeeting.contact.third_name}}
              </p>
              <p class="tel-window">{{nearMeeting.contact.phone}}</p>
            </div>
            <p class="theme-window">{{nearMeeting.topic}}</p>
            <div class="data-and-time-window">
              <p class="data-window">{{nearMeeting.date}}</p>
              <p class="time-window">{{nearMeeting.start}} - {{nearMeeting.end}}</p>
            </div>
            <p class="place-window">{{nearMeeting.place.name}}</p>
          </div>
        </div>
        <div class="green-button-add">
          <buttons></buttons>
        </div>
      </div>

      <div class="list-meet">
        <h4 class="list-meet-text">Список встреч</h4>

        <div class="all-meet">
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
      </div>
    </div>
  </main>
</template>

<script>
import buttons from "@/components/UI/button-add.vue";
import api from "@/api";

export default {
  components: { buttons },

  data() {
    return {
      meetings: [
        {
          client: {},
          place: {}
        }
      ],
      nearMeeting: {
        contact: {},
        client: {},
        place: {}
      }
    };
  },

  async mounted() {
    this.meetings = await api.getMeetings();
    this.prepareMeetingsForDisplay();
    this.nearMeeting = this.meetings[0];
  },

  methods: {
    async prepareMeetingsForDisplay() {
      
      const clients = await api.getClients();
      const mathchIdToClient = new Map();
      clients.forEach(item => {
        mathchIdToClient.set(item.id, item);
      });

      const places = await api.getPlaces();
      const matchIdToPlace = new Map();
      places.forEach(item => {
        matchIdToPlace.set(item.id, item);
      });

      this.meetings.forEach(meeting => {
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

<style>
.window-near-meet {
  width: 774px;
  height: 567px;
  background: #ffffff;
  border: 1px solid #bdbdbd;
  border-radius: 30px;
  margin-bottom: 60px;
}

p {
  font-size: 20px;
  margin: 0;
}
.near-meet-text {
  margin-top: 31px;
  margin-bottom: 31px;
}

.near-meet {
  margin-right: 40px;
}
.green-button-add {
  display: flex;
  justify-content: center;
}

.data-time-place-item {
  display: flex;
}

.window-tel-and-contact {
  display: flex;
}

.pencil-meet {
  margin-top: 38px;

  margin-right: 23px;
  margin-left: auto;
  display: flex;
}

.all-meet-item {
  width: 356px;
  height: 127px;
  background: #f1fbf2;
  border: 0.714062px solid #f1fbf2;
  box-shadow: 0px 2.85625px 2.85625px rgba(0, 0, 0, 0.25);
  border-radius: 10px;
}

.all-meet-item {
  margin-bottom: 19px;
}

.contact-window {
  margin-right: 334px;
  margin-bottom: 53px;
}

.data-and-time-window {
  display: flex;
}

.container-window {
  margin-left: 40px;
}

.container {
  display: flex;
}

.all-meet {
  width: 406px;
  height: 766px;
  background: #00b268;
  border-radius: 20px;
  overflow: auto;
}
.theme-window {
  margin-bottom: 123px;
}

.list-all-meet {
  margin-left: 16px;
}

.data-window {
  margin-bottom: 18px;
  margin-right: 439px;
}
.client-wiindow {
  margin-bottom: 26px;
}
</style>
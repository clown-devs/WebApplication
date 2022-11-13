<template>
  <div class="day-container">
    <ul class="grid-container">
      <li class="grid-row" ref="infoBox" v-for="item in this.cells" :key="item" >
        <div class="row-cell-container">
          <span class="time-item"> {{ item }} </span>
          <div class="half-time-line"></div>
        </div>

        <div
            class="meeting-event-element"
            v-if="meetingEvents.get(item) !== undefined"
            :style="{top: meetingEvents.get(item)['start-coordinate']}"
        >
          {{ meetingEvents.get(item).topic }}
        </div>
      </li>
    </ul>
  </div>
</template>

<script>

import Popup from "@/components/UI/Popup";
export default {
  name: "CalendarDaySlice",
  components: {Popup},
  data() {
    return {
      lineHeight: '',

      cells: [],

      meetingEvents: new Map(),
    }
  },

  mounted() {
    this.createTimeGrid();

    let meetings = [
      {
        start: '00:00',
        end: '3:25',
        topic: 'Govno iz jopi'
      },

      {
        start: '05:13',
        end: '06:13',
        topic: 'Clown devs'
      },

      {
        start: '11:16',
        end: '14:00',
        topic: 'Bozhedom'
      }
    ];

    this.handleMeetingEvents(meetings);

    this.$nextTick(function () {
      this.lineHeight = this.matchHeight();
    })

  },



  methods: {
    createTimeGrid() {
      for (let i = 0; i < 24; i++) {
        this.cells.push(i);
      }
    },

    handleMeetingEvents(meetingEvents) {
      meetingEvents.forEach(event => {
        let key;
        let hours = event.start.slice(0, 2);

        key = hours[0] === '0' ? +hours[1] : +hours;

        let minutes = +event.start.slice(3, event.start.length);
        event['start-coordinate'] = String(-100 + Math.round(minutes / 60 * 100)) + '%';

        this.meetingEvents.set(key, event);
      });
    },

    matchHeight () {
      return this.$refs.infoBox[0].clientHeight;
    },
  }
}
</script>

<style scoped>
.day-container {
  border: 1px solid black;
}

.grid-container {
  list-style: none;
  margin: 0;
}

.row-cell-container {
  display: flex;
  flex-direction: row;
  height: 100%;
  align-items: center;
}

.grid-row {
  height: 60px;
  border-bottom: 1px solid #BDBDBD;
}

.time-item {
  font-family: "Exo 2", serif;
  font-weight: 500;
  font-size: 1.25rem;
}

.meeting-event-element {
  margin-left: 2rem;
  background: rgba(255, 255, 255, 0.5);
  box-shadow: 0 2px 2px rgba(0, 0, 0, 0.25);
  border-radius: 5px;
  position: relative;
  top: -100%;
}

.half-time-line {
  border: 1px dashed #BDBDBD;
  width: 100%;
  margin-left: 25px;
}
</style>
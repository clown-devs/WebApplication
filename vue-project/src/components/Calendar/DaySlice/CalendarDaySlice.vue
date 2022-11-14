<template>
  <div class="day-container" ref="gridLine">
    <ul class="grid-container">
      <li class="grid-row" ref="infoBox" v-for="item in this.cells" :key="item" >
        <div class="row-cell-container">
          <span class="time-item"> {{ item }} </span>
          <div class="half-time-line"></div>
        </div>

        <div
            class="meeting-event-element"
            v-if="meetingEvents.get(item) !== undefined"
            :style="{
              top: meetingEvents.get(item)['start-coordinate'],
              height: meetingEvents.get(item)['end-coordinate']
            }"
        >
          {{ meetingEvents.get(item).topic }}
        </div>
      </li>
      <li
          class="red-line"
          :style="{top: getRedLinePosition()}"
          v-if="isToday"
          :key="updateRedLineComponent"
      ></li>
    </ul>
  </div>
</template>

<script>
import api from "@/api";
import {HALF_MINUTES} from "@/helpers/constant";

export default {
  name: "CalendarDaySlice",

  props: {
    selectedPlace: {
      type: Object,
      default: {id: 1}
    },

    selectedDate: {
      type: Date,
      default: new Date()
    }
  },

  data() {
    return {
      lineHeight: 0,

      cells: [],

      meetingEvents: new Map(),

      updateRedLineComponent: 0,
    }
  },

  created() {
    window.setInterval(() => {
      this.updateRedLineComponent += 1;
    }, HALF_MINUTES);
  },

  async mounted() {
    this.createTimeGrid();

    let meetings = await api.getMeetingsByDate(this.selectedDate);

     await this.$nextTick(function () {
      this.lineHeight = this.matchHeight();
      this.handleMeetingEvents(meetings);
      this.scrollToStartWorkTime();
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
        let hours = event.start.slice(0, 2);
        let key = hours[0] === '0' ? +hours[1] : +hours;

        event['start-coordinate'] = this.calculateStartCoordinateForEvent(
            event.start
        );

        event['end-coordinate'] = this.calculateEndCoordinateForEvent(
            event.start,
            event.end
        );

        this.meetingEvents.set(key, event);
      });
    },

    matchHeight () {
      return this.$refs.infoBox[0].clientHeight;
    },

    calculateStartCoordinateForEvent(startTime) {
      startTime = this.deleteSecondsInTime(startTime);

      let minutes = +startTime.slice(3, startTime.length);
      return String(-100 + Math.round(minutes / 60 * 100)) + '%';
    },

    calculateEndCoordinateForEvent(startTime, endTime) {
      startTime = this.deleteSecondsInTime(startTime);
      endTime = this.deleteSecondsInTime(endTime);

      const startHours = +startTime.slice(0, 2);
      const startMinutes = +startTime.slice(3, startTime.length);

      const endHours = +endTime.slice(0, 2);
      const endMinutes = +endTime.slice(3, endTime.length);

      const startDate = new Date();
      startDate.setHours(startHours);
      startDate.setMinutes(startMinutes);

      const endDate = new Date();
      endDate.setHours(endHours);
      endDate.setMinutes(endMinutes);

      const delta = ( (endDate - startDate) / 1000 / 60 / 60 );
      const result = delta * this.lineHeight;
      return String(result) + 'px';
    },

    deleteSecondsInTime(timeString) {
      //format: hh:mm:ss
      return timeString.slice(0, timeString.length - 3);
    },

    scrollToStartWorkTime() {
      this.$refs.gridLine.scrollTop = this.$refs.gridLine.scrollHeight / 3;
    },

    getRedLinePosition() {
      const today = new Date();
      today.setHours(0, 0, 0, 0);

      const now = new Date();

      const delta = (now - today) / 1000 / 60 / 60 / 24 * 100;
      return String(delta) + '%';
    }

  },

  watch: {
    async selectedDate(newDate) {
      let meetings = await api.getMeetingsByDate(newDate);
      this.meetingEvents.clear();
      this.handleMeetingEvents(meetings);
    },

    async selectedPlace(newPlace) {
      let meetings = await api.getMeetingsInPlaceByDate(this.selectedDate, newPlace);
      this.meetingEvents.clear();
      this.handleMeetingEvents(meetings);
    },
  },

  computed: {
    isToday() {
      let selectedDate = this.selectedDate;
      selectedDate.setHours(0, 0, 0, 0);

      let today = new Date();
      today.setHours(0, 0, 0, 0);

      return (selectedDate - today) === 0;
    }
  }

}
</script>

<style scoped>
.day-container {
  border-right: 1px solid #BDBDBD;
  max-width: 1199px;
  height: 100%;
  overflow: auto;
}

.grid-container {
  list-style: none;
  margin: 0;
  padding: 0;
  position: relative;
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

.red-line {
  background-color: #f00;
  height: 1px;
  position: absolute;
  width: 100%;
}
</style>
<template>
  <div class="day-container" ref="gridLine">
    <ul class="grid-container">
      <li class="grid-row" ref="infoBox" v-for="item in this.cells" :key="item" >
        <div class="row-cell-container">
          <div class="divider-cell-line-container" ref="forStartCalculate">
            <span
                class="time-item"
                ref="lineTime"
                v-show="!hideTimeInLine.get(item)"
            > {{ prepareHoursForDisplay(item) }} </span>
            <div class="divider-cell-line" ref="dividerLine"></div>
          </div>
        </div>

        <div
            class="meeting-event-element"
            v-if="meetingEvents.get(item) !== undefined"
            :style="{
              top: meetingEvents.get(item)['start-coordinate'],
              height: meetingEvents.get(item)['end-coordinate']
            }"
        >
          <label class="topic-meeting-event">{{ meetingEvents.get(item).topic }}</label>
          <label class="place-meeting-event"> {{ meetingEvents.get(item).place_name }} </label>
        </div>
      </li>
      <li
          class="red-line-container"
          :style="{top: getRedLinePosition()}"
          v-if="isToday"
          :key="updateRedLineComponent"
      >
        <label class="current-time"> {{ displayCurrentTime() }} </label>
        <div class="red-point" ref="redPoint"></div>
        <div class="red-line"></div>
      </li>
    </ul>
  </div>
</template>

<script>
import api from "@/api";
import {MINUTE} from "@/helpers/constant";

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
      gridLineHeight: 0,

      dividerLineDivHeight: 0,

      redPointHeight: 0,

      dividerLineHeight: 0,

      cells: [],

      meetingEvents: new Map(),

      updateRedLineComponent: 0,

      hideTimeInLine: new Map(),

      now: new Date()
    }
  },

  created() {
    window.setInterval(() => {
      this.updateRedLineComponent += 1;
      this.now = new Date();
    }, MINUTE);
  },

  async mounted() {
    this.createTimeGrid();

    let meetings = await api.getMeetingsByDate(this.selectedDate);

    for (let i = 0; i < 24; i++) {
      this.hideTimeInLine.set(i, false);
    }

     await this.$nextTick(function () {
      this.gridLineHeight = this.getClientHeight(this.$refs.infoBox[0]);
      this.dividerLineDivHeight = this.getClientHeight(this.$refs.forStartCalculate[0]);
      this.redPointHeight = this.getClientHeight(this.$refs.redPoint);
      this.dividerLineHeight = this.getClientHeight(this.$refs.dividerLine[0]);
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

    getClientHeight(ref) {
      return ref.clientHeight;
    },

    calculateStartCoordinateForEvent(startTime) {
      startTime = this.deleteSecondsInTime(startTime);

      let minutes = +startTime.slice(3, startTime.length);

      return String(
          -100 + (minutes / 60 * 100) + this.getStartOffsetByDividerLineInPercent() * 100
      ) + '%';
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
      const result = delta * this.gridLineHeight;
      return String(result) + 'px';
    },

    deleteSecondsInTime(timeString) {
      //format: hh:mm:ss
      return timeString.slice(0, timeString.length - 3);
    },

    scrollToStartWorkTime() {
      this.$refs.gridLine.scrollTop = this.$refs.gridLine.scrollHeight / 3
          - this.dividerLineDivHeight / 2;
    },

    getRedLinePosition() {
      const today = new Date();
      today.setHours(0, 0, 0, 0);

      const now = new Date();

      let delta = (now - today)  * 100 / 1000 / 60 / 60 / 24;
      delta += this.getStartOffsetByDividerLineInPercent() * 100 / 24;
      delta -= this.getOffsetForRedLineByRedCircle() * 100 / 24;

      return String(delta) + '%';
    },

    prepareHoursForDisplay(hours) {
      return hours < 10 ? '0' + hours + ':00' : hours + ':00';
    },

    getStartOffsetByDividerLineInPercent() {
      return (this.dividerLineDivHeight  - this.dividerLineHeight) / this.gridLineHeight / 2 ;
    },

    getOffsetForRedLineByRedCircle() {
      return this.redPointHeight / this.gridLineHeight / 2;
    },

    displayCurrentTime() {
      const currentDate = new Date();
      const hours = currentDate.getHours();
      const minutes = currentDate.getMinutes();

      const hoursString = hours < 10 ? '0' + String(hours) : String(hours);
      const minutesString = minutes < 10 ? '0' + String(minutes) : String(minutes);

      return hoursString + ':' + minutesString;
    },
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

    now(newMoment) {
      if (newMoment.getHours() + 1 === 24) return;

      if (newMoment.getMinutes() >= 50) {
        this.hideTimeInLine.set(newMoment.getHours() + 1, true);
        return;
      }

      if (newMoment.getMinutes() <= 10) {
        this.hideTimeInLine.set(newMoment.getHours(), true);
        return;
      }

      this.hideTimeInLine.set(newMoment.getHours(), false);
    }
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
  flex-direction: column;
  height: 100%;
  justify-content: space-between;
}

.grid-row {
  height: 60px;
}

.time-item {
  font-family: "Exo 2", serif;
  font-weight: 500;
  font-size: 12px;
  margin-right: 10px;
  width: 33px;
}

.current-time {
  font-family: "Exo 2", serif;
  font-weight: 500;
  font-size: 12px;
  margin-right: 4px;
  width: 33px;
  color: red;
}

.meeting-event-element {
  margin-left: 41px;
  background: rgba(255, 255, 255, 0.5);
  box-shadow: 0 2px 2px rgba(0, 0, 0, 0.25);
  border-radius: 5px;
  position: relative;
  top: -100%;
  display: flex;
  flex-direction: column;
  cursor: pointer;
}

.red-line-container {
  position: absolute;
  width: 100%;
  display: flex;
  align-items: center;
  flex-direction: row;
}

.red-line {
  background-color: #f00;
  height: 1px;
  width: 100%;
}

.red-point {
  border-radius: 50px;
  background-color: red;
  height: 12px;
  width: 12px;
}

.topic-meeting-event {
  margin: 2px 0 0 2rem;
  font-family: "Exo 2", serif;
  font-weight: 700;
  font-size: 18px;
  cursor: pointer;
}

.place-meeting-event {
  margin: 2px 0 0 2rem;
  font-family: "Exo 2", serif;
  font-weight: 700;
  font-size: 16px;
  cursor: pointer;
}

.divider-cell-line {
  height: 1px;
  width: 100%;
  background-color: #bdbdbd;
}

.divider-cell-line-container {
  display: flex;
  flex-direction: row;
  align-items: center;
  height: 12px;
  width: 100%;
}
/* Animations and hovers */
</style>
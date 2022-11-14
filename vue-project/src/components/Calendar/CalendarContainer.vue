<template>
  <div class="calendar-main-container">
    <h1>Календарь</h1>
    <div class="calendar-header-level">
      <div class="blank-space"></div>
      <calendar-header
          class="calendar-header"
          @switchSlice="showSelectedSlice"
          @selectedDate="handleSelectDate"
      ></calendar-header>
    </div>

    <div class="calendar-body-level">
      <div class="blank-mini-calendar">
        <meet-places @takePlace="handleTakePlace"></meet-places>
      </div>
      <div class="slice-container">
        <calendar-day-slice
            class="day-slice-container"
            v-if="isPressedDayBtn"
            :selected-place="selectedPlace"
            :selected-date="selectedDate"
        ></calendar-day-slice>

        <calendar-week-slice v-if="isPressedWeekBtn"></calendar-week-slice>
        <calendar-month-slice v-if="isPressedMonthBtn"></calendar-month-slice>
      </div>
    </div>
  </div>
</template>

<script>
import CalendarDaySlice from "@/components/Calendar/DaySlice/CalendarDaySlice";
import CalendarWeekSlice from "@/components/Calendar/WeekSlice/CalendarWeekSlice";
import CalendarMonthSlice from "@/components/Calendar/MonthSlice/CalendarMonthSlice";
import CalendarHeader from "@/components/Calendar/CalendarHeader";
import MeetPlaces from "@/components/Calendar/SmallCalendar/MeetPlaces";

export default {
  name: "CalendarContainer",

  components: {
    CalendarDaySlice,
    CalendarWeekSlice,
    CalendarMonthSlice,
    CalendarHeader,
    MeetPlaces,
  },

  data() {
    return {
        isPressedDayBtn: false,
        isPressedWeekBtn: false,
        isPressedMonthBtn: false,
        selectedPlace: {},
        selectedDate: new Date(),
      }
    },

  methods: {
    showSelectedSlice(slicesState) {
      this.isPressedDayBtn = slicesState.isPressedDayBtn;
      this.isPressedMonthBtn = slicesState.isPressedMonthBtn;
      this.isPressedWeekBtn = slicesState.isPressedWeekBtn;
    },

    handleTakePlace(place) {
      this.selectedPlace = place;
    },

    handleSelectDate(date) {
      this.selectedDate = date;
    }
  },

}

</script>

<style scoped>
.calendar-main-container {
  height: 100%;
}

h1 {
  font-family: "Exo 2", serif;
  font-weight: 700;
  font-size: 2rem;
  margin-left: 10.14%;
  margin-bottom: 30px;
  margin-top: 0;
  padding-top: 30px;
}

.calendar-header-level {
  display: flex;
  flex-direction: row;
}

.blank-space {
  flex: 1;
}

.calendar-header {
  flex: 2;
  margin-right: 10.14%;
  padding-right: 5px;
}

.calendar-body-level {
  display: flex;
  flex-direction: row;
  height: 70%;
}

.slice-container {
  flex: 2;
  margin-right: 10.14%;
  /*overflow-y: auto;*/
}

.blank-mini-calendar {
  flex: 1;
}
</style>
<template>
  <div class="calendar-main-container">
    <div class="calendar-body-level">
      <div class="blank-mini-calendar">
        <meet-places @takePlace="handleTakePlace"></meet-places>
      </div>
      <div class="slice-container">
        <calendar-header
            class="calendar-header"
            @switchSlice="showSelectedSlice"
            @selectedDate="handleSelectDate"
        ></calendar-header>
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
import SmallCalendar from "@/components/Calendar/SmallCalendar/SmallCalendar";

export default {
  name: "CalendarContainer",

  components: {
    SmallCalendar,
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
      propsSelectedDay: new Date(),
      emitSelectedDay: new Date(),
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
    },
    takeDayOnCalendar(day) {
      this.emitSelectedDay = day;
    }
  },

}

</script>

<style scoped>
.calendar-main-container {
  height: 100%;
  margin-left: 10.14%;
  margin-right: 10.14%;
  margin-top: 50px;
}

.calendar-body-level {
  padding-top: 30px;
  display: flex;
  flex-direction: row;
  height: 75%;
}

.slice-container {
  flex: 2;
  margin-right: 10.14%;
}

.blank-mini-calendar {
  /*display: flex;*/
  /*flex-direction: column;*/
  /*margin-right: 62px;*/
  /*justify-content: space-between;*/
  flex: 1;
}

.small-calendar {
  height: 47%;
}

.meet-places {
  height: 47%;
}
</style>
<template>
  <calendar-day-slice v-show="isPressedDayBtn"></calendar-day-slice>
  <calendar-week-slice v-show="isPressedWeekBtn"></calendar-week-slice>
  <calendar-month-slice v-show="isPressedMonthBtn"></calendar-month-slice>
  <small-calendar></small-calendar>
  <meet-places @takePlace='onTest' ></meet-places>
</template>

<script>
import CalendarDaySlice from "@/components/Calendar/CalendarDaySlice";
import CalendarWeekSlice from "@/components/Calendar/CalendarWeekSlice";
import CalendarMonthSlice from "@/components/Calendar/CalendarMonthSlice";
import SmallCalendar from "@/components/Calendar/SmallCalendar/SmallCalendar";
import MeetPlaces from "@/components/Calendar/SmallCalendar/MeetPlaces";

export default {
  name: "CalendarContainer",

  components: {
    SmallCalendar,
    CalendarDaySlice,
    CalendarWeekSlice,
    CalendarMonthSlice,
    MeetPlaces
  },

  data() {
    return {
      isPressedDayBtn: true,
      isPressedWeekBtn: false,
      isPressedMonthBtn: false,
      selectedDate: new Date(),
    }
  },

  methods: {
    activateSliceBtn(isPressedButtonsState = {day: false, week: false, month: false}) {
      this.isPressedMonthBtn = isPressedButtonsState.month;
      this.isPressedWeekBtn = isPressedButtonsState.week;
      this.isPressedDayBtn = isPressedButtonsState.day;
    },
    onTest(data) {
      console.log(data)
    }
  },

  watch: {
    isPressedDayBtn(newState) {
      if (newState) {
        this.activateSliceBtn({day: true, month: false, week: false});
      }
    },

    isPressedWeekBtn(newState) {
      if (newState) {
        this.activateSliceBtn({day: false, week: true, month: false});
      }
    },

    isPressedMonthBtn(newState) {
      if (newState) {
        this.activateSliceBtn({day: false, week: false, month: true});
      }
    },
  }
}
</script>

<style scoped>

</style>
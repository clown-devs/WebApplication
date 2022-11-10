<template>
  <calendar-header @switchSlice="showSelectedSlice"></calendar-header>
  <calendar-day-slice v-show="isPressedDayBtn"></calendar-day-slice>
  <calendar-week-slice v-show="isPressedWeekBtn"></calendar-week-slice>
  <calendar-month-slice v-show="isPressedMonthBtn"></calendar-month-slice>
</template>

<script>
import CalendarDaySlice from "@/components/Calendar/CalendarDaySlice";
import CalendarWeekSlice from "@/components/Calendar/CalendarWeekSlice";
import CalendarMonthSlice from "@/components/Calendar/CalendarMonthSlice";
import CalendarHeader from "@/components/Calendar/CalendarHeader";

export default {
  name: "CalendarContainer",

  components: {
    CalendarDaySlice,
    CalendarWeekSlice,
    CalendarMonthSlice,
    CalendarHeader
  },

  data() {
    return {
        isPressedDayBtn: false,
        isPressedWeekBtn: false,
        isPressedMonthBtn: false,
      }
    },

  methods: {
    showSelectedSlice(slicesState) {
      console.log(slicesState);
      this.isPressedDayBtn = slicesState.isPressedDayBtn;
      this.isPressedMonthBtn = slicesState.isPressedMonthBtn;
      this.isPressedWeekBtn = slicesState.isPressedWeekBtn;
    }
  },

}

</script>

<style scoped>

</style>
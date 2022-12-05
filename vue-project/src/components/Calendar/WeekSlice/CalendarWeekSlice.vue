<template>
  <div class="week-container">
    <ul class="days-container">
      <time-line :set-grid-line-height="gridRowHeight"></time-line>
      <li class="day-item" v-for="day in 7">
        <calendar-day-slice
            :turned-on-scroll="false"
            :hide-time="true"
            :set-grid-row-height="gridRowHeight"
        ></calendar-day-slice>
      </li>
    </ul>
  </div>
</template>

<script>
import CalendarDaySlice from "@/components/Calendar/DaySlice/CalendarDaySlice";
import TimeLine from "@/components/Calendar/WeekSlice/TimeLine";

export default {
  name: "CalendarWeekSlice",

  components: {
    TimeLine,
    CalendarDaySlice,
  },

  data() {
    return {
      gridRowHeight: '60px',
    }
  },
}
</script>

<style scoped>
.week-container {
  height: 100%;
  overflow: auto;
}

.days-container {
  list-style: none;
  display: flex;
  flex-direction: row;
  padding: 0;
  margin: 0;
}

.day-item {
  flex: 1;
}
</style>
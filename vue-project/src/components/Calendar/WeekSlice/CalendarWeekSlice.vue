<template>
  <div class="week-container">
    <ul class="days-container">
      <li class="day-item" v-for="day in 7">
        <calendar-day-slice></calendar-day-slice>
      </li>
    </ul>
  </div>
</template>

<script>
import CalendarDaySlice from "@/components/Calendar/DaySlice/CalendarDaySlice";

export default {
  name: "CalendarWeekSlice",

  components: {
    CalendarDaySlice
  }
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
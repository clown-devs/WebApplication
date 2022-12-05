<template>
<div class="time-line-container">
  <div class="time-line-row" v-for="hours in 24" :style="{height: setGridLineHeight}">
    <span class="time-item">
      {{ displayHours(hours - 1) }}
    </span>
  </div>
</div>
</template>

<script>
import displayTime from "@/helpers/display-time.js";

export default {
  name: "TimeLine",

  components: {
    displayTime,
  },

  props: {
    setGridLineHeight: {
      type: String,
      default: '60px',
    }
  },

  methods: {
    displayHours(hours) {
      return displayTime.prepareHoursForDisplay(hours);
    }
  }
}

</script>

<style scoped>
.time-line-container {
  display: flex;
  flex-direction: column;
}

.time-item {
  font-family: "Exo 2", serif;
  font-weight: 500;
  font-size: 12px;
  margin-right: 10px;
  width: 33px;
  margin-top: 0;
  height: 12px;
  vertical-align: text-top;
}

.time-line-row {
  display: flex;
  flex-direction: column;
  align-self: flex-start;
}
</style>
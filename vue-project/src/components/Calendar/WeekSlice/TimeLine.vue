<template>
<div class="time-line-container">
  <span class="time-item" v-for="hours in 23">
    {{ displayHours(hours) }}
  </span>
</div>
</template>

<script>
import displayTime from "@/helpers/display-time.js";

export default {
  name: "TimeLine",

  components: {
    displayTime,
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
}
</style>
<template>
<div class="time-line-container">
  <span class="time-item" v-for="time in 23"> {{ time }} </span>
</div>
</template>

<script>

export default {
  name: "TimeLine",

  props: {
    time: {
      type: String,
      default: '05:00'
    }
  }
}

</script>

<style scoped>
.time-item {
  font-family: "Exo 2", serif;
  font-weight: 500;
  font-size: 12px;
  margin-right: 10px;
  width: 33px;
}
</style>
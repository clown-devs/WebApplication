<template>
  <div class="tooltip-wrapper">
    <slot></slot>
    <div class="tooltip-element" v-if="isShow">
      <div class="tooltip-header"></div>
      <div class="tooltip-content-header"></div>
      <div class="tooltip-content">
        <div class="content-title-container">
          <label class="content-title">{{ event.topic }}</label>
        </div>
        <div class="event-date-container">
          <label class="start time"> {{ deleteSecondsInTime(event.start) }}</label>
          <label class="event-date"> {{ event.date.toLocaleDateString()}},</label>
          <label class="end time"> {{ getWeekDay(event.date) }} - {{ deleteSecondsInTime(event.end) }}</label>
        </div>
        <label class="event-place">{{ event.place_name }}</label>
        <label class="event-creater"> {{ event.creator_name }}</label>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "EventTooltip",

  props: {
    event: {
      type: Object,
      default: {}
    },

    isShow: {
      type: Boolean,
      default: false
    }
  },

  data() {
    return {}
  },

  mounted() {
    this.event.date = new Date(this.event.date);
  },

  methods: {
    getWeekDay(date) {
      const shorts = {
        'понедельник': 'Пн',
        'вторник': 'Вт',
        'среда': 'Ср',
        'четверг': 'Чт',
        'пятница': 'Пт',
        'суббота': 'Сб',
        'воскресенье': 'Вс'
      };

      const option = {weekday: 'long'};
      let weekday = date.toLocaleString('ru', option);
      return shorts[weekday];
    },

    deleteSecondsInTime(time) {
      return time.slice(0, 5);
    }
  },
}
</script>

<style scoped>
.tooltip-wrapper {
  position: relative;
  height: 100%;
}

.tooltip-element {
  position: absolute;
  left: 50%;
  top: 50%;
  z-index: 1000;
  opacity: .9;
}

.tooltip-element {
  min-width: 400px;
  max-height: 220px;
  background-color: #E1EEE3;
  margin-left: 5px;
}

.tooltip-header {
  height: 20px;
  background-color: #00B268;
  box-shadow: 0 0 2px rgba(0, 0, 0, 0.25);
}

.tooltip-content {
  height: 200px;
  box-shadow: 0 2px 2px rgba(0, 0, 0, 0.25);
  display: flex;
  flex-direction: column;
  align-items: center;
}

.content-title {
  margin-top: 33px;
  font-family: "Exo 2", serif;
  font-weight: 700;
  font-size: 20px;
  color: #009657;
}

.content-title-container {
  display: flex;
  align-items: center;
  justify-content: center;
}

.event-date-container {
  margin-top: 15px;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  width: 50%;
  font-family: "Exo 2", serif;
  font-weight: 500;
  font-size: 14px;
}

.event-place,
.event-creater {
  margin-top: 15px;
  font-family: "Exo 2", serif;
  font-weight: 500;
  font-size: 14px;
}

.content-title {
  margin: 0 2rem 0 2rem;
}
</style>
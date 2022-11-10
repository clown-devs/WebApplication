<template>
  <div class="calendar-header-container">
    <div class="calendar-header-btns">
      <button class="today-btn slice" @click="pressTodayBtn">Сегодня</button>

      <button
          class="day-btn slice" @click="pressDayBtn"
          :class="{activate: this.isPressedDayBtn}"
      >День</button>

      <button
          class="week-btn slice"
          @click="pressWeekBtn"
          :class="{activate: this.isPressedWeekBtn}"
      >Неделя</button>

      <button
          class="month-btn slice"
          @click="pressMonthBtn"
          :class="{activate: this.isPressedMonthBtn}"
      >Месяц</button>

      <button class="prev-date-btn"></button>
      <button class="next-date-btn"></button>
      <button class="free-place-btn"></button>
    </div>
    <div class="design-line-element"></div>
  </div>
</template>

<script>
export default {
  name: "CalendarHeader",

  data() {
    return {
      isPressedDayBtn: true,
      isPressedWeekBtn: false,
      isPressedMonthBtn: false,
      selectedDate: new Date(),
    }
  },

  mounted() {
    this.$emit("switchSlice", this.sliceState);
  },

  methods: {
    activateSliceBtn(isPressedButtonsState = {day: false, week: false, month: false}) {
      this.isPressedMonthBtn = isPressedButtonsState.month;
      this.isPressedWeekBtn = isPressedButtonsState.week;
      this.isPressedDayBtn = isPressedButtonsState.day;
    },

    pressDayBtn() {
      this.isPressedDayBtn = !this.isPressedDayBtn;
      this.activateSliceBtn({day: true, month: false, week: false});
      this.$emit("switchSlice", this.sliceState);
    },

    pressWeekBtn() {
      this.isPressedWeekBtn = !this.isPressedWeekBtn;
      this.activateSliceBtn({day: false, week: true, month: false});
      this.$emit("switchSlice", this.sliceState);
    },

    pressMonthBtn() {
      this.isPressedMonthBtn = !this.isPressedMonthBtn;
      this.activateSliceBtn({day: false, week: false, month: true});
      this.$emit("switchSlice", this.sliceState);
    },

    pressTodayBtn() {
      this.isPressedDayBtn = true;
      this.activateSliceBtn({day: true, month: false, week: false});
      this.selectedDate = new Date();
      this.$emit("switchSlice", this.sliceState);
    }
  },

  computed: {
    sliceState() {
      return {
        "isPressedDayBtn": this.isPressedDayBtn,
        "isPressedWeekBtn": this.isPressedWeekBtn,
        "isPressedMonthBtn": this.isPressedMonthBtn,
        "selectedDate": this.selectedDate
      }
    }
  },

}
</script>

<style scoped>
.calendar-header-btns {
  display: flex;
  flex-direction: row;
  flex: 1;
  margin-bottom: 13px;
}

.slice {
  font-family: "Exo 2", serif;
  font-weight: 700;
  font-size: 1.25rem;
  margin-right: 30px;
  color: #424242;
  background: none;
  border: none;
  cursor: pointer;
  transition-duration: 0.5s;
}

.design-line-element {
  border: 1px solid #000000;
}

.activate {
  color: #00B268;
  transition-duration: 0.5s;
}
</style>
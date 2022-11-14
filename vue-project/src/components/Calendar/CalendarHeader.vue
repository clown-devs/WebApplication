<template>
  <div class="calendar-header-container">
    <div class="calendar-header-btns">
      <button class="today-btn" @click="pressTodayBtn">Сегодня</button>

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

      <div class="select-time-container">
        <label class="selected-date"> {{ this.formatDate(selectedDate) }} </label>
        <i class="material-icons prev-date-btn"
           @click="selectedNewDate(-1)"
        > arrow_back_ios </i>
        <i class="material-icons next-date-btn"
           @click="selectedNewDate()"
        > arrow_forward_ios </i>
      </div>
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
    },

    selectedNewDate(direction = 1) {
      let selectedDate = this.selectedDate;
      selectedDate.setDate(selectedDate.getDate() + direction);

      this.selectedDate = new Date(selectedDate);
      this.$emit('selectedDate', this.selectedDate);
    },

    formatDate(date) {
      let options = {
        year: 'numeric',
        month: 'long',
        day: 'numeric',
      };

      return date.toLocaleString('ru', options);
    },
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
  justify-content: space-between;
}

.slice {
  font-family: "Exo 2", serif;
  font-weight: 700;
  font-size: 1.25rem;
  margin-right: 21px;
  color: #424242;
  background: none;
  border: none;
  cursor: pointer;
  transition-duration: 0.5s;
  padding: 0;
}

.design-line-element {
  border: 1px solid #000000;
}

.activate {
  color: #00B268;
  transition-duration: 0.5s;
}

.material-icons {
  outline: none;
}

.prev-date-btn,
.next-date-btn {
  transition: 0.5s;
}

.selected-date {
  font-family: "Exo 2", serif;
  font-weight: 700;
  font-size: 1.25rem;
  margin-right: 27px;
}

.today-btn {
  background: #FFFFFF;
  box-shadow: 0 4px 4px rgba(0, 0, 0, 0.25);
  border: 0;
  border-radius: 15px;
  height: 40px;
  width: 157px;
  font-family: "Exo 2", serif;
  font-weight: 700;
  font-size: 1.25rem;
  margin-right: 21px;
  cursor: pointer;
  transition: 0.5s;
}

.select-time-container  {
  justify-self: flex-end;
  align-self: center;
}

/* Animations and hovers */
.prev-date-btn:hover {
  color: #00B268;
  transition: 0.5s;
  cursor: pointer;
}

.next-date-btn:hover {
  color: #00B268;
  transition: 0.5s;
  cursor: pointer;
}

.today-btn:hover {
  background-color: lightgrey;
  transition: 0.2s;
}
</style>
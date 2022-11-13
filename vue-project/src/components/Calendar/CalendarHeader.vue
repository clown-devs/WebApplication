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

      <i class="material-icons prev-date-btn"
         @click="selectedNewDate(-1)"
      > arrow_back_ios </i>

      <i class="material-icons next-date-btn"
         @click="selectedNewDate()"
      > arrow_forward_ios </i>

      <label class="selected-date"> {{ this.formatDate(selectedDate) }} </label>
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
      const createdDate = new Date(selectedDate);
      console.log(createdDate);

      this.selectedDate = createdDate;
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

.next-date-btn {
  margin-right: 30px;
}

.selected-date {
  font-family: "Exo 2", serif;
  font-weight: 700;
  font-size: 1.25rem;
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

</style>
<template>
  <div class="small-calendar-container">
    <header class="header-small-calendar">
      <h3 class="header-date"> {{ this.month[currentMonth] }} {{ currentYear }} </h3>
      <i @click='changeMonth' class="material-icons icon left" id="prev"> arrow_left </i>
      <i @click='changeMonth' class="material-icons icon right" id="next"> arrow_right </i>
    </header>
    <hr>

    <main class="main-grid">
      <ul class="calendar">
        <li v-for="nameDay in daysName" :key="nameDay" class="day-name"> {{ nameDay }}</li>

        <li v-for="prevDay in prevDays" :key="prevDay" class="day inactive"> {{ prevDay }}</li>

        <li v-for="day in days" :key="day" class="day current-day" v-bind:class="{choseDay: getIsDay(day) }">
          <button class="day-button" @click="getIsDayButton" v-bind:class="{active: getIsToday(day)}"> {{ day }} </button>
        </li>

        <li v-for="nextDay in nextDays" :key="nextDay" class="day inactive"> {{ nextDay }}</li>
      </ul>
    </main>
  </div>
</template>

<script>
export default {
  props: {
    selectedDate: String
  },

  data() {
    return {
      date: "",
      currentYear: 0,
      currentMonth: 0,
      month: [
        "Январь", "Февраль", "Март", "Апрель", "Май", "Июнь", "Июль", "Август", "Сентябрь", "Октябрь", "Ноябрь", "Декабрь",
      ],
      days: [],
      daysForSelectedDay: [],
      daysName: ["П", "В", "С", "Ч", "П", "С", "В"],
      prevDays: [],
      nextDays: [],
      lastDateOfMonth: "",
      firstDayOfMonth: "",
      lastDateOfLastMonth: "",
      numberOfDay: "",
      isDay: false,
    };
  },

  mounted() {
    // assign a Date
    this.date = new Date();
    this.currentYear = this.date.getFullYear();
    this.currentMonth = this.date.getMonth();

    this.renderCalendar();
  },

  methods: {
    renderCalendar() {
      //получаем число день в недели в котором начинается месяц если 2 то со вторника и тд
      this.firstDayOfMonth = new Date(this.currentYear, this.currentMonth, 1).getDay();
      //Получаем количество дней в месяце
      this.lastDateOfMonth = new Date(this.currentYear, this.currentMonth + 1, 0).getDate();
      //Получаем сколько дней у этого месяца в последней недели
      this.lastDayOfMonth = new Date(this.currentYear, this.currentMonth, this.lastDateOfMonth).getDay();
      //Получаем количество дней в предыдущем месяце
      this.lastDateOfLastMonth = new Date(this.currentYear, this.currentMonth, 0).getDate();

      if (this.firstDayOfMonth === 0) {
        this.firstDayOfMonth = 7;
      }

      //Get in array end days prev Month
      for (let i = this.firstDayOfMonth - 1; i > 0; i--) {
        this.prevDays.push(this.lastDateOfLastMonth - i + 1);
      }
      //Get in array start days next Month
      for (let i = this.lastDayOfMonth; i < 7; ++i) {
        this.nextDays.push(i - this.lastDayOfMonth + 1);
      }
      //Get in array days current Month
      for (let i = 1; i <= this.lastDateOfMonth; ++i) {
        this.days.push(i);
        this.daysForSelectedDay.push(false);
      }
    },

    changeMonth(icon) {
      this.currentMonth = icon.target.id === "prev" ? this.currentMonth - 1 : this.currentMonth + 1;
      if (this.currentMonth < 0 || this.currentMonth > 11) {
        this.date = new Date(this.currentYear, this.currentMonth, this.date.getDate());
        this.currentYear = this.date.getFullYear();
        this.currentMonth = this.date.getMonth();
      }

      this.prevDays = [];
      this.days = [];
      this.nextDays = [];
      this.renderCalendar();
    },

    getIsToday(day) {
      if (day === this.date.getDate() && this.currentMonth === new Date().getMonth() && this.currentYear === new Date().getFullYear()) {
        return true;
      } else
        return false;
    },

    getIsDay(day) {
      if (this.daysForSelectedDay[day - 1] === true) {
        return true
      }
    },

    getIsDayButton(day) {
      for (let i = 1; i <= this.lastDateOfMonth; ++i) {
        this.daysForSelectedDay[i - 1] = false;
      }
      this.daysForSelectedDay[day.target.innerText - 1] = true;
    }

  }
}
</script>

<style scoped>
/* General style */
.small-calendar-container {
  margin-left: 500px;
  border: 1px solid #BDBDBD;
  background: #FFFFFF;
  border-radius: 30px;
  width: 300px;
}

/* Header style */

.header-small-calendar {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 17px;
}

.header-date {
  margin: 0 auto 0 40px;
  padding: 0;

}

.icon {
  height: 40px;
  width: 40px;
  text-align: center;
  line-height: 40px;
  font-size: 40px;
  cursor: pointer;
  margin: 0 1px;
  border-radius: 50%;
}


.left {
  margin-right: 5px;
  outline: none;
}

.right {
  margin-right: 40px;
  outline: none;
}

hr {
  border: 0;
  width: 85%;
  border-bottom: 1px solid #BDBDBD;
  margin: 0 auto 11px auto;
}

/* Main style */

.main-grid {
  width: 260px;
  padding-right: 40px;
}

.calendar {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 9px;
  margin: 0 0 25px 0;

}

.calendar li {
  position: relative;
  width: 24px;
  height: 24px;
  line-height: 24px;
}

.day {
  list-style: none;
  text-align: center;
  cursor: pointer;
  z-index: 1;
}

.day.inactive {
  color: #aaa;
}

.day.active {
  color: #fff;
}

.day-button {
  border: 0;
  background: none;
  padding: 0;
  vertical-align: center;
  align-items: center;
  cursor: pointer;
}

.choseDay {
  border-radius: 50%;
  background: #7de5b8;
}

.calendar .day-name {
  font-weight: 700;
  font-size: 14px;
  list-style: none;
  text-align: center;
}


/* Animation main */
.icon:hover {
  background: #f2f2f2;
}

.day-button::before {
  position: absolute;
  content: "";
  height: 22px;
  width: 22px;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 50%;
  z-index: -1;

}

.day-button::after {
  position: absolute;
  content: "";
  inset: 0;
  z-index: -1;
}

.day-button:hover::before:not(.choseDay) {
  background: #f2f2f2;
}

.day-button.active::before {
  background: #00B268;
}

.day-button.active {
  color: white;
}

.icon:active {
  color: #00B268;
}

.day-button:nth-child(7n), .day-button:nth-child(7n - 1) {
  color: #CC2222;
}
</style>
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
        <li class="day-name">П</li>
        <li class="day-name">В</li>
        <li class="day-name">С</li>
        <li class="day-name">Ч</li>
        <li class="day-name">П</li>
        <li class="day-name">С</li>
        <li class="day-name">В</li>
        <li v-for="prevDay in prevDays" :key="prevDay" class="day inactive"> {{ this.prevDays[prevDay - 1] }} </li>

        <li v-for="day in days" :key="day" class="day"> {{ this.days[day - 1] }} </li>
      </ul>
    </main>
  </div>
</template>

<script>
export default {

  data() {
    return {
      date: "",
      currentYear: "",
      currentMonth: "",
      month: [
        "Январь", "Февраль", "Март", "Апрель", "Май", "Июнь", "Июль", "Август", "Сентябрь", "Октябрь", "Ноябрь", "Декабрь",
      ],
      days: [],
      prevDays: [],
      lastDateOfMonth: "",
      firstDayOfMonth: "",
      lastDateOfLastMonth: "",
      numberOfDay: "",
    };
  },

  mounted() {
    // assign a Date
    this.date = new Date();
    this.currentYear = this.date.getFullYear();
    this.currentMonth = this.date.getMonth();

    this.ConsoleLog();
    this.renderCalendar();
  },

  methods: {
    ConsoleLog() {
    },

    renderCalendar() {
      this.firstDayOfMonth = new Date(this.currentYear, this.currentMonth, 1).getDay();
      this.lastDateOfMonth = new Date(this.currentYear, Number(this.currentMonth) + 1, 0).getDate();
      this.lastDateOfLastMonth = new Date(this.currentYear, this.currentMonth, 0).getDate();

      for (let i = this.firstDayOfMonth; i > 1; i--) {
        this.prevDays.push(this.lastDateOfLastMonth - i + 1);
      }

      console.log(this.prevDays )

      for (let i = 1; i <= this.lastDateOfMonth; ++i) {
        this.days.push(i);
      }
    },

    changeMonth(icon) {
      this.currentMonth = icon.target.id === "prev" ? this.currentMonth - 1 : this.currentMonth + 1;
      this.prevDays = [];
      this.days = [];
      this.renderCalendar();
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
  height: 39px;
  width: 39px;
  text-align: center;
  line-height: 39px;
  font-size: 39px;
  cursor: pointer;
  margin: 0 1px;
  border-radius: 50%;
  color: #878787;
}

.left {
  margin-right: 5px;
}

.right {
  margin-right: 40px;
}

hr {
  border: 0;
  width: 85%;
  border-bottom: 1px solid #BDBDBD;
  margin: 11px auto 11px auto;
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


.calendar .day-name {
  font-weight: 700;
  font-size: 14px;
  list-style: none;
  text-align: center;
}

.grid {
  display: grid;
  margin-left: 500px;
  grid-template-columns: 100px 100px;
  grid-template-rows: 100px 100px;

  width: 200px;
  height: 200px;

  gap: 10px;
}

.header-grid {
  border: 1px solid green;
}

.header-grid {
  grid-row: 1 / -1;
}

/* Animation main */
.icon:hover {
  background: #f2f2f2;
}

.day::before {
  position: absolute;
  content: "";
  height: 23px;
  width: 23px;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 50%;
  z-index: -1;
}

.day:hover::before {
  background: #f2f2f2;
}

.day.active::before {
  background: #00B268;
}

.icon:active {
  color: #00B268;
}
</style>
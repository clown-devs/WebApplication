<template>
  <div class="small-calendar-container">
    <header class="header-small-calendar">
      <h2 class="header-date">Октябрь 2022</h2>
      <i class="material-icons icon left"> arrow_left </i>
      <i class="material-icons icon right"> arrow_right </i>
    </header>
    <hr>

    <main class="main-grid">
      <ul class="calendar">
        <li class="day day-name">П</li>
        <li class="day day-name">В</li>
        <li class="day day-name">С</li>
        <li class="day day-name">Ч</li>
        <li class="day day-name">П</li>
        <li class="day day-name">С</li>
        <li class="day day-name">В</li>

        <li class="day">1</li>
        <li class="day">2</li>
        <li class="day">3</li>
        <li class="day">4</li>
        <li class="day">5</li>
        <li class="day">6</li>
        <li class="day">7</li>
        <li class="day">8</li>
        <li class="day">9</li>
        <li class="day">10</li>
        <li class="day">11</li>
        <li class="day">12</li>
        <li class="day">13</li>
        <li class="day">14</li>
        <li class="day">15</li>
        <li class="day">16</li>
        <li class="day">17</li>
        <li class="day">18</li>
        <li class="day">19</li>
        <li class="day">20</li>
        <li class="day">21</li>
        <li class="day">22</li>
        <li class="day">23</li>
        <li class="day">24</li>
        <li class="day">25</li>
        <li class="day">26</li>
        <li class="day">27</li>
        <li class="day">28</li>
        <li class="day">29</li>
        <li class="day">30</li>
        <li class="day">31</li>
      </ul>
    </main>
  </div>
</template>

<script>
export default {}
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
  margin: 0 10px 0 0;
  padding: 0;
}

.icon {
  height: 33px;
  width: 33px;
  text-align: center;
  line-height: 33px;
  font-size: 33px;
  cursor: pointer;
  margin: 0 1px;
  border-radius: 50%;
  color: #878787;
}

.left {
  margin-right: 5px;
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
}

.calendar {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 9px;
  margin: 0 0 25px 0;
}

.day {
  list-style: none;
  text-align: center;
}

.calendar .day-name {
  font-width: 700px;
  font-size: 14px;
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
</style>
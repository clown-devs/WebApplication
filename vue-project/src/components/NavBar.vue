<template>
  <header class="header">
    <nav>
      <div class="title-name title-names">
        <h3 class="logo-name">Sber clients</h3>
      </div>

        <ul class="nav-bar-list">
        <li class="nav-bar-list-item" @click="$router.push('/')">
          Главная страница
        </li>

        <li class="nav-bar-list-item" @click="$router.push('/clients')">
          Мои клиенты
        </li>

        <li class="nav-bar-list-item" @click="$router.push('/calendar')">
          Календарь
        </li>

        <li class="nav-bar-list-item" @click="$router.push('/results')">
          Отчёты по встречам
        </li>
        </ul>

      <button class="nav-bar-notify-button">
        <img src="/svg/notify.svg" alt="" class="nav-bar-notify-icon" />
      </button>

      <div class="nav-bar-icon-button">
        <button class="nav-bar-add-meet">
          <img src="/svg/add-meet.svg" alt="" class="nav-bar-icon-add" />
        </button>

        <button class="nav-bar-profile">
          <img src="/svg/profile.svg" alt="" class="nav-bar-profile-icon" />
        </button>
        <div class="burger">
          <div class="line line1"></div>
          <div class="line line2"></div>
          <div class="line line3"></div>
      </div>
        <button @click="logOutWrapper" class="log-out">Выйти</button>
    </div>
    </nav>

    <hr class="line-header" />
  </header>

  <!-- <main>
    <div class="container">
      <div class="content-text">
        <p>
          Lorem ipsum dolor sit amet consectetur, adipisicing elit. Perferendis
          iusto cum voluptatum veniam itaque, earum, veritatis reiciendis
          distinctio eaque unde magni illo adipisci consequuntur, dignissimos
          vel alias facere ratione nihil!
        </p>
        <p>
          Lorem ipsum dolor sit amet consectetur, adipisicing elit. Perferendis
          iusto cum voluptatum veniam itaque, earum, veritatis reiciendis
          distinctio eaque unde magni illo adipisci consequuntur, dignissimos
          vel alias facere ratione nihil!
        </p>
        <p>
          Lorem ipsum dolor sit amet consectetur, adipisicing elit. Perferendis
          iusto cum voluptatum veniam itaque, earum, veritatis reiciendis
          distinctio eaque unde magni illo adipisci consequuntur, dignissimos
          vel alias facere ratione nihil!
        </p>
        <p>
          Lorem ipsum dolor sit amet consectetur, adipisicing elit. Perferendis
          iusto cum voluptatum veniam itaque, earum, veritatis reiciendis
          distinctio eaque unde magni illo adipisci consequuntur, dignissimos
          vel alias facere ratione nihil!
        </p>
      </div>
    </div>
  </main> -->
</template>

<script>
import { mapMutations } from "vuex";

export default {
  components: {},
  data: () => ({
    width: 0,
  }),
  methods: {
    ...mapMutations(["logOut"]),

    logOutWrapper() {
      this.$router.push("/login");
      this.logOut();
    },
    slider() {
      const burger = document.querySelector(".burger"),
        nav = document.querySelector(".nav-bar-list"),
        links = document.querySelector(".nav-bar-list li");

      burger.addEventListener("click", () => {
        nav.classList.toggle("nav-active");

        burger.classList.toggle("toggle");

        links.forEach((link, index) => {
          if (link.style.animation) {
            link.style.animation = "";
          } else {
            link.style.animation = "navLinksFade .5s ease ";
          }
        });
      });
    },
    updateWidth() {
      if ((this.width = window.innerWidth) <= 768) {
        slider();
      }
    },
  },
  created() {
    window.addEventListener("resize", this.updateWidth);
  },
};
</script>

<style scoped>
button {
  background-color: white;
  border: none;
  margin: 0;
  padding: 0;
}

.header {
  background-color: #fff;
}

* {
  padding: 0;
  margin: 0;

  box-sizing: border-box;
}

html,
body {
  overflow-x: hidden;
}

nav {
  display: flex;
  justify-content: space-around;
  align-items: center;
  min-height: 110px;
  margin-left: 200px;
  margin-right: 200px;
}

.nav-bar-list {
  display: flex;
  justify-content: space-around;
  width: 50%;
}

.nav-bar-list li {
  list-style: none;
}

.nav-bar-icon-button {
  display: flex;
  justify-content: space-around;
  width: 20%;
}

.burger {
  cursor: pointer;
  display: none;
  align-self: center;
  }

.line {
  width: 25px;
  height: 3px;
  margin: 5px;
  background-color: #47af52;
  transition: all 0.4s ease;
  -webkit-transition: all 0.4s ease;
  -moz-transition: all 0.4s ease;
  -ms-transition: all 0.4s ease;
  -o-transition: all 0.4s ease;
}

@media screen and (max-width: 1024px) {
  .nav-bar-list {
    width: 80%;
  }
  .nav-bar-icon-button {
    width: 35%;
  }
  nav {
    display: flex;
    justify-content: space-around;
    align-items: center;
    min-height: 110px;
    margin-left: 100px;
    margin-right: 100px;
  }
  }

@media screen and (max-width: 768px) {
  .burger {
    display: block;
}

  .nav-bar-icon-button {
    width: 55%;
  }
  nav {
    display: flex;
    justify-content: space-around;
    align-items: center;
    min-height: 110px;
    margin: 0;
}

.nav-bar-list {
    width: 10%;

    position: absolute;
    right: 0;
    height: 93vh;
    top: 7vh;
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 50%;
    /*background-color: #bceec0;
    transform: translateX(100%);
    -webkit-transform: translateX(100%);
    -moz-transform: translateX(100%);
    -ms-transform: translateX(100%);
    -o-transform: translateX(100%);
    transition: transform 0.5s ease-in;
    -webkit-transition: transform 0.5s ease-in;
    -moz-transition: transform 0.5s ease-in;
    -ms-transition: transform 0.5s ease-in;
    -o-transition: transform 0.5s ease-in;*/
  }

  .nav-bar-list li {
    opacity: 0;
  }
}

.nav-active {
  transform: translateX(0%);
  -webkit-transform: translateX(0%);
  -moz-transform: translateX(0%);
  -ms-transform: translateX(0%);
  -o-transform: translateX(0%);
}

@keyframes navLinksFade {
  from {
    opacity: 0;
    transform: translateX(50px);
    -webkit-transform: translateX(50px);
    -moz-transform: translateX(50px);
    -ms-transform: translateX(50px);
    -o-transform: translateX(50px);
  }

  to {
    opacity: 1;
    transform: translateX(0px);
    -webkit-transform: translateX(0px);
    -moz-transform: translateX(0px);
    -ms-transform: translateX(0px);
    -o-transform: translateX(0px);
  }
}

.toggle .line1 {
  transform: rotate(-45deg) translate(-5px, -6px);
  -webkit-transform: rotate(-45deg) translate(-5px, -6px);
  -moz-transform: rotate(-45deg) translate(-5px, -6px);
  -ms-transform: rotate(-45deg) translate(-5px, -6px);
  -o-transform: rotate(-45deg) translate(-5px, -6px);
}

.toggle .line2 {
  opacity: 0;
}

.toggle .line3 {
  transform: rotate(45deg) translate(-5px, -6px);
  -webkit-transform: rotate(45deg) translate(-5px, -6px);
  -moz-transform: rotate(45deg) translate(-5px, -6px);
  -ms-transform: rotate(45deg) translate(-5px, -6px);
  -o-transform: rotate(45deg) translate(-5px, -6px);
}
/*.nav-bar-list {
  display: flex;
  justify-content: space-around;
}

.nav-bar-icon-button {
  display: flex;
  margin-left: auto;
  justify-content: center;
  align-items: center;
}

.nav-bar-icon-button-two {
  margin-left: auto;
}

.title-name {
  margin-right: 60px;
}

ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.line-header {
  width: 100%;
}

.nav-bar-list-item {
  font-size: 16px;
  margin-right: 52px;
}

.nav-bar-add-meet {
  margin-right: 34px;
}

.log-out {
  background-color: white;
  border: none;
  margin: 0;
  padding: 0;
}
.nav-bar-notify-button {
  height: 40px;

  background-color: white;
  border: none;
  margin: 0;
  padding: 0;
}

.nav-bar-add-meet {
  height: 60px;

  background-color: white;
  border: none;
  margin-top: 0;
  margin-bottom: 0;
  margin-left: 0;
  padding: 0;
  display: flex;
}

.nav-bar-profile {
  height: 70px;

  background-color: white;
  border: none;
  margin-top: 0;
  margin-bottom: 0;
  margin-left: 0;
  padding: 0;
  margin-right: 20px;
}*/

@media (max-width: 1520px) {
  .container {
    max-width: 970px;
  }
}

@media (max-width: 992px) {
  .container {
    max-width: 750px;
  }
}

@media (max-width: 767px) {
  .container {
    max-width: none;
  }
}
</style>
<template>
  <header class="header">
    <nav>
      <div class="title-name title-names">
        <h3 class="logo-name">Sber clients</h3>
      </div>
      <ul class="nav-bar-list">
        <input type="checkbox" name="" id="checkbox_toggle" />
        <label for="checkbox_toggle" class="hamburger"
          ><span class="hamburger-line"></span
        ></label>

        <div class="menu-items">
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

          <li class="nav-bar-list-item button-exit" @click="logOutWrapper">
            <button class="log-out-item">Выйти</button>
          </li>
        </div>
      </ul>

      <button class="nav-bar-notify-button">
        <img src="/svg/notify.svg" alt="" class="nav-bar-notify-icon" />
      </button>

      <button class="nav-bar-add-meet">
        <img src="/svg/add-meet.svg" alt="" class="nav-bar-icon-add" />
      </button>

      <button class="nav-bar-profile">
        <img src="/svg/profile.svg" alt="" class="nav-bar-profile-icon" />
      </button>

      <button @click="logOutWrapper" class="log-out">Выйти</button>
    </nav>

    <hr class="line-header" />
  </header>
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

.menu-items {
  display: flex;
  justify-content: space-around;
  align-items: center;
}

html,
body {
  overflow-x: hidden;
}
.button-exit {
  display: none;
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

.nav-bar-list {
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

.log-out-item {
  background-color: #90dc97;
  border: none;
  font-size: 16px;
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
}

.nav-bar-icon-button {
  display: flex;
  justify-content: space-around;
  width: 20%;
}

#checkbox_toggle {
  display: none;
}

.hamburger {
  display: none;
  font-size: 50px;
  cursor: pointer;
}

.hamburger-line {
  background-color: #47af52;
  display: block;
  height: 2px;
  position: relative;
  width: 24px;
}

.hamburger-line::before,
.hamburger-line::after {
  background-color: #47af52;
  content: "";
  display: block;
  position: absolute;
  transition: all 0.3s;
  width: 100%;
  height: 100%;
}

.hamburger-line::before {
  top: 8px;
}

.hamburger-line::after {
  top: -8px;
}

@media (max-width: 1200px) {
  nav {
    margin: 0;
    display: flex;
    flex-direction: row-reverse;
  }

  nav .nav-bar-notify-button {
    margin-right: auto;
  }
  .button-exit {
    display: block;
  }

  .log-out-item {
    font-size: 24px;
  }

  .menu-items li {
    transition: all 0.5s;
    display: block;
    padding: 25px;
    font-size: 24px;
    margin-right: 0;
    border: 1px solid white;
  }
  .menu-items li:hover {
    background-color: #46d554;
  }

  .title-name,
  .log-out {
    display: none;
  }

  .menu-items {
    display: none;
    max-height: 0;
    overflow: hidden;
    position: absolute;
    background-color: #90dc97;
    text-align: center;
    right: 0;
    left: 0;
    margin-top: 77px;
  }

  .hamburger {
    display: block;
    padding: 20px 0;
    margin-right: 10px;
    margin-left: auto;
  }
  #checkbox_toggle:checked ~ .menu-items {
    display: block;
    max-height: 100%;
    transition: all 0.5s;
  }

  #checkbox_toggle:checked ~ .hamburger .hamburger-line {
    background-color: transparent;
  }

  #checkbox_toggle:checked ~ .hamburger .hamburger-line::before {
    transform: rotate(-45deg);
    top: 0;
  }
  #checkbox_toggle:checked ~ .hamburger .hamburger-line::after {
    transform: rotate(45deg);
    top: 0;
  }
}
</style>
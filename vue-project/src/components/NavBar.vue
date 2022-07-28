<template>
  <header class="header">
    <nav>
      <div class="container header-container">

        <div class="title-name title-names">
          <h3 class="logo-name">Sber clients</h3>
        </div>

        <div class="header-burger">
          <span></span>
        </div>

        <ul class="nav-bar-list">
          <li class="nav-bar-list-item" @click="$router.push('/')">Главная страница</li>

          <li class="nav-bar-list-item" @click="$router.push('/clients')">Мои клиенты</li>

          <li class="nav-bar-list-item" @click="$router.push('/calendar')">Календарь</li>

          <li class="nav-bar-list-item" @click="$router.push('/results')">Отчёты по встречам</li>
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
        </div>

        <div>
          <button @click="logOutWrapper" class="log-out">Выйти</button>
        </div>
      </div>
    </nav>

    <hr class="line-header" />
  </header>

</template>

<script>
import { mapMutations } from "vuex";

export default {
  components: {},

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
.container {
  display: flex;
  align-items: center;
}

.header {
  background-color: #fff;
}

* {
  padding: 0;
  margin: 0;

  box-sizing: border-box;
}

html, body {
  overflow-x: hidden;
}

nav {
  padding: 15px 10px;
  display: flex;
  justify-content: space-around;
  align-items: center;
  min-height: 110px;
  
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
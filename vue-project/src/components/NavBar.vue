<template>
  <header class="header">
    <nav>
      <div class="title-name title-names">
        <h3 class="logo-name" @click="$router.push('/')">Sber clients</h3>
      </div>
      <ul class="nav-bar-list">
        <input type="checkbox" name="" id="checkbox_toggle" />
        <label for="checkbox_toggle" class="hamburger">
          <span class="hamburger-line"></span>
        </label>

        <div class="menu-items">


          <li class="nav-bar-list-item" @click="$router.push('/clients')">
            <router-link to="/clients" class="inactive" active-class="active"
              >Мои клиенты</router-link
            >
          </li>

          <li class="nav-bar-list-item" @click="$router.push('/calendar')">
            <router-link to="/calendar" class="inactive" active-class="active"
              >Календарь</router-link
            >
          </li>

          <li class="nav-bar-list-item" @click="$router.push('/results')">
            <router-link to="/results" class="  inactive" active-class="active"
              >Отчёты по встречам</router-link
            >
          </li>

          <li class="nav-bar-list-item button-exit" @click="logOutWrapper">
            <button class="log-out-item">Выйти</button>
          </li>
        </div>
      </ul>
      <button class="nav-bar-notify-button">
        <img src="/svg/notify.svg" alt="" class="nav-bar-notify-icon" />
      </button>
      <div class="popup_wrapper">
        <v-popup
          v-if="isInfoPopupVisible"
          @closePopup="closeInfoPopup"
          class="modal-window"
        >
          <template v-slot:header>
            <h1 class="popup-title">Доступные комнаты</h1>
          </template>

          <template v-slot:content>
            <select
              name=""
              id=""
              v-if="isLoadedPlaces"
              class="select-room"
              v-model="selected"
            >
              <option v-for="place in places" :key="place">
                {{ place.name }}
              </option>
            </select>
            <loading-indicate v-else></loading-indicate>
          </template>
        </v-popup>
      </div>

      <button
        class="nav-bar-add-meet v-modal_add-meet"
        id="btn-add"
        @click="showPopupInfo"
      >
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
import vPopup from "@/components/UI/Popup.vue";
import LoadingIndicate from "@/components/UI/LoadingIndicate.vue";
import api from "@/api";

export default {
  components: { vPopup, LoadingIndicate },

  methods: {
    ...mapMutations(["logOut"]),

    logOutWrapper() {
      this.$router.push("/login");
      this.logOut();
    },

    showPopupInfo() {
      this.isInfoPopupVisible = true;
    },
    closeInfoPopup() {
      this.isInfoPopupVisible = false;
    },
  },

  data() {
    return {
      isInfoPopupVisible: false,
      isLoadedPlaces: false,
      places: [],
      selected: "test",
    };
  },
  
  async mounted() {
    this.places = await api.getPlaces();
    this.isLoadedPlaces = true;
  },
};
</script>

<style scoped>
@import url("../assests/fonts/Exo_2/stylesheet.css");

button {
  background-color: white;
  border: none;
  margin: 0;
  padding: 0;
}

.header {
  background-color: #fff;
  position: fixed;
  width: 100%;
  top: 0;
  left: 0;
  z-index: 50;
}

* {
  padding: 0;
  margin: 0;
  font-family: "Exo 2", serif;
  box-sizing: border-box;
  font-weight: 700;
}

.popup-title {
  display: flex;
  align-items: center;
  justify-content: center;
  flex: 2;
}

.select-room {
  display: flex;
  width: 100%;
  margin: 1rem;
}

.inactive {
  text-decoration: none;
  color: black;
}

.log-out {
  font-size: 1rem;
}

.logo-name {
  font-weight: 900;
  font-size: 1.5rem;
  cursor: pointer;
}

.select-room {
  width: 250px;
}

.menu-items {
  display: flex;
  justify-content: space-between;
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
  justify-content: space-between;
  align-items: center;
  min-height: 110px;
  margin-left: 10.14%;
  flex: 2;
  margin-right: 10.14%;
}

.nav-bar-list {
  display: flex;
  justify-content: space-between;
  width: 50%;
  cursor: pointer;
}

.nav-bar-list li {
  list-style: none;
}

.title-name {
  margin-right: 10px;
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
  cursor: pointer;
}

.nav-bar-add-meet {
  margin-right: 34px;
  cursor: pointer;
}

.log-out {
  background-color: white;
  border: none;
  margin: 0;
  padding: 0;
  cursor: pointer;
}

.log-out-item {
  background-color: #90dc97;
  border: none;
  font-size: 16px;
  margin: 0;
  padding: 0;
  cursor: pointer;
}

.nav-bar-notify-button {
  height: 40px;
  cursor: pointer;
  background-color: white;
  border: none;
  padding: 0;
}

.nav-bar-add-meet {
  height: 55px;
  background-color: white;
  border: none;
  padding: 0;
  display: flex;
  cursor: pointer;
  margin: 0 34px 0 0;
}

.nav-bar-profile {
  height: 70px;
  background-color: white;
  border: none;
  padding: 0;
  margin: 0 20px 0 0;
  cursor: pointer;
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
  height: 4px;
  position: relative;
  width: 35px;
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
  top: 9px;
}

.hamburger-line::after {
  top: -9px;
}

@media (max-width: 1200px) {
  nav {
    display: flex;
    min-height: 90px;
    justify-content: center;
    margin: 0 40px 0 40px;
  }
  nav .nav-bar-notify-button {
    margin-right: 20px;
  }
  .button-exit {
    display: block;
  }

  .nav-bar-profile {
    margin-right: 20px;
  }
  .log-out-item {
    display: none;
  }

  .nav-bar-add-meet {
    margin-right: 20px;
  }

  .title-name {
    display: none;
  }
  .nav-bar-list {
    margin-right: auto;
  }
}

@media (max-width: 992px) {
  nav {
    display: flex;
    min-height: 90px;
    justify-content: center;
    margin: 0 10px 0 10px;
  }
  nav .nav-bar-notify-button {
    margin-left: auto;
    margin-right: 20px;
  }
  .button-exit {
    display: block;
  }

  .nav-bar-profile {
    margin-right: 20px;
  }
  .log-out-item {
    display: none;
  }

  .nav-bar-add-meet {
    margin-right: 20px;
  }

  .title-name {
    display: none;
  }
}

@media (max-width: 768px) {
  nav {
    display: flex;
    min-height: 70px;
    align-items: center;
    justify-content: center;
    margin: 0 10px 0 10px;
  }
  nav .nav-bar-notify-button {
    margin-left: auto;
    margin-right: 20px;
    height: 30px;
  }
  .button-exit {
    display: block;
  }
  .nav-bar-profile {
    margin-right: 0;
    height: 55px;
  }

  .nav-bar-profile-icon {
    height: 55px;
  }

  .nav-bar-icon-add {
    height: 45px;
  }

  .nav-bar-notify-icon {
    height: 30px;
  }
  .log-out-item {
    font-size: 24px;
    display: flex;
  }
  .nav-bar-add-meet {
    margin-right: 20px;
    height: 45px;
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
    display: flex;
    max-height: 0;
    overflow: hidden;
    position: absolute;
    background-color: #90dc97;

    right: 0;
    left: 0;
    margin-top: 57px;
  }
  .hamburger {
    display: block;
    padding: 20px 0;
    margin-right: auto;
  }
  #checkbox_toggle:checked ~ .menu-items {
    display: block;
    max-height: 400px;
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

/* Hovers and animations */

.logo-name:hover,
.log-out:hover {
  opacity: 0.9;
  transition: 0.5s;
  color: #00b268;
}
.inactive:hover {
  opacity: 0.9;
  transition: 0.5s;
  color: #00b268;
}
.nav-bar-list-item:hover {
  opacity: 0.9;
  transition: 0.5s;
  color: #00b268;
}

.nav-bar-notify-button:hover,
.nav-bar-icon-add:hover {
  opacity: 0.7;
  transition: 0.5s;
}
</style>


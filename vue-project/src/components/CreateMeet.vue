<template>
  <div class="popup_wrapper" ref="popup_wrapper">
    <form class="v-popup animation" @submit.prevent>
      <div class="v-popup__header">
        <slot name="header"></slot>
        <span>
          <i class="material-icons" @click="closePopup"> close </i>
        </span>
      </div>

      <div class="v-popup__content">
        <div class="info-meeting">
          <div class="info-meet-main">
            <div class="client-container">
              <input type="text" class="client-input" placeholder="Клиент" />
            </div>

            <div class="contact-container">
              <input
                type="text"
                class="contact-input"
                placeholder="Контактное лицо"
              />
            </div>

            <div class="theme-container">
              <input type="text" class="theme-input" placeholder="Тема" />
            </div>

            <div class="date-icon-all">
              <div class="date-container">
                <input type="date" class="date-input" placeholder="Дата" />
              </div>

              <!-- <div class="icon-date">
                <button class="date-button">
                  <img src="/svg/date.svg" alt="" class="date-icon" />
                </button>
              </div> -->
            </div>

            <div class="time-start-end">
              <div class="time-start">
                <input
                  type="time"
                  class="date-start-input"
                  placeholder="Время начала"
                />
              </div>

              <div class="time-end">
                <input
                  type="time"
                  class="date-end-input"
                  placeholder="Время окончания"
                />
              </div>
            </div>

            <div class="place-container">
              <input
                type="text"
                class="place-input"
                placeholder="Место встречи"
              />
            </div>
          </div>

          <div class="info-worker">
            <div class="worker-container">
              <input
                type="text"
                class="worker-input"
                placeholder="Сотрудники"
              />
            </div>
          </div>
        </div>

        <div class="info-note">
          <div class="note-container">
            <input type="text" class="note-input" placeholder="Заметки" />
          </div>
        </div>
      </div>

      <div class="v-popup__footer">
        <slot name="footer"></slot>
      </div>
    </form>
  </div>
</template>

<script>
export default {
  props: {},

  data() {
    return {};
  },

  methods: {
    closePopup() {
      this.$emit("closePopup");
    },
  },

  mounted() {
    let vm = this;
    document.addEventListener("click", function (item) {
      if (item.target === vm.$refs["popup_wrapper"]) {
        vm.closePopup();
      }
    });
  },
};
</script>

<style scoped lang="scss">
.material-icons {
  cursor: pointer;
}

.popup_wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  position: fixed;
  right: 0;
  left: 0;
  top: 0;
  background: rgba(0, 0, 0, 0.5);
  bottom: 0;
  z-index: 100;
}

.info-meeting, .time-start-end{
  display: flex;
}


.worker-input {
  width: 350px;
  padding-bottom: 200px;
  background: #f5f5f5;
  border-bottom: 1px solid #7a7474;
  border-radius: 10px;
}

.info-worker {
  margin-left: auto;
}

.client-input,
.contact-input,
.place-input {
  background: #f5f5f5;
  border-bottom: 1px solid #7a7474;
  border-radius: 7px;
  padding-right: 56%;
  height: 20px;
  width: 400px;
  margin-bottom: 10px;
}

.theme-input {
  background: #f5f5f5;
  border-bottom: 1px solid #7a7474;
  border-radius: 7px;
  padding-right: 56%;
  height: 30px;
  width: 400px;
  margin-bottom: 10px;
}

.date-input {
    background: #f5f5f5;
  border-bottom: 1px solid #7a7474;
  border-radius: 7px;
  padding-right: 56%;
  height: 20px;
  width: 400px;
  margin-bottom: 10px;
}

.date-start-input {
    background: #f5f5f5;
  border-bottom: 1px solid #7a7474;
  border-radius: 7px;
  height: 20px;
  margin-bottom: 10px;
  margin-right: 160px;
  width: 100px;
}

.note-input {
    background: #f5f5f5;
  border-bottom: 1px solid #7a7474;
  border-radius: 7px;
  padding-right: 56%;
  height: 60px;
  width: 448px;
  margin-bottom: 10px;
}

.date-end-input {
    background: #f5f5f5;
  border-bottom: 1px solid #7a7474;
  border-radius: 7px;
  height: 20px;
  margin-bottom: 10px;
  margin-left: auto;
}
.v-popup {
  border-radius: 25px;
  position: fixed;
  background: #ffffff;
  box-shadow: 0 0 17px 0 #e7e7e7;
  z-index: 10;
  width: 74%;

  &__header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin: 2rem;
    font-family: "Exo 2";
    font-weight: 700;
    font-size: 1.5rem;
  }

  &__footer {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    margin: 2rem;
  }

  &__content {
    display: flex;
    flex-direction: column;
    gap: 2rem;
    margin: 2rem;
  }

  .close_modal {
    color: #ffffff;
    background: #00b268;
  }
}

.animation {
  animation-duration: 0.55s;
  animation-fill-mode: both;
  animation-name: fadeInDown;
}

@keyframes fadeInDown {
  0% {
    opacity: 0;
    transform: translate3d(0px, -100%, 0px);
  }
  100% {
    opacity: 1;
    transform: none;
  }
}

@media (max-width: 1200px) {
  .v-popup {
    top: 150px;
  }
}

@media (max-width: 768px) {
  .v-popup {
    top: 0;
    width: 100vw;
    height: 100vh;
    border-radius: 0;
  }
}

/* Animations and hovers */

.material-icons {
  transition-duration: 0.5s;
}

.material-icons:hover {
  transition-duration: 0.5s;
  color: #00b268;
}
</style>
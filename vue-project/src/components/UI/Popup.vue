<template>
  <div class="popup_wrapper" ref="popup_wrapper">
    <div class="v-popup animation">
      <div class="v-popup__header">
        <span>Доступные переговорки</span>
        <span>
          <i class="material-icons" @click="closePopup"> close </i>
        </span>
      </div>
      <div class="v-popup__content">
        <slot></slot>
      </div>
    </div>
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
    document.addEventListener('click', function (item) {
        if (item.targer === vm.$refs['popup_wrapper']) {
            vm.closePopup()
        }
    })
  },

};
</script>

<style scoped lang="scss">
.popup_wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  position: fixed;
  right: 0;
  left: 0;
  top: 0;
  

  background: rgba(70, 70, 70, 0.4);    
  bottom: 0;
}

.v-popup {
  border-radius: 25px;
  padding: 16px;
  position: fixed;
  top: 50px;
  width: 450px;
  height: 500px;
  background: #ffffff;
  box-shadow: 0 0 17px 0 #e7e7e7;
  z-index: 10;
  &__header,
  &__footer {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  &__content {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: 70px;
  }
  .close_modal {
    padding: 8px;
    color: #ffffff;
    background: #00B268;
    margin-top: 285px;
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
  nav {
    margin: 0;
    display: flex;
    margin-left: 10px;
    margin-right: 10px;
    display: flex;
    justify-content: center;
  }
  nav .nav-bar-notify-button {
    margin-left: auto;
    margin-right: 20px;
    border-radius: 25px;
  }
  .button-exit {
    display: block;
  }

  .nav-bar-profile {
    margin-right: 0;
  }

  .v-popup {
  padding: 16px;
  position: fixed;
  top: 50px;
  width: 250px;
  height: 500px;
  background: #ffffff;
  box-shadow: 0 0 17px 0 #e7e7e7;
  z-index: 10;
  &__header,
  &__footer {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  &__content {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: 70px;
  }
  .close_modal {
    padding: 8px;
    color: #ffffff;
    background: #00B268;
    margin-top: 285px;
  }

}

  .log-out-item {
    font-size: 24px;
  }

  .nav-bar-add-meet {
    margin-right: 20px;
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
    margin-right: auto;
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
<template>
  <div class="popup_wrapper" ref="popup_wrapper">
    <div class="v-popup animation">
      <div class="v-popup__header">
        <slot name="header"></slot>
        <span>
          <i class="material-icons" @click="closePopup"> close </i>
        </span>
      </div>
      <div class="v-popup__content">
        <slot name="content"></slot>
      </div>
      <div class="v-popup__footer">
        <slot name="footer"></slot>
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
  background: rgba(70, 70, 70, 0.4);
  bottom: 0;
}

.v-popup {
  border-radius: 25px;
  position: fixed;
  top: 200px;
  background: #ffffff;
  box-shadow: 0 0 17px 0 #e7e7e7;
  z-index: 10;
  width: 500px;
  height: 500px;

  &__header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin: 1rem;
  }

  &__footer {
    display: flex;
    justify-content: center;
    align-items: center;
  }

  &__content {
    display: flex;
    justify-content: center;
    align-items: center;
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
<template>
  <div class="show-info">
    <button class="show-info-content">{{ buttonTitle }}</button>
  </div>
</template>

<script>
export default {
  props: {
    showInfo: {
        type: Boolean,
        default: false
    }
  },

  computed: {
    buttonTitle: function() {
        return !this.showInfo ? "Подробнее" : "Скрыть";
    }  
  }
};
</script>

<style scoped>
.show-info {
  display: flex;
  align-items: center;
  justify-content: flex-end;
}

.show-info-content {
  cursor: pointer;
  background-color: #f1fbf2;
  border-width: 0;
  font-weight: 700;
  font-size: 1.25rem;
  font-family: "Exo 2";
  color: #7a7474;
}

/* Animations and hovers */
.show-info-content {
  transition-duration: 0.5s;
}

.show-info-content:hover {
  transition-duration: 0.5s;
  color: #00b268;
}
</style>
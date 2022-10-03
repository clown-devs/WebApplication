<template>
  <div class="direction-user-item" v-if="this.directionUser !== undefined">
    <div class="direction-user__header">
      <div class="direction-user-name">
        <span class="direction-user-name-content">
          {{ directionUser.first_name }} {{ directionUser.second_name }}
          {{ directionUser.third_name }}
        </span>
      </div>
    </div>
  </div>
  <div class="direction-user empty" v-else>
    <span>Нет сотрудников</span>
  </div>
</template>

<script>
export default {
  data() {
    return {
      directionUser: this.directionUserData,
    };
  },

  props: {
    directionUserData: {
      type: Object,
      default: undefined,
    },
  },
};
</script>

<style scoped>

/* Main styles */
.direction-user-item {
  min-height: 70px;
  border: 0.72px solid #f1fbf2;
  box-shadow: 0px 2.86px 2.86px rgba(0, 0, 0, 0.25);
  border-radius: 10px;
  background-color: #fff;
  margin: 10px 0 10px 0;
  margin-left: 11.5%;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.empty {
  align-items: center;
  font-family: "Exo 2";
  font-weight: 700;
}

.direction-user__header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 70px;
  width: 100%;
}

.direction-user-name-content {
  margin-left: 20px;
  font-family: "Exo 2";
  font-weight: 700;
  font-size: 1.25rem;
  cursor: default;
}

/* Animations and hovers */
.direction-user-item {
    transition-duration: 0.5s;
}

.direction-user-item:hover {
  transition-duration: 0.5s;
  background: #e5e4e4;
}
</style>
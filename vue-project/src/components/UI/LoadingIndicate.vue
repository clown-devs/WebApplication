<template>
  <div class="loading-container">
    <div class="loading-indicator"></div>
    <p class="information-text">Загрузка данных...</p>
  </div>
</template>

<script>
export default {};
</script>

<style scoped>

.loading-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  margin-top: 2rem;
}

.loading-indicator {
  width: 80px;
  height: 80px;
  border: 8px solid;
  border-radius: 50%;
  border-color: green #e2eee3;
  animation: animate 2s linear infinite;
}

@keyframes animate {
  to {
    transform: rotate(360deg);
  }
}
</style>
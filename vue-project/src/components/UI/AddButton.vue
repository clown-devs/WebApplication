<template>
  <button class="add-meet-but">
    <slot></slot>
  </button>
</template>

<script>
export default {};
</script>

<style scoped>
button {
  background-color: white;
  border: none;
  margin: 0;
  padding: 0;
  cursor: pointer;
}

.add-meet-but {
  width: 558px;
  height: 6.48vh;
  background-color: #00b268;
  font-size: 24px;
  color: #ffffff;
  box-shadow: 0px 2.85625px 2.85625px rgba(0, 0, 0, 0.25);
  border-radius: 20px;
}

/* Animations and hovers */

.add-meet-but {
  transition: 0.3s;
}

.add-meet-but:hover {
  transform: translateY(-6px);
  border: 1px solid #fff;
  transition: 0.3s;
}
</style>
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

.add-meet-but:hover,
.add-meet-but:focus {
  border: 2px solid darkgreen;
  transition: 0.3s;
}

@media (max-width: 768px) {
  .add-meet-but {
  width: 358px;
  height: 6.48vh;
  background-color: #00b268;
  font-size: 24px;
  color: #ffffff;
  box-shadow: 0px 2.85625px 2.85625px rgba(0, 0, 0, 0.25);
  border-radius: 20px;
}
}
</style>
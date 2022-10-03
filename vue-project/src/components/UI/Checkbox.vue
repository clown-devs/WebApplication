<template>
  <label class="checkbox-container">
    <input
      @input="(event) => $emit('update:checked', event.target.checked)"
      type="checkbox"
      :checked="checked"
      class="checkbox"
    />
    <span class="check-style"></span>
    <span class="option-content">{{ option }}</span>
  </label>
</template>

<script>
export default {
  props: {
    option: String,

    checked: {
      type: Boolean,
    },
  },
};
</script>

<style scoped>
* {
  box-sizing: border-box;
}

body {
  margin: 0;
  padding: 50px;
}

.checkbox-container {
  margin-bottom: 25px;
  display: block;
}

.checkbox {
  position: absolute;
  width: 1px;
  height: 1px;
  overflow: hidden;
  clip: rect(0 0 0 0);
  padding-left: 20px;
}

.check-style {
  position: absolute;
  width: 16px;
  height: 16px;
  border: 2px solid green;
}

.checkbox:checked + .check-style {
  background-image: url("/public/svg/check.svg");
}

.checkbox:focus + .check-style {
  box-shadow: 0 0 0 2px black;
}

.option-content {
  margin-left: 20px;
}
</style>

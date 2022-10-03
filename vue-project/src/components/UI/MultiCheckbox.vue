<template>
  <div class="multicheckbox-container">
    <checkbox 
        v-for="option in options" :key="option"
        v-model:checked="option.isSelected"
        :option="option.name"
        class="option-item"
    ></checkbox>
  </div>
</template>

<script>
import Checkbox from "@/components/UI/Checkbox.vue";

export default {
  components: {
    Checkbox,
  },

  props: {
    dataOptions: {
      type: Array,
      default: []
    }
  },

  data() {
    return {
      options: this.dataOptions
    };
  },
};
</script>

<style scoped>
.multicheckbox-container {
  overflow-y: auto;
  width: 100%;
  max-height: 200px;
}

.option-item {
  margin-top: 2px;
  margin-left: 5px;
}
</style>
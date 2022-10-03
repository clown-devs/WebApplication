<template>
  <div class="multicheckbox-container">
    <checkbox 
        v-for="option in options" :key="option"
        v-model:checked="option.isSelected"
        :option="option.name"
    ></checkbox>
  </div>
</template>

<script>
import Checkbox from "@/components/UI/Checkbox.vue";

export default {
  components: {
    Checkbox,
  },

  data() {
    return {
      options: [
        {
          name: "option1option1option1option1",
          id: 1,
          isSelected: false,
        },
        {
          name: "option2option2option2",
          id: 2,
          isSelected: false,
        },
        {
          name: "option3",
          id: 3,
          isSelected: false,
        },
      ],
    };
  },
};
</script>

<style scoped>
</style>
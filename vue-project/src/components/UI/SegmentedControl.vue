<template>
  <div class="segmented-control-container">
    <button 
        class="segmented-control-item" 
        :class="{active: isFirstItemActive}"
        @click="touchItem"
    >
      <slot class="segmented-control-item-name" name="item1"></slot>
    </button>
    <button 
        class="segmented-control-item" 
        :class="{active: !isFirstItemActive}"
        @click="touchItem"
    >
      <slot class="segmented-control-item-name" name="item2"></slot>
    </button>
  </div>
</template>

<script>
export default {
    data() {
        return {
            isFirstItemActive: true
        }
    },

    methods: {
        touchItem() {
            this.isFirstItemActive = !this.isFirstItemActive;
            this.$emit("selected", this.isFirstItemActive);
        }
    }
};
</script>

<style scoped>
.segmented-control-container {
  display: flex;
  justify-content: space-between;
  margin: 0;
}

.segmented-control-item {
  border-radius: 20px 20px 0px 0px;
  background-color: #e2eee3;
  flex: 1;
  border: 0;
  color: black;
  font-family: "Exo 2";
  font-size: 1rem;
  font-weight: 500;
  text-align: center;
  vertical-align: middle;
  padding: 0;
  margin: 0;
}

.active {
    background-color: #00B268;
    color: #ffffff;
}
</style>
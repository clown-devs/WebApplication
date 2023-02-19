<template>
  <div class="meet-places">
    <ul class="meet-places-list">
      <loading-indicate v-if="!places.length"></loading-indicate>
      <li v-for="place in places" :key="place" class="meet-place-item">
        <button @click='takePlace(place)' class="meet-place-item-button"> {{ place.name }}</button>
      </li>
    </ul>
  </div>
</template>

<script>
import api from "@/api";
import LoadingIndicate from "@/components/UI/LoadingIndicate";

export default {
  components: {
    LoadingIndicate
  },

  data() {
    return {
      places: [],
    };
  },

  async created() {
    this.places = await api.getPlaces();
  },

  methods: {
    takePlace(place) {
       this.$emit('takePlace', place)
    }
  }
};
</script>

<style scoped>
/* Style for main block*/
.meet-places {
  justify-content: center;
  display: flex;

}

/* Style of list */
.meet-places-list {
  background: #FFFFFF;
  border: 1px solid #BABABA;
  border-radius: 30px;
  margin: 0;
  padding: 3% 15% 3% 15%;
  overflow: scroll;
  overflow-x: auto;
}

.meet-places-list::-webkit-scrollbar {
  width: 0;
  height: 0;
}

.meet-place-item {
  list-style: none;
}

.meet-place-item button {
  list-style: none;
  display: flex;
  align-items: center;
  letter-spacing: 0.05em;
  line-height: 25px;
  font-weight: 500;
  cursor: pointer;
  font-size: 15px;
  color: #7A7474;
  white-space: nowrap;
}

.meet-place-item-button {
  background: none;
  border: 0;

}

.meet-place-item button:hover {
  color: #00B268;
}
</style>
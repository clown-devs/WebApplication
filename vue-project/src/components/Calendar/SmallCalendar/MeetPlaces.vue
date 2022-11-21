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
  margin-top: 5px;
  margin-left: 120px;
/*  ВОТ ТУТ ЗАТЕРЕТЬ МАРГИН*/
}

/* Style of list */
.meet-places-list {
  background: #FFFFFF;
  border: 1px solid #BABABA;
  border-radius: 30px;
  min-height: 20%;
  min-width: 30%;
  margin: 0;
  padding: 15px 40px 15px 40px;
}

.meet-place-item {
  list-style: none;
}

.meet-place-item button {
  list-style: none;
  display: flex;
  align-items: center;
  letter-spacing: 0.05em;
  line-height: 30px;
  font-weight: 500;
  cursor: pointer;
  font-size: 20px;
  color: #7A7474;
}

.meet-place-item-button {
  background: none;
  border: 0;

}

.meet-place-item button:hover {
  color: #00B268;
}
</style>
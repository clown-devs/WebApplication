<template>
  <div class="contact-item">
    <div class="contact-item__header">
      <div class="contact-name">
        <span class="contact-name-content">
          {{ contact.first_name }} {{ contact.second_name }}
          {{ contact.third_name }}
        </span>
      </div>
      <div class="header-buttons">
        <edit-button @click="editContact"></edit-button>
        <show-info-btn
          class="show-info-btn"
          @click="touchContactInfoButton"
          :showInfo="showContactInfo"
        ></show-info-btn>
      </div>
    </div>
    <div class="contact-details__content" v-if="showContactInfo">
      <div class="contact-email">
        <p class="contact-email-content">Почта: {{ contact.email }}</p>
      </div>
      <div class="contact-position">
        <p class="contact-position-content">
          Должность: {{ contact.position }}
        </p>
      </div>
      <div class="contact-number">
        <p class="contact-number-content">Телефон: {{ contact.phone }}</p>
      </div>
    </div>
  </div>
</template>

<script>
import ShowInfoBtn from "@/components/UI/ShowInfoButton.vue";
import EditButton from "@/components/UI/EditButton.vue";

export default {
  components: { ShowInfoBtn, EditButton },

  data() {
    return {
      showContactInfo: false,
      contact: this.contactData
    };
  },

  props: {
    contactData: {
      type: Object,
      default: {},
    },
  },

  methods: {
    touchContactInfoButton() {
      this.showContactInfo = !this.showContactInfo;
    },

    async editContact() {
      const contact = this.contact;
      contact.email = "test@gmail.com";
      this.contact = await api.editContact(contact);

      this.$emit("edit", this.contact);
    },
  },
};
</script>

<style scoped>
.contact-item {
  min-height: 70px;
  border: 0.72px solid #f1fbf2;
  box-shadow: 0px 2.86px 2.86px rgba(0, 0, 0, 0.25);
  border-radius: 10px;
  background-color: #f1fbf2;
  margin: 10px 0 10px 0;
  margin-left: 11.5%;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.contact-item__header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 70px;
  width: 100%;
}

.contact-name-content {
  margin-left: 20px;
  font-family: "Exo 2";
  font-weight: 700;
  font-size: 1.25rem;
  transition-duration: 0.5s;
}

.header-buttons {
  margin-right: 2rem;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.contact-managers-content {
  font-family: "Exo 2";
  font-weight: 700;
  font-size: 1.25rem;
  color: #7a7474;
  cursor: pointer;
}

.contact-details__content {
  margin-left: 20px;
  font-size: 1.25rem;
  font-weight: 500;
  font-family: "Exo 2";
}

.contact-number {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.contact-position-content {
  margin: 0;
}

.contact-managers-content {
  margin-right: 2rem;
}

/* Animations and hovers */

.contact-item,
.contact-email-content,
.contact-position-content,
.contact-number-content {
  transition-duration: 0.5s;
}

.contact-item:hover {
  transform: translateY(-3px);
  transition-duration: 0.5s;
}

.contact-name-content:hover,
.contact-email-content:hover,
.contact-position-content:hover,
.contact-number-content:hover {
  transition-duration: 0.5s;
  color: #00b268;
}
</style>
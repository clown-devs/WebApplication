<template>
  <div 
    class="contact-item" 
    v-if="this.contact !== undefined"
    @click="touchContactInfoButton"
  >
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
          :showInfo="showContactInfo"
        ></show-info-btn>
      </div>
    </div>
    <div class="contact-details__content" v-if="showContactInfo">
      <div class="contact-email">
        <p class="contact-email-content">Почта: {{ contact.email }}</p>
        <i class="material-icons copy-btn" @click="copyContentToBuffer(contact.email)"> content_copy </i>
      </div>
      <div class="contact-position">
        <p class="contact-position-content">
          Должность: {{ contact.position }}
        </p>
        <i class="material-icons copy-btn" @click="copyContentToBuffer(contact.position)"> content_copy </i>
      </div>
      <div class="contact-number">
        <p class="contact-number-content">Телефон: {{ contact.phone }}</p>
        <i class="material-icons copy-btn" @click="copyContentToBuffer(contact.phone)"> content_copy </i>
      </div>
    </div>
  </div>
  <div class="contact-item empty" v-else>
    <span>Нет контактов</span>
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
      contact: this.contactData,
    };
  },

  props: {
    contactData: {
      type: Object,
      default: undefined,
    },
  },

  methods: {
    touchContactInfoButton() {
      this.showContactInfo = !this.showContactInfo;
    },

    async editContact() {
      this.$emit("edit", this.contact);
    },

    copyContentToBuffer(content) {
      navigator.clipboard.writeText(content === null ? "" : content);
    }
  },
};
</script>

<style scoped>
.contact-item {
  min-height: 70px;
  border: 0.72px solid #f1fbf2;
  box-shadow: 0px 2.86px 2.86px rgba(0, 0, 0, 0.25);
  border-radius: 10px;
  background-color: #fff;
  margin: 10px 0 10px 0;
  margin-left: 11.5%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  cursor: pointer;
}

.empty {
  align-items: center;
  font-family: "Exo 2";
  font-weight: 700;
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
  cursor: default;
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

.contact-position,
.contact-email {
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

.copy-btn {
  margin-right: 2rem;
  color: #7a7474;
  cursor: pointer;
}

.contact-email,
.contact-position-content,
.contact-number {
  cursor: default;
}

/* Animations and hovers */

.contact-item,
.copy-btn {
  transition-duration: 0.5s;
}

.contact-item:hover {
  transition-duration: 0.5s;
  background: #e5e4e4;
}

.copy-btn:hover {
  transition-duration: 0.5s;
  color: #00b268;
}
</style>
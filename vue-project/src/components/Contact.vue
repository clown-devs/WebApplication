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
        <div class="contact-info">
          <a class="contact-info-content" @click="touchContactInfoButton">{{
            buttonTitle
          }}</a>
        </div>
        <a class="edit-contact">...</a>
      </div>
    </div>
    <div class="contact-details__content" v-if="showContactInfo">
      <div class="contact-inn">
        <p class="contact-inn-content">ИНН: {{ contact.client_inn }}</p>
      </div>
      <div class="contact-position">
        <p class="contact-position-content">
          Должность: {{ contact.position }}
        </p>
      </div>
      <div class="contact-number">
        <p class="contact-number-content">Телефон: {{ contact.phone }}</p>
        <div class="contact-managers">
          <a class="contact-managers-content">
            Менеджеры
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      buttonTitle: "Подробнее",
      showContactInfo: false,
    };
  },

  props: {
    contact: {
      type: Object,
      default: {},
    },
  },

  methods: {
    touchContactInfoButton() {
      if (!this.showContactInfo) {
        this.buttonTitle = "Скрыть";
      } else {
        this.buttonTitle = "Подробнее";
      }

      this.showContactInfo = !this.showContactInfo;
    },
  },
};
</script>

<style scoped>
.contact-item {
  min-height: 70px;
  border: 0.714062px solid #f1fbf2;
  box-shadow: 0px 2.85625px 2.85625px rgba(0, 0, 0, 0.25);
  border-radius: 10px;
  background-color: #f1fbf2;
  margin: 10px 0 10px 0;
  margin-left: 223px;
  display: flex;
  flex-direction: column;
}

.contact-item__header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
}

.contact-name-content {
  margin-left: 20px;
  font-family: "Exo 2";
  font-weight: 700;
  font-size: 1.25rem;
  transition-duration: 1s;
}

.header-buttons {
  margin-right: 2rem;
  display: flex;
  align-items: center;
  gap: 1rem;
}

.edit-contact,
.contact-managers-content {
  font-family: "Exo 2";
  font-weight: 700;
  font-size: 1.25rem;
  color: #7a7474;
}

.contact-info,
.edit-contact {
  cursor: pointer;
}

.contact-info-content {
  font-size: 1.25rem;
  font-weight: 700;
  font-family: "Exo 2";
  color: #7a7474;
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
.contact-info-content,
.edit-contact {
  transition-duration: 0.5s;
}

.contact-item:hover {
  border: 3px solid green;
  transition-duration: 0.5s;
}

.contact-name-content:hover,
.contact-info-content:hover,
.edit-contact:hover {
  transition-duration: 0.5s;
  color: #00b268;
}
</style>
<template>
  <div class="login-container">
    <form @submit.prevent>
      <h1 class="title">Вход в систему</h1>
      <input
        type="text"
        id="login"
        name="user_login"
        v-model="username"
        placeholder="Введите логин..."
        :class="{ invalid: v$.username.$error }"
      />
      <small v-if="v$.username.$error" class="validate-error-message">
        Логин не должен быть пустым!
      </small>
      
      <input
        type="password"
        id="password"
        name="user_password"
        v-model="password"
        placeholder="Введите пароль..."
        :class="{ invalid: v$.password.$error}"
      />
      <small v-if="v$.password.$dirty && v$.password.required.$invalid" class="validate-error-message">
        Пароль является обязательным полем!
      </small>
      <small v-if="v$.password.minLength.$invalid" class="validate-error-message">
        Длина пароля должна быть больше 3 символов!
      </small>

      <div class="save-container">
        <input
          type="checkbox"
          id="save"
          name="selected_save"
          v-model="isSavedSession"
        />
        <label for="save" class="save-label">Запомнить</label>
      </div>
      <button type="submit" @click="touchLogIn">Далее</button>

      <div v-if="isHaveErrorAuth()" class="error-container">
        {{ this.$store.state.auth.errorMessage }}
      </div>
    </form>
  </div>
</template>

<script>
import { mapActions } from "vuex";
import useVuelidate from "@vuelidate/core";
import { required, minLength } from "@vuelidate/validators";

export default {
  setup() {
    return {
      v$: useVuelidate(),
    };
  },

  data() {
    return {
      username: "",
      password: "",
      isSavedSession: false,
      passwordPlaceholder: "Введите пароль...",
      loginPlaceholder: "Введите логин...",
    };
  },

  validations() {
    return {
      username: { required },
      password: { required, minLength: minLength(4) },
    };
  },

  methods: {
    ...mapActions(["logIn"]),

    async touchLogIn() {

      const isCorrectForm = await this.v$.$validate();
      if (!isCorrectForm) {
        return;
      }

      const isAuth = await this.logIn({
        username: this.username,
        password: this.password,
        isSavedSession: this.isSavedSession,
      });

      if (isAuth) {
        this.$router.push("/");
        return;
      }

      this.$router.push("/login");
    },

    isHaveErrorAuth() {
      return this.$store.state.auth.errorMessage !== "";
    },

    isValidForm() {
      const MAX_PASSWORD_LEN = 4;

      this.validationErrors = [];

      if (this.username.length && this.password.length >= MAX_PASSWORD_LEN) {
        return true;
      }

      return false;
    },
  },
};
</script>

<style scoped>
/* Main */

.login-container {
  width: 100vw;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}

form {
  background: rgba(255, 255, 255, 0.58);
  border: 1px solid #47af52;
  border-radius: 30px;
  width: 596px;
  height: 660px;
  display: flex;
  flex-direction: column;
  align-items: center;
  flex-wrap: wrap;
}

h1 {
  color: #00b268;
  font-size: 2rem;
  text-align: center;
  font-weight: normal;
  margin-top: 16.36%;
  margin-bottom: 13.03%;
  width: 80.8%;
}

input {
  width: 80.8%;
  background: rgba(255, 255, 255, 0);
  border-bottom-color: #00b268;
  border-top-style: none;
  border-left-style: none;
  border-right-style: none;
}

input,
label {
  font-size: 1.25rem;
  color: #00000085;
}

input[type="text"] {
  margin-bottom: 18.18%;
}

input[type="password"] {
  margin-bottom: 5.45%;
  background: rgba(255, 255, 255, 0);
}

.save-container {
  display: flex;
  align-self: flex-start;
  align-items: center;
  margin-left: 9.56%;
  margin-bottom: 8.48%;
}

input[type="checkbox"] {
  min-width: 20px;
  min-height: 20px;
  margin-right: 10px;
}

button {
  background: #47af52;
  color: #ffffff;
  font-size: 1.5rem;
  border-radius: 30px;
  border-width: 0;
  width: 237px;
  height: 60px;
}

.error-container {
  margin-top: 2.5rem;
  color: red;
  font-weight: bolder;
}

/* Responcive CSS */

@media (max-width: 1200px) {
  form {
    width: 60vw;
    height: 70vh;
    justify-content: center;
    gap: 2.5rem;
  }

  h1,
  input[type="text"],
  input[type="password"],
  label,
  input[type="checkbox"],
  button,
  .save-container,
  .error-container {
    margin-top: 0;
    margin-bottom: 0;
  }

  button {
    width: 50%;
  }
}

@media (max-width: 992px) {
  form {
    gap: 2rem;
    height: 100vh;
    width: 100vw;
    border: inherit;
    border-radius: 0;
  }
}

@media (max-width: 767px) {
  form {
    height: 100vh;
    width: 100vw;
    border-radius: 0;
    border: inherit;
  }

  button {
    width: 80%;
    height: 2.5rem;
  }
}

@media (max-height: 415px) {
  form {
    height: 100vh;
    gap: 1.3rem;
  }
}

/* Validation styles */

.invalid {
  border-bottom-color: red;
}

.validate-error-message {
  color: red;
  font-weight: bolder;
}
</style>
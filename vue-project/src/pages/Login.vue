<template>
  <div class="login">
    <form action="/my-handling-form-page" method="post">
      <h1 class="title">Вход в систему</h1>
      <input type="text" id="log" name="user_login" value="Логин" />
      <input
        type="password"
        id="password"
        name="user_password"
        value="Пароль"
      />
      <div class="save-container">
        <input type="checkbox" id="save" name="selected_save" />
        <label for="save" class="save-label">Запомнить</label>
      </div>
      <div class="submit-button">
        <button type="submit">Далее</button>
      </div>
    </form>
  </div>
</template>

<script>
export default {};
</script>

<style scoped>
.login {
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
  width: 31.04%;
  height: 61.11%;
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
  padding: 1.125rem 2.9rem;
}

@media (max-width: 1200px) {
  form {
    min-height: 660px;
    min-width: 596px;
    justify-content: center;
    gap: 2rem;
  }

  h1,
  .save-container,
  button,
  input[type="text"],
  input[type="password"] {
    margin-top: 0;
    margin-bottom: 0;
  }

}

@media (max-width: 992px) {
  form {
    min-width: 100vw;
    min-height: 100vh;
    justify-content: center;
    gap: 2rem;
    font-size: 2rem;
    border-radius: 0;
  }

  h1,
  .save-container,
  button,
  label,
  input[type="text"],
  input[type="password"] {
    margin-top: 0;
    margin-bottom: 0;
  }

  h1 {
    font-size: 3rem;
  }

  .save-container,
  button,
  label,
  input[type="text"],
  input[type="password"] {
    font-size: 2rem;
  }

  button {
    width: 50vw;
    padding: 1rem 2rem;
  }

}

@media (max-width: 767px) {
  form {
    min-height: 100vh;
    min-width: 100vw;
    border-radius: 0;
    gap: 1rem;
  }

   h1 {
    font-size: 2.5rem;
  }

  .save-container,
  button,
  label,
  input[type="text"],
  input[type="password"] {
    font-size: 1.2rem;
  }

  button {
    width: 80.8vw;
  }

}
</style>
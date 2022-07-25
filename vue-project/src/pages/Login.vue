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
      <button type="submit">Далее</button>
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
  .save-container {
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
</style>
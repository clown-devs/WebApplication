<template>
  <div class="login">
    <form action="/my-handling-form-page" method="post">
      <label>Вход в систему</label>
      <input type="text" id="log" name="user_login" value="Логин" />
      <input type="text" id="password" name="user_password" value="Пароль" />
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
  justify-content: space-around;
  flex-shrink: 2;
}

label {
  color: #00b268;
  font-size: 2rem;
  text-align: center;
}

button {
  background: #47af52;
  color: #ffffff;
  font-size: 1.5rem;
  border-radius: 30px;
  border-width: 0;
  min-height: 100%;
  min-width: 100%;
}

.submit-button {
    width: 39.76%;
    height: 9.09%;
}

.save-container {
  display: flex;
  align-self: flex-start;
  align-items: center;
  margin-left: 9.56%;
}

input,
.save-label {
  font-size: 1.25rem;
  color: #00000085;
}

input[type="checkbox"] {
  min-height: 17px;
  min-width: 17px;
}

input {
  width: 80.8%;
  background: rgba(255, 255, 255, 0);
  border-bottom-color: #00b268;
  border-top-style: none;
  border-left-style: none;
  border-right-style: none;
}
</style>
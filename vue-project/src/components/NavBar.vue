<template>
  <header class="header">
    <div class="container header-container">
      <div class="title-name title-names">Sber clients</div>

      <div class="header-burger">
        <span></span>
      </div>
      <div class="header-nav-bar">
        <ul class="nav-bar-list">
          <li class="nav-bar-list-item" @click="$router.push('/')">Главная страница</li>

          <li class="nav-bar-list-item" @click="$router.push('/clients')">Мои клиенты</li>

          <li class="nav-bar-list-item" @click="$router.push('/calendar')">Календарь</li>

          <li class="nav-bar-list-item" @click="$router.push('/results')">Отчёты по встречам</li>
        </ul>
      </div>

      <button class="nav-bar-notify-button">
        <img src="/svg/notify.svg" alt="" class="nav-bar-notify-icon" />
      </button>
      <div class="nav-bar-icon-button">
        <button class="nav-bar-add-meet">
          <img src="/svg/add-meet.svg" alt="" class="nav-bar-icon-add" />
        </button>

        <button class="nav-bar-profile">
          <img src="/svg/profile.svg" alt="" class="nav-bar-profile-icon" />
        </button>
      </div>
    </div>
    <hr class="line-header" />
  </header>

  <!-- <main>
    <div class="container">
      <div class="content-text">
        <p>
          Lorem ipsum dolor sit amet consectetur, adipisicing elit. Perferendis
          iusto cum voluptatum veniam itaque, earum, veritatis reiciendis
          distinctio eaque unde magni illo adipisci consequuntur, dignissimos
          vel alias facere ratione nihil!
        </p>
        <p>
          Lorem ipsum dolor sit amet consectetur, adipisicing elit. Perferendis
          iusto cum voluptatum veniam itaque, earum, veritatis reiciendis
          distinctio eaque unde magni illo adipisci consequuntur, dignissimos
          vel alias facere ratione nihil!
        </p>
        <p>
          Lorem ipsum dolor sit amet consectetur, adipisicing elit. Perferendis
          iusto cum voluptatum veniam itaque, earum, veritatis reiciendis
          distinctio eaque unde magni illo adipisci consequuntur, dignissimos
          vel alias facere ratione nihil!
        </p>
        <p>
          Lorem ipsum dolor sit amet consectetur, adipisicing elit. Perferendis
          iusto cum voluptatum veniam itaque, earum, veritatis reiciendis
          distinctio eaque unde magni illo adipisci consequuntur, dignissimos
          vel alias facere ratione nihil!
        </p>
      </div>
    </div>
  </main> -->
</template>

<script>
export default {};
</script>

<style scoped>
.header-container {
  display: flex;
  align-items: center;
}

.container {
  max-width: 1520px;
  height: 110px;
  margin-left: 200px;
  margin-right: 200px;
}

@media (max-width: 1520px) {
  .container {
    max-width: 970px;
  }
}

@media (max-width: 992px) {
  .container {
    max-width: 750px;
  }
}

@media (max-width: 767px) {
  .container {
    max-width: none;
  }
}

.nav-bar-list {
  display: flex;
}

.nav-bar-icon-button {
  display: flex;
  margin-left: auto;
  justify-content: center;
  align-items: center;
}

.nav-bar-icon-button-two {
  margin-left: auto;
}

.title-name {
  margin-right: 60px;
}

ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.line-header {
  width: 100%;
}


.nav-bar-list-item {
  font-size: 16px;
  margin-right: 52px;
}

.nav-bar-add-meet {
  margin-right: 34px;
}

.nav-bar-notify-button {
  height: 40px;

  background-color: white;
  border: none;
  margin: 0;
  padding: 0;
}

.nav-bar-add-meet {
  height: 60px;

  background-color: white;
  border: none;
  margin-top: 0;
  margin-bottom: 0;
  margin-left: 0;
  padding: 0;
  display: flex;
}

.nav-bar-profile {
  height: 70px;

  background-color: white;
  border: none;
  margin-top: 0;
  margin-bottom: 0;
  margin-left: 0;
  padding: 0;
}
</style>
<template>
  <div class="popup_wrapper" ref="popup_wrapper">
    <form class="v-popup animation" @submit.prevent>
      <div class="v-popup__header">
        <span class="popup-title"> Отчёт о встрече </span>
        <span>
          <i class="material-icons" @click="closeReportPopup"> close </i>
        </span>
      </div>

      <div class="v-popup__content">
        <div class="main-info">
          <h5 class="main-title">Общая информация</h5>
          <table class="table-main-info" cellpadding="5" cellspacing="5">
            <tr class="table-row-main-info">
              <td class="row-data-main-info first-column-main-meet">
                Источник:
              </td>
              <td class="row-data-main-info">
                <input type="text" class="input-report-meet">
              </td>
            </tr>
            <tr class="table-row-main-info">
              <td class="row-data-main-info ">
                Дата:
              </td>
              <td class="row-data-main-info">
                <input type="text" class="input-report-meet">
              </td>
            </tr>
            <tr class="table-row-main-info">
              <td class="row-data-main-info ">
                Место:
              </td>
              <td class="row-data-main-info">
                <input type="text" class="input-report-meet">
              </td>
            </tr>
          </table>
        </div>

        <div class="members">
          <h5>1. Участники</h5>
          <table class="table-members-info" cellpadding="5" cellspacing="5">
            <tr class="table-row-members-headers">
              <th class="table-header-members first-column-members">ФИО</th>
              <th class="table-header-members">Должность</th>
            </tr>
            <tr class="table-row-members">
              <td class="row-data-members-title">Со стороны ПАО "Сбербанк России"</td>
            </tr>
            <tr class="table-row-members" v-for="employer in selectedUsersInMeeting">
              <td class="row-data-members" v-if="makeNewEmployer(employer)">
                {{ employer.name }}
              </td>
              <td class="row-data-members" v-else>
                <input type="text" class="input-report-meet">
              </td>
              <td class="row-data-members">
                <input type="text" class="input-report-meet">
              </td>
            </tr>
            <tr class="table-row-members">
              <td></td>
              <td class="add-new-employer">
                <button class="button-add-new-employer" @click="addNewEmployer">+</button>
              </td>
            </tr>
            <tr class="table-row-members">
              <td class="row-data-members-title">Со стороны клиента</td>
            </tr>
            <tr class="table-row-members" v-for="client in newClients">
              <td class="row-data-members">
                <input type="text" class="input-report-meet">
              </td>
              <td class="row-data-members">
                <input type="text" class="input-report-meet">
              </td>
            </tr>
            <tr class="add-new-clients">
              <td></td>
              <td class="add-new-employer">
                <button class="button-add-new-clients" @click="addNewClients">+</button>
              </td>
            </tr>
          </table>
        </div>

        <div class="content-meet">
          <h5 class="main-title">Содержание встречи</h5>
          <table class="table-content-meet" cellpadding="5" cellspacing="5">
            <tr class="table-row-content-meet">
              <td class="row-data-content-meet">
                <textarea class="content-meet-input"></textarea>
              </td>
            </tr>
          </table>
        </div>

        <div class="accepted-decision">
          <h5 class="accepted-decision-title">Принятые решения и слудующие шаги</h5>
          <table class="table-accepted-decision" cellpadding="5" cellspacing="5">
            <tr class="table-row-accepted-headers">
              <th class="table-header-accepted first-column-accepted-decision">№</th>
              <th class="table-header-accepted">Мероприятия</th>
              <th class="table-header-accepted">Участники</th>
              <th class="table-header-accepted">Срок</th>
              <th class="table-header-accepted">Комментарии</th>
            </tr>
            <tr class="table-row-accepted-decision" v-for="decision in newAcceptedDecision">
              <td class="row-data-accepted-decision">
                {{ decision.id }}
              </td>
              <td class="row-data-accepted-decision">
                <input type="text" class="input-report-meet">
              </td>
              <td class="row-data-accepted-decision">
                <input type="text" class="input-report-meet">
              </td>
              <td class="row-data-accepted-decision">
                <input type="text" class="input-report-meet">
              </td>
              <td class="row-data-accepted-decision">
                <input type="text" class="input-report-meet">
              </td>
            </tr>
            <tr class="add-new-accepted-decision">
              <td></td>
              <td></td>
              <td></td>
              <td></td>
              <td class="add-new-employer">
                <button class="button-add-new-decision" @click="addNewAcceptedDecision">+</button>
              </td>
            </tr>
          </table>
        </div>
      </div>

      <div class="v-popup__footer">
        <addButton>Сохранить</addButton>
      </div>
    </form>
  </div>
</template>

<script>
import addButton from "@/components/UI/AddButton";

export default {
  components: {addButton},

  data() {
    return {
      tmpIdEmployer: 25,
      lengthSelectedEmployersBeforeAdd: this.selectedUsersInMeeting.length,
      newClients: [],
      tmpIdClient: 25,
      newAcceptedDecision: [],
      tmpIdAcceptedDecision: 0,
    }
  },

  props: {
    selectedUsersInMeeting:
        {
          type: Object,
          default: null,
        },

  },

  methods: {
    closeReportPopup() {
      this.$emit("closeReportPopup");
    },

    makeNewEmployer(employer) {
      return employer.name !== "";
    },

    addNewEmployer() {
      let newEmployer = {
        name: "",
        id: this.tmpIdEmployer + 1,
        direction: "",
      }
      this.tmpIdEmployer += 1;
      this.selectedUsersInMeeting.push(newEmployer)
    },

    addNewClients() {
      let newClient = {
        name: "",
        id: this.tmpIdClient + 1,
        position: "",
      }
      this.tmpIdClient += 1;
      this.newClients.push(newClient)
    },

    addNewAcceptedDecision() {
      let newDecision = {
        id: this.tmpIdAcceptedDecision + 1,
        event: "",
        members: [],
        term: "",
        comment: ""
      }
      this.tmpIdAcceptedDecision += 1;
      this.newAcceptedDecision.push(newDecision)
    }
  },
  async mounted() {
    let vm = this;

    document.addEventListener("click", function (item) {
      if (item.target === vm.$refs["popup_wrapper"]) {
        vm.closeReportPopup();
      }
    });
  }
}
</script>

<style scoped lang="scss">
//Main style popup

.popup_wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  position: fixed;
  right: 0;
  left: 0;
  top: 0;
  background: rgba(0, 0, 0, 0.5);
  bottom: 0;
  z-index: 100;
  overflow: auto;
}

.v-popup {
  border-radius: 25px;
  position: fixed;
  background: #ffffff;
  box-shadow: 0 0 17px 0 #e7e7e7;
  z-index: 10;
  width: 74%;
  height: 90%;
  overflow: auto;

  &__header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin: 2rem;
    font-family: "Exo 2", serif;
    font-weight: 700;
    font-size: 1.5rem;
  }

  &__footer {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    margin: 2rem;
  }

  &__content {
    flex-direction: column;
    gap: 2rem;
    margin: 2rem 3rem 0 3rem;
    overflow-y: hidden;
  }
}

.popup-title {
  flex: 2;
  display: flex;
  align-items: center;
  justify-content: center;
}

//Table css main info
.table-main-info, .table-members-info, .table-content-meet, .table-accepted-decision {
  border: 1px solid black;
  width: 100%;
  border-collapse: collapse;
}

.row-data-main-info, .table-row-members, .table-header-members, .row-data-members, .table-header-accepted, .table-row-accepted-decision, .row-data-accepted-decision {
  border: 1px solid black;
}

.first-column-main-meet {
  width: 20%;
}

.first-column-members {
  width: 40%;
}

.first-column-accepted-decision {
  width: 5%;
}

.row-data-members-title {
  font-weight: 600;
}

.input-report-meet, .content-meet-input {
  border: 0;
  outline: none;
  width: 100%;
}

.content-meet-input {
  resize: none;
  height: 50px;
}

.button-add-new-employer, .button-add-new-clients, .button-add-new-decision {
  display: flex;
  margin-left: auto;
  margin-right: 5px;
  border-radius: 50%;
  border-width: 1px;
  cursor: pointer;
  font-size: 20px;
  border-bottom-color: #7a7474;
}

.input-report-name {
  display: none;
}

//Animation
.animation {
  animation-duration: 0.55s;
  animation-fill-mode: both;
  animation-name: fadeInDown;
}

.material-icons {
  cursor: pointer;
}

.material-icons {
  transition-duration: 0.5s;
}

.material-icons:hover {
  transition-duration: 0.5s;
  color: #00b268;
}

.v-popup::-webkit-scrollbar {
  display: none;
}
</style>
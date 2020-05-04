<template>
  <div class="moderate">
    <div class="container">
      <h1>Moderation dashboard</h1>
      <hr />
      <AuthForm v-if="token === null" @authenticated="setToken" />
      <div v-else id="dashboard">
        <div id="controlPanel">
          <RoomList
            id="roomList"
            ref="roomList"
            :token="token"
            :selectedRoom="selectedRoom"
            @select="room => (selectedRoom = room)"
          />
          <StateList
            id="stateList"
            v-if="selectedRoom != null"
            :selectedState="selectedState"
            @select="state => (selectedState = state)"
          />
        </div>
        <div id="submissions">
          <strong v-if="selectedState"
            ><span
              >Submissions in '{{ selectedRoom }}' of state '{{
                selectedState
              }}':</span
            >
            <font-awesome-icon @click="refresh" class="icn" icon="redo-alt"
          /></strong>
          <strong v-else>Select a room and state</strong>
          <hr />
          <SubmissionList
            id="submissionList"
            ref="submissionList"
            v-if="selectedRoom != null"
            :token="token"
            :selectedRoom="selectedRoom"
            :selectedState="selectedState"
            @refresh="refresh"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import AuthForm from "@/components/AuthForm.vue";
import RoomList from "@/components/RoomList.vue";
import StateList from "@/components/StateList.vue";
import SubmissionList from "@/components/SubmissionList.vue";
export default {
  name: "moderate",
  components: {
    AuthForm,
    RoomList,
    StateList,
    SubmissionList
  },
  data() {
    return {
      token: null,
      selectedRoom: null,
      selectedState: null
    };
  },
  methods: {
    setToken(token) {
      this.token = token;
    },
    refresh() {
      this.$refs.submissionList.refresh();
      this.$refs.roomList.refresh();
    }
  }
};
</script>

<style lang="scss" scoped>
.moderate {
  height: 100%;
  background-color: var(--background-join);
  color: var(--primary-join);

  background-image: url("/img/stripe.svg");
  background-repeat: repeat-y;
  background-position-x: 0.8rem;

  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.container {
  padding: 0 1.5rem 1rem 3.5rem;
  height: 100%;

  display: flex;
  flex-direction: column;

  @media (min-width: 992px) {
    padding-left: 8rem;
    padding-right: 6rem;
  }

  font-family: monospace;
  font-size: 1.2rem;
  color: var(--primary-join);
}

hr {
  margin: $spacer * 2 0;
  border: 0;
  border-bottom: 1px dashed var(--secondary-join);
}

p {
  margin-bottom: 1.5rem;
  line-height: 1.2;
}
strong {
  margin-top: $spacer;
  display: flex;
  justify-content: space-between;
}
.icn {
  margin: 0 $spacer * 2;
}

#dashboard {
  display: flex;
  height: 0;
  flex-grow: 1;
}
@media (orientation: portrait) {
  #dashboard {
    font-size: 50%;
  }
}
#controlPanel {
  margin-right: 4 * $spacer;
}
#submissions {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
}
#submissionList {
  overflow-y: scroll;
  height: 0;
  flex-grow: 1;
}
</style>

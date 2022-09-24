<template>
  <div class="moderate">
    <div class="container">
      <header>
        <span id="sign-out" v-if="token" @click="discardToken">
          <font-awesome-icon icon="sign-out-alt" title="sign out" /> Sign Out
        </span>
        <h1>Moderation dashboard</h1>
        <font-awesome-icon
          v-if="token"
          id="refresh"
          :class="{ active: this.refreshing }"
          :disabled="refreshing"
          @click="refresh"
          class="icn"
          icon="redo-alt"
        />
      </header>

      <hr />

      <AuthForm v-if="token === null" @authenticated="setToken" />

      <div v-else id="dashboard">
        <div id="controlPanel">
          <RoomList
            id="roomList"
            ref="roomList"
            :token="token"
            :rooms="rooms"
            :selectedRoom="selectedRoom"
            :disabled="refreshing"
            @select="
              room => {
                this.selectedRoom = room;
                this.refresh();
              }
            "
          />

          <StateList
            id="stateList"
            ref="stateList"
            v-if="selectedRoom"
            :selectedState="selectedState"
            :selectedRoom="selectedRoom"
            :disabled="refreshing"
            @select="
              state => {
                this.selectedState = state;
                this.refresh();
              }
            "
          />
        </div>

        <div id="moderatedMessages">
          <strong v-if="selectedState && selectedRoom">
            {{ { invisible: "Invisible", visible: "Visible" }[selectedState] }}
            Moderated messages in room '{{ selectedRoom.Name }}':
          </strong>
          <strong v-else-if="!selectedRoom">Select a room.</strong>
          <strong v-else-if="!selectedState">Select a state.</strong>

          <hr />

          <ModeratedMessageList
            id="moderatedMessageList"
            ref="moderatedMessageList"
            v-if="selectedRoom && selectedState"
            :token="token"
            :selectedRoomName="selectedRoom.Name"
            :selectedState="selectedState"
            :disabled="refreshing"
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
import ModeratedMessageList from "@/components/ModeratedMessageList.vue";

export default {
  name: "moderate",
  components: {
    AuthForm,
    RoomList,
    StateList,
    ModeratedMessageList
  },
  data() {
    return {
      token: null,
      selectedRoom: null,
      selectedState: "visible",
      rooms: [],
      refreshing: false
    };
  },
  methods: {
    setToken(token) {
      this.token = token;
      this.refresh();
    },
    discardToken() {
      this.token = null;
    },
    refresh() {
      //We don't try to refresh if we're already refreshing
      if (this.refreshing) {
        return;
      }

      this.refreshing = true;
      setTimeout(
        function() {
          this.refreshing = false;
        }.bind(this),
        250
      );

      const url =
        window.location.origin +
        "/api/?method=get_moderated_rooms&token=" +
        this.token;
      const options = {
        method: "GET"
      };

      fetch(url, options)
        .then(resp => {
          return resp.text();
        })
        .then(result => {
          this.rooms = JSON.parse(result) || [];

          //If the moderated message list exists we need to refresh that too.
          if (this.$refs.moderatedMessageList) {
            this.$nextTick().then(this.$refs.moderatedMessageList.refresh);
          }

          //If a room is selected we need to update it with potentially new details
          if (this.selectedRoom) {
            for (var room of this.rooms) {
              if (room.Name === this.selectedRoom.Name) {
                this.selectedRoom = room;
                return;
              }
            }
          }
        });
    }
  }
};
</script>

<style lang="scss" scoped>
.moderate {
  height: 100%;
  background-color: var(--background-join);
  color: var(--primary-join);

  background-image: url("../../public/img/stripe.svg");
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

  @media (orientation: portrait) {
    #sign-out,
    #dashboard {
      font-size: 50%;
    }
  }
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
}

header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;

  #sign-out {
    width: 100%;
    margin-top: $spacer * 2;
  }

  h1 {
    max-width: 80%;
  }

  #refresh {
    margin: 0 $spacer * 2;
    transition: transform 0.25s;
    &.active {
      transform: rotate(360deg);
      color: $grey-l;
    }
  }
}

#dashboard {
  display: flex;
  height: 0;
  flex-grow: 1;
}
#controlPanel {
  margin-right: 4 * $spacer;
}
#moderatedMessages {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
}
#moderatedMessageList {
  overflow-y: scroll;
  height: 0;
  flex-grow: 1;
}
</style>

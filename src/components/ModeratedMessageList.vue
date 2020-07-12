<template>
  <ul>
    <li v-if="filteredModeratedMessages.length == 0">
      <p>There's nothing here!</p>
    </li>
    <ModeratedMessage
      class="moderatedMessage"
      v-for="moderatedMessage in filteredModeratedMessages"
      :key="moderatedMessage.msg.hash"
      :token="token"
      :moderatedMessage="moderatedMessage"
      :selectedRoomName="selectedRoomName"
      :selectedState="selectedState"
      :disabled="disabled"
      @changeState="
        newState => {
          moderatedMessage.State = newState;
          $emit('refresh');
        }
      "
    />
  </ul>
</template>

<script>
import ModeratedMessage from "@/components/ModeratedMessage.vue";

import RunlengthEncoder from "../assets/js/runlengthEncoder.js";
import { Message } from "../assets/js/message.js";
import colour from "../assets/js/colours.js";

export default {
  components: {
    ModeratedMessage
  },
  props: ["token", "selectedRoomName", "selectedState", "disabled"],
  data() {
    return {
      moderatedMessages: []
    };
  },
  mounted() {
    this.refresh();
  },
  computed: {
    filteredModeratedMessages() {
      return this.moderatedMessages.filter(
        moderatedMessage => moderatedMessage.State == this.selectedState
      );
    }
  },
  methods: {
    refresh() {
      const url =
        window.location.origin +
        "/api/?method=get_moderated_messages&token=" +
        this.token +
        "&room_id=" +
        this.selectedRoomName;
      const options = {
        method: "GET"
      };

      fetch(url, options)
        .then(resp => {
          return resp.text();
        })
        .then(result => {
          this.moderatedMessages = JSON.parse(result) || [];
          this.moderatedMessages.map(
            moderatedMessage =>
              (moderatedMessage.msg = new Message(
                RunlengthEncoder.decode(moderatedMessage.Message.Payload.Data),
                moderatedMessage.Message.Payload.Span,
                null,
                colour(moderatedMessage.Message.Payload.ColourIndex),
                moderatedMessage.Message.Time
              ))
          );
        });
    }
  }
};
</script>

<style lang="scss" scoped>
ul {
  width: 100%;
  margin: 0;
  padding: 0;

  p {
    text-align: center;
  }
}

.moderatedMessage {
  max-width: 60vw;
  margin: 0 auto;

  margin-bottom: $spacer * 2;
  padding: 0;
  display: block;
  width: 100%;

  > * {
    width: 100%;
    margin: 0;
    margin-bottom: $spacer;
  }
}
</style>

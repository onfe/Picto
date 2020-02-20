<template lang="html">
  <section>
    <div class="history">
      <div class="message" v-for="msg in history" v-bind:key="msg.id">
        <CanvasMessage v-if="msg.type == 'normal'" v-bind="msg" />
        <Announcement v-else-if="msg.type == 'announcement'" v-bind="msg" />
        <div class="text" v-else>{{ msg.text }}</div>
      </div>
    </div>
  </section>
</template>

<script>
import CanvasMessage from "@/components/CanvasMessage.vue";
import Announcement from "@/components/Announcement.vue";

export default {
  components: {
    CanvasMessage,
    Announcement
  },
  computed: {
    history() {
      return this.$store.state.messages.history;
    }
  }
};
</script>

<style lang="scss" scoped>
section {
  position: relative;
  height: 100%;
}

.history {
  padding: 1vw;
  overflow-y: scroll;
  display: flex;
  height: 100%;
  flex-direction: column;
}

.message {
  margin-top: 1vw;
}

/*
Problem: FireFox just doesn't scroll on a flexbox with flex-direction:column-reverse.

| Potential Solution            | Negatives                    | Positives                   |
| ----------------------------- | ---------------------------- | --------------------------- |
| flex-direction:column-reverse | Doesn't work on FireFox.     | Works on everything except  |
| on .history                   |                              | FireFox perfectly.          |
| ----------------------------- | ---------------------------- | --------------------------- |
| transform: scaleY(-1)         | Scroll direction is reversed | https://open.spotify.com/tr |
| on .history and .message      | (doesn't affect mobile);     | ack/5foxQ24C0x7W0B2OD46AJg? | 
|                               | it's just really dumb        | si=joaaiGIsTES52UQVX5bNoQ   | 

https://open.spotify.com/track/5foxQ24C0x7W0B2OD46AJg?si=joaaiGIsTES52UQVX5bNoQ
*/
.history, .message {
  transform: scaleY(-1);
}

.text {
  font-size: 0.85em;
  text-align: center;
  color: $grey;
  font-family: monospace;
}

</style>

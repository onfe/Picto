<template lang="html">
  <div id="submission">
    <CanvasMessage :msg="message" />
    <div class="controls">
      <ul class="size btn">
        <li title="Publish" :class="method == 'publish' ? 'selected' : ''">
          <font-awesome-icon class="icn" icon="thumbs-up" @click="publish" />
        </li>
        <li title="Reject" :class="method == 'reject' ? 'selected' : ''">
          <font-awesome-icon class="icn" icon="thumbs-down" @click="reject" />
        </li>
      </ul>
      <ul class="size btn">
        <li title="Submit" :class="method === '' ? 'disabled' : ''">
          <font-awesome-icon class="icn" icon="paper-plane" @click="submit" />
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import CanvasMessage from "@/components/CanvasMessage.vue";
import { Message } from "@/assets/js/message.js";
import colour from "@/assets/js/colours.js";
import RunlengthEncoder from "../assets/js/runlengthEncoder.js";

export default {
  components: {
    CanvasMessage
  },
  props: ["submission", "roomName", "token"],
  data() {
    return {
      method: ""
    };
  },
  computed: {
    message() {
      return new Message(
        RunlengthEncoder.decode(this.submission.Message.Data),
        this.submission.Message.Span,
        this.submission.Message.Sender,
        colour(this.submission.Message.ColourIndex),
        null
      );
    }
  },
  methods: {
    submit() {
      if (this.method == "") {
        return;
      }

      var url =
        window.location.origin +
        `/api/?method=${this.method}_submission&token=${this.token}&room_id=${this.roomName}&submission_id=${this.submission.ID}`;
      var options = {
        method: "GET"
      };

      setTimeout(() => {
        fetch(url, options)
          .then(resp => {
            return resp.text();
          })
          .then(
            function() {
              this.$emit("remove");
            }.bind(this)
          );
      }, 250);
    },
    reject() {
      this.method = this.method != "reject" ? "reject" : "";
    },
    publish() {
      this.method = this.method != "publish" ? "publish" : "";
    }
  }
};
</script>

<style lang="scss" scoped>
.controls,
.btn {
  height: $sidebar-width;
  display: flex;
  justify-content: flex-start;
}
.controls {
  padding: $spacer;
}
.btn li {
  flex-shrink: 1;
  width: $sidebar-width - 2 * $spacer;
  height: $sidebar-width - 2 * $spacer;
  padding-bottom: 0;
}
.disabled {
  color: var(--button-selected);
}
</style>

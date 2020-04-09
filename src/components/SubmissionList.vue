<template lang="html">
  <div class="submissions">
    <Submission
      class="message"
      v-for="submission in submissions"
      v-bind:key="submission.ID"
      :submission="submission"
      :roomName="roomName"
      :token="token"
      @remove="updateSubmissions"
    />
  </div>
</template>

<script>
import Submission from "@/components/Submission.vue";

export default {
  components: {
    Submission
  },
  props: ["token", "roomName"],
  data() {
    return {
      submissions: []
    };
  },
  methods: {
    updateSubmissions() {
      const url =
        window.location.origin +
        `/api/?method=get_submissions&token=${this.token}&room_id=${this.roomName}`;
      const options = {
        method: "GET"
      };

      setTimeout(() => {
        fetch(url, options)
          .then(resp => {
            return resp.text();
          })
          .then(result => {
            this.submissions = JSON.parse(result) || [];
            if (typeof this.submissions != "object") {
              this.submissions = [];
            }
          });
      }, 250);
    }
  },
  watch: {
    roomName() {
      this.updateSubmissions();
    }
  },
  mounted() {
    this.updateSubmissions();
  }
};
</script>

<style lang="scss" scoped>
.submissions {
  padding: $spacer;
  overflow-y: scroll;
  overflow-x: hidden;
  display: flex;
  flex-direction: column;
  height: 100%;
}

.message {
  margin-top: $spacer;
  width: $ratio-width;
}

.firefox {
  .history {
    flex-direction: column !important;
  }

  .history,
  .message {
    transform: scaleY(-1);
  }
}

.text {
  font-size: 0.85em;
  text-align: center;
  color: $grey;
  font-family: monospace;
}
</style>

<template>
  <ul>
    <li v-if="filteredSubmissions.length == 0">
      <p>There's nothing here!</p>
    </li>
    <Submission
      class="submission"
      v-for="submission in filteredSubmissions"
      :key="submission.msg.hash"
      :token="token"
      :submission="submission"
      :selectedRoomName="selectedRoomName"
      :selectedState="selectedState"
      :disabled="disabled"
      @changeState="
        newState => {
          submission.State = newState;
          $emit('refresh');
        }
      "
    />
  </ul>
</template>

<script>
import Submission from "@/components/Submission.vue";

import RunlengthEncoder from "../assets/js/runlengthEncoder.js";
import { Message } from "../assets/js/message.js";
import colour from "../assets/js/colours.js";

export default {
  components: {
    Submission
  },
  props: ["token", "selectedRoomName", "selectedState", "disabled"],
  data() {
    return {
      submissions: []
    };
  },
  mounted() {
    this.refresh();
  },
  computed: {
    filteredSubmissions() {
      return this.submissions.filter(
        submission => submission.State == this.selectedState
      );
    }
  },
  methods: {
    refresh() {
      const url =
        window.location.origin +
        "/api/?method=get_submissions&token=" +
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
          this.submissions = JSON.parse(result) || [];
          this.submissions.map(
            submission =>
              (submission.msg = new Message(
                RunlengthEncoder.decode(submission.Message.Payload.Data),
                submission.Message.Payload.Span,
                null,
                colour(submission.Message.Payload.ColourIndex),
                submission.Message.Time
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

.submission {
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

<template>
  <ul>
    <li
      v-for="(submission, index) in this.submissions.filter(
        submission => submission.State == selectedState
      )"
      :key="index"
    >
      <p>Author: "{{ submission.Message.Payload.Sender }}"</p>
      <CanvasMessage :msg="submission.msg" />
    </li>
  </ul>
</template>

<script>
import CanvasMessage from "@/components/CanvasMessage.vue";
import RunlengthEncoder from "../assets/js/runlengthEncoder.js";
import { Message } from "../assets/js/message.js";
import colour from "../assets/js/colours.js";

export default {
  components: {
    CanvasMessage
  },
  props: ["token", "selectedRoom", "selectedState"],
  data() {
    return {
      submissions: []
    };
  },
  mounted() {
    const url =
      window.location.origin +
      "/api/?method=get_submissions&token=" +
      this.token +
      "&room_id=" +
      this.selectedRoom;
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
};
</script>

<style lang="scss" scoped>
ul {
  width: 100%;
  margin: 0;
  padding: 0;
}

li {
  margin: 0;
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

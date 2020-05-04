<template>
  <li>
    <p>Author: '{{ submission.Message.Payload.Sender }}'</p>
    <CanvasMessage :msg="submission.msg" />
    <ul id="options">
      <li :class="{ btn: true, selected: newState == 'published' }">
        <font-awesome-icon
          @click="newState = 'published'"
          class="icn"
          icon="check"
          title="publish"
        />
      </li>
      <li :class="{ btn: true, selected: newState == 'submitted' }">
        <font-awesome-icon
          @click="newState = 'submitted'"
          class="icn"
          icon="hourglass"
          title="unpublish"
        />
      </li>
      <li :class="{ btn: true, selected: newState == 'deleted' }">
        <font-awesome-icon
          @click="newState = 'deleted'"
          class="icn"
          icon="times"
          title="reject"
        />
      </li>
      <li :class="{ btn: true, selected: newState == 'offensive' }">
        <font-awesome-icon
          @click="newState = 'offensive'"
          class="icn"
          icon="exclamation-triangle"
          title="ban"
        />
      </li>
      <div class="sep" />
      <li id="submit" class="btn">
        <font-awesome-icon
          @click="newState != submission.State ? submit() : null"
          class="icn"
          icon="share"
          title="submit"
        />
      </li>
    </ul>
  </li>
</template>

<script>
import CanvasMessage from "@/components/CanvasMessage.vue";

export default {
  components: {
    CanvasMessage
  },
  props: ["token", "submission", "selectedRoomName", "selectedState"],
  data() {
    return {
      newState: this.submission.State
    };
  },
  methods: {
    submit() {
      var url =
        window.location.origin +
        "/api/?token=" +
        this.token +
        "&room_id=" +
        this.selectedRoomName +
        "&submission_id=" +
        this.submission.ID;

      if (this.newState == "submitted" || this.newState == "published") {
        url += "&method=set_submission_state&state=" + this.newState;
      } else if (this.newState == "deleted" || this.newState == "offensive") {
        url +=
          "&method=reject_submission&offensive=" +
          (this.newState == "offensive");
      }

      const options = {
        method: "GET"
      };

      fetch(url, options).then(
        function() {
          /**
           * Should probably check that the state did actually change 
           * successfully on the server. Whatever, it's non-essential. There's
           * a refresh button for now. A later patch...
           */
          this.$emit("changeState", this.newState);
        }.bind(this)
      );
    }
  }
};
</script>

<style lang="scss" scoped>
ul,
p {
  margin: $spacer 0;
}
#options {
  display: flex;
  justify-content: flex-end;
  padding: 0;
  width: 100%;
  text-align: right;

  li {
    text-align: center;
    display: inline-block;
    height: $sidebar-width/2;
    width: $sidebar-width/2;
    margin: 0 $spacer;
  }
  .icn {
    width: 100%;
    height: 100%;
    padding: $spacer/2;
  }
  li.selected {
    background: $grey-l;
    font-weight: bold;
  }

  .sep {
    border-left: 1px dashed var(--secondary-join);
    margin: 0 $spacer;
  }
}
</style>

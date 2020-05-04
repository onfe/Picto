<template>
  <li>
    <p>Author: '{{ submission.Message.Payload.Sender }}'</p>
    <CanvasMessage :msg="submission.msg" />
    <ul id="options">
      <li :class="{ btn: true, selected: newState == 'published' }">
        <font-awesome-icon
          @click="disabled ? null : (newState = 'published')"
          :class="{
            icn: true,
            disabled: disabled,
            active: submission.State == 'published'
          }"
          icon="check"
          title="publish"
        />
      </li>
      <li :class="{ btn: true, selected: newState == 'submitted' }">
        <font-awesome-icon
          @click="disabled ? null : (newState = 'submitted')"
          :class="{
            icn: true,
            disabled: disabled,
            active: submission.State == 'submitted'
          }"
          icon="hourglass"
          title="unpublish"
        />
      </li>
      <li :class="{ btn: true, selected: newState == 'deleted' }">
        <font-awesome-icon
          @click="disabled ? null : (newState = 'deleted')"
          :class="{ icn: true, disabled: disabled }"
          icon="times"
          title="reject"
        />
      </li>
      <li :class="{ btn: true, selected: newState == 'offensive' }">
        <font-awesome-icon
          @click="disabled ? null : (newState = 'offensive')"
          :class="{ icn: true, disabled: disabled }"
          icon="exclamation-triangle"
          title="ban"
        />
      </li>
      <div class="sep" />
      <li id="submit" class="btn">
        <font-awesome-icon
          @click="!disabled && newState != submission.State ? submit() : null"
          :class="{
            icn: true,
            disabled: disabled || newState == submission.State
          }"
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
  props: [
    "token",
    "submission",
    "selectedRoomName",
    "selectedState",
    "disabled"
  ],
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
           * successfully on the server and report an error?
           * For now, refreshes should keep the data presented accurate.
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

    &.selected {
      border: 1px dashed var(--secondary-join);
      font-weight: bold;
    }
  }

  .icn {
    width: 100%;
    height: 100%;
    padding: $spacer/2;

    &.active,
    &.disabled {
      color: $grey-l;
    }
  }

  .sep {
    border-left: 1px dashed var(--secondary-join);
    margin: 0 $spacer;
  }
}
</style>

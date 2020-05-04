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
        />
      </li>
      <li :class="{ btn: true, selected: newState == 'held' }">
        <font-awesome-icon
          @click="newState = 'held'"
          class="icn"
          icon="pause"
        />
      </li>
      <li :class="{ btn: true, selected: newState == 'deleted' }">
        <font-awesome-icon
          @click="newState = 'deleted'"
          class="icn"
          icon="times"
        />
      </li>
      <li id="submit" class="btn">
        <font-awesome-icon @click="submit" class="icn" icon="share" />
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
  props: ["token", "submission", "selectedRoom", "selectedState"],
  data() {
    return {
      newState: this.submission.State
    };
  },
  methods: {
    submit() {
      this.$emit("changeState", this.newState);
      console.log(this.newState);
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
  display: inline-block;
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
    text-decoration: underline;
  }
}
#submit {
  text-decoration: underline;
}
</style>

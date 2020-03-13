<template lang="html">
  <section v-bind:class="{ keyboard: true, show }">
    <div class="simple-keyboard"></div>
  </section>
</template>

<script>
import Keyboard from "simple-keyboard";
import "simple-keyboard/build/css/index.css";

export default {
  name: "SimpleKeyboard",
  props: {
    show: Boolean
  },
  data: () => ({
    keyboard: null
  }),
  mounted() {
    this.keyboard = new Keyboard({
      onChange: this.onChange,
      onKeyPress: this.onKeyPress,
      layoutName: "default",
      layout: {
        default: [
          "q w e r t y u i o p",
          "a s d f g h j k l",
          "{shift} z x c v b n m {backspace}",
          "{numbers} {space} . ,"
        ],
        shift: [
          "Q W E R T Y U I O P",
          "A S D F G H J K L",
          "{shift} Z X C V B N M {backspace}",
          "{numbers} {space} . ,"
        ],
        numbers: ["1 2 3", "4 5 6", "7 8 9", "{abc} 0 {backspace}"]
      },
      display: {
        "{numbers}": "123",
        "{space}": "space",
        "{send}": "Send",
        "{escape}": "esc ⎋",
        "{tab}": "tab ⇥",
        "{backspace}": "⌫",
        "{capslock}": "caps lock ⇪",
        "{shift}": "⇧",
        "{controlleft}": "ctrl ⌃",
        "{controlright}": "ctrl ⌃",
        "{altleft}": "alt ⌥",
        "{altright}": "alt ⌥",
        "{metaleft}": "cmd ⌘",
        "{metaright}": "cmd ⌘",
        "{abc}": "ABC"
      }
    });
  },
  methods: {
    onKeyPress(button) {
      /**
       * If you want to handle the shift and caps lock buttons
       */
      console.log(button);
      if (button === "{shift}") this.handleShift();
      else if (button === "{numbers}") this.handleNumber();
      else if (button === "{backspace}")
        this.$store.dispatch("compose/backspace");
      else if (button === "{send}") this.$store.dispatch("compose/send");
      else if (button === "{space}") this.$store.dispatch("compose/write", " ");
      else {
        this.$store.dispatch("compose/write", button);
        this.handleShift(false);
      }
    },
    handleShift(force) {
      let currentLayout = this.keyboard.options.layoutName;
      var shiftToggle;
      if (force === undefined) {
        shiftToggle = currentLayout === "default" ? "shift" : "default";
      } else {
        shiftToggle = force ? "shift" : "default";
      }

      this.keyboard.setOptions({
        layoutName: shiftToggle
      });
    },
    handleNumber() {
      console.log("number");
    }
  },
  watch: {
    input(input) {
      this.keyboard.setInput(input);
    }
  }
};
</script>

<style lang="scss" scoped>
.keyboard {
  max-height: 0;
  overflow: scroll;

  &.show {
    max-height: 100vh;
    overflow: visible;
  }

  transition: max-height 300ms ease;
}
</style>

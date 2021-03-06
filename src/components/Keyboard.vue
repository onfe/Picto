<template lang="html">
  <section v-bind:class="{ show }">
    <div id="picto-keyboard" class="simple-keyboard"></div>
  </section>
</template>

<script>
import Keyboard from "simple-keyboard";
import "@/assets/scss/keyboard.scss";

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
      theme: "picto-keyboard",
      layout: {
        default: [
          "q w e r t y u i o p",
          "a s d f g h j k l",
          "{shift} z x c v b n m {backspace}",
          "{alt} {space} , ."
        ],
        shift: [
          "Q W E R T Y U I O P",
          "A S D F G H J K L",
          "{shift} Z X C V B N M {backspace}",
          "{alt} {space} , ."
        ],
        specials: [
          "1 2 3 4 5 6 7 8 9 0",
          "@ # £ _ & - + ( ) /",
          "{specials2} * \" ' : ; ! ? {backspace}",
          "{abc} {space} , ."
        ],
        specials2: [
          "~ ` | • < > ÷ × ¶ ?",
          "€ ¥ $ ¢ ^ ° = { } \\",
          "{specials} % © ® [ ] ¡ ¿ {backspace}",
          "{abc} {space} , ."
        ]
      },
      buttonTheme: [
        {
          class: "small",
          buttons: ". , {abc} {alt}"
        }
      ],
      display: {
        "{specials}": "123",
        "{alt}": "123",
        "{specials2}": "=\\<",
        "{space}": "SPACE",
        "{backspace}": "DEL",
        "{shift}": "SHIFT",
        "{abc}": "abc"
      }
    });
  },
  methods: {
    onKeyPress(button) {
      if (button === "{shift}") this.handleShift();
      else if (button === "{alt}") this.handleLayoutChange("specials");
      else if (button === "{specials}") this.handleLayoutChange("specials");
      else if (button === "{specials2}") this.handleLayoutChange("specials2");
      else if (button === "{abc}") this.handleLayoutChange("default");
      else if (button === "{backspace}")
        this.$store.dispatch("compose/backspace");
      else if (button === "{send}") this.$store.dispatch("compose/send");
      else if (button === "{space}") this.$store.dispatch("compose/write", " ");
      else {
        this.$store.dispatch("compose/write", button);
        if (this.keyboard.options.layoutName == "shift") {
          this.handleShift(false);
        }
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
    handleLayoutChange(layout) {
      this.keyboard.setOptions({
        layoutName: layout
      });
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
section {
  max-height: 0vh;
  overflow: hidden;

  &.show {
    max-height: 15em;
  }

  transition: max-height 400ms ease-in-out;
}
</style>

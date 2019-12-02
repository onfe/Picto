<template>
  <section class="toolbox">
    <ul class="tool">
      <li v-bind:class="{ selected: isPencil, rainbow: isRainbow }">
        <font-awesome-icon @click="pencil" class="icn" icon="pencil-alt" />
      </li>
      <li v-bind:class="{ selected: isEraser }">
        <font-awesome-icon @click="eraser" class="icn" icon="eraser" />
      </li>
    </ul>

    <ul class="size">
      <li v-bind:class="{ selected: isSmall }">
        <font-awesome-icon
          @click="small"
          class="icn"
          :icon="['far', 'dot-circle']"
        />
      </li>
      <li v-bind:class="{ selected: isLarge }">
        <font-awesome-icon @click="large" class="icn" icon="circle" />
      </li>
    </ul>
    <ul class="keyboard">
      <li>
        <font-awesome-icon @click="keyboard" class="icn" icon="keyboard" />
      </li>
    </ul>
    <input autocapitalize="none" id="text-input" />
  </section>
</template>

<script>
export default {
  computed: {
    isPencil() {
      return this.$store.state.compose.tool == "pencil";
    },
    isEraser() {
      return this.$store.state.compose.tool == "eraser";
    },
    isLarge() {
      return this.$store.state.compose.size == "large";
    },
    isSmall() {
      return this.$store.state.compose.size == "small";
    },
    isRainbow() {
      return this.$store.state.compose.rainbow;
    }
  },
  methods: {
    pencil() {
      this.$store.dispatch("compose/pencil");
    },
    eraser() {
      this.$store.dispatch("compose/eraser");
    },
    small() {
      this.$store.dispatch("compose/small");
    },
    large() {
      this.$store.dispatch("compose/large");
    },
    keyboard() {
      const el = document.getElementById("text-input");
      el.focus();
    },
    handleInput(e) {
      e.preventDefault();
      e.stopPropagation();
      const el = document.getElementById("text-input");
      if (el.value.length < 1) {
        this.$store.dispatch("compose/backspace");
      } else if (el.value[1] != "℗") {
        const chr = el.value[1];
        this.$store.dispatch("compose/write", chr);
      }
      el.value = "℗";
    }
  },
  mounted() {
    const el = document.getElementById("text-input");
    el.value = "℗";
    el.addEventListener("input", this.handleInput);
  },
  beforeDestroy() {
    const el = document.getElementById("text-input");
    el.removeEventListener("input", this.handleInput);
  }
};
</script>

<style lang="scss" scoped>
section {
  display: flex;
  flex-direction: column;
  background: #fff;
  padding: $spacer;
}

ul {
  border-radius: $spacer;
  overflow: hidden;
  list-style-type: none;
  padding: 0;
  margin: 0;
  position: relative;
  background: $almost-white;
  margin-bottom: 1vw;

  &:last-of-type {
    margin: 0;
  }

  li {
    width: 100%;
    padding-bottom: 100%;
    position: relative;
    border-bottom: 1px solid #fff;

    &:last-of-type {
      border: 0;
    }

    .icn {
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      padding: $spacer;
    }

    &.selected {
      background: $grey-l;
    }

    &.rainbow .icn {
      animation: rainbowbg 8s infinite;
    }
  }
}

.keyboard {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
}

#text-input {
  width: 1px;
  height: 1px;
  opacity: 0;
  position: absolute;
}

@include rainbow("rainbowbg", "color");
</style>

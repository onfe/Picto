<template lang="html">
  <ToggleButton
    class="darkMode"
    icon="moon"
    iconActive="sun"
    :active="theme == 'dark'"
    @click.native="setTheme(theme == 'dark' ? 'light' : 'dark')"
  />
</template>

<script>
import ToggleButton from "@/components/ToggleButton.vue";
export default {
  components: {
    ToggleButton
  },
  data() {
    return {
      theme: "light"
    };
  },
  methods: {
    setTheme(theme) {
      this.theme = theme;
      document.getElementsByTagName("html")[0].setAttribute("class", theme);
      this.$cookies.set("theme", this.theme);
      console.log(this.$cookies.get("theme"));
    }
  },
  mounted() {
    //If there is a theme cookie set it takes priority.
    if (this.$cookies.isKey("theme")) {
      this.setTheme(this.$cookies.get("theme"));
    } else {
      //Otherwise we default to prefers-color-scheme
      if (window.matchMedia("(prefers-color-scheme: dark)").matches) {
        this.setTheme("dark");
      } //Don't need to check prefers-color-scheme:light as light theme is default.
    }
  }
};
</script>

<template lang="html">
  <ul class="btn">
    <li @click="setTheme(nextTheme)">
      <font-awesome-icon class="icn" :icon="themeIcon" />
    </li>
  </ul>
</template>

<script>
export default {
  data() {
    return {
      theme: "light",
      themes: ["light", "dark", "pink"]
    };
  },
  computed: {
    nextTheme() {
      return this.themes[
        (this.themes.indexOf(this.theme) + 1) % this.themes.length
      ];
    },
    themeIcon() {
      switch (this.theme) {
        case "light":
          return "sun";
        case "dark":
          return "moon";
        case "pink":
          return "ice-cream";
      }
      return "sun";
    }
  },
  methods: {
    setTheme(theme) {
      this.theme = theme;
      document.getElementsByTagName("html")[0].setAttribute("class", theme);
      if (window.matchMedia("(prefers-color-scheme: dark)").matches && theme == "dark"
      || !window.matchMedia("(prefers-color-scheme: dark)").matches && theme == "light") {
        this.$cookies.remove("theme")
      } else {
        this.$cookies.set("theme", this.theme);
      }
    }
  },
  mounted() {
    //If there is a theme cookie set it takes priority.
    if (this.$cookies.isKey("theme")) {
      this.setTheme(this.$cookies.get("theme"));
    } else {
      //Otherwise we default to prefers-color-scheme and don't set a cookie.
      if (window.matchMedia("(prefers-color-scheme: dark)").matches) {
        this.theme = "dark";
        document.getElementsByTagName("html")[0].setAttribute("class", "dark");
      } //Don't need to check prefers-color-scheme:light as light theme is default.
    }
  }
};
</script>

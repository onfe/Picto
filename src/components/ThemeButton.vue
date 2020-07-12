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
      theme: "system",
      themes: ["system", "light", "dark", "pink"]
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
      return "paint-roller";
    }
  },
  methods: {
    setTheme(theme) {
      this.theme = theme;
      document.getElementsByTagName("html")[0].setAttribute("class", theme);
      this.$cookies.set("theme", this.theme);
    }
  },
  mounted() {
    //If there is a theme cookie set, we apply that theme.
    if (this.$cookies.isKey("theme") && this.$cookies.get("theme") != "system") {
      this.setTheme(this.$cookies.get("theme"));
    }
  }
};
</script>

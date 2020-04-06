<template lang="html">
  <div class="update-man" :title="hoverTitle">
    <button @click="handleClick" v-if="updateable" class="update">
      <font-awesome-icon
        class="icn pad"
        :class="{ spin: updating }"
        icon="sync-alt"
      />{{ updating ? "Updating" : "Update" }}
    </button>
    <div v-else class="version">
      <font-awesome-icon class="icn pad" :icon="icon" />
      <a :href="link">{{ version }}</a>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    version: {
      type: String,
      default: "v?.?.?"
    },
    link: {
      type: String,
      default: "#"
    },
    status: {
      type: String,
      default: "normal"
    }
  },
  data() {
    return {
      updating: false
    };
  },
  computed: {
    updateable() {
      return this.status === "update-ready";
    },
    preparing() {
      return this.status === "update-preparing";
    },
    icon() {
      return this.preparing ? "download" : "code-branch";
    },
    hoverTitle() {
      if (this.preparing) {
        return "Preparing update";
      } else if (this.updatable) {
        return "Click to update";
      } else {
        return null;
      }
    }
  },
  methods: {
    handleClick() {
      this.$emit("update");
      this.updating = true;
    }
  }
};
</script>

<style lang="scss" scoped>
.update-man {
  display: flex;
  font-size: inherit;
  justify-content: center;

  .icn {
    margin-right: 1ch;
  }

  font-size: inherit;
  font-family: inherit;
}

.update {
  font-family: inherit;
  background: $green-d;
  cursor: pointer;
  padding: 0.25em;
  margin: -0.25em;
  border: 0;
  border-radius: 0.25em;
  color: #fff;

  &:focus {
    outline: none;
  }

  .icn.spin {
    animation: spin 1s infinite linear;
  }
}

.version {
  a {
    color: inherit;
    text-decoration-style: dotted;
    transition: color 200ms ease-in-out;

    &:hover {
      color: $grey-d;
    }
  }
}

@keyframes spin {
  from {
    transform: scale(1) rotate(0deg);
  }
  to {
    transform: scale(1) rotate(360deg);
  }
}
</style>

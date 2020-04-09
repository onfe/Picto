<template>
  <div class="moderate">
    <div class="container">
      <h1>Moderation dashboard</h1>

      <hr />

      <AuthForm v-if="token === null" @authenticated="setToken" />

      <p v-if="token != null">Select a room:</p>
      <div v-if="token != null" class="room-selector">
        <select v-model="gallery" class="font-pixel" @click="updateRooms">
          <option
            v-for="gallery in galleries"
            v-bind:key="gallery.Name"
            v-bind:value="gallery"
          >
            {{ gallery.Name }} ({{ gallery.Pop }}/{{ gallery.Cap }}) -
            {{
              gallery.Submissions +
                (gallery.Submissions == 1 ? " submission" : " submissions")
            }}
          </option>
        </select>
      </div>

      <p v-if="gallery != null">{{ gallery.Desc }}</p>

      <SubmissionList
        v-if="gallery != null"
        class="submission-list"
        :token="token"
        :roomName="gallery.Name"
      />
    </div>
  </div>
</template>

<script>
import AuthForm from "@/components/AuthForm.vue";
import SubmissionList from "@/components/SubmissionList.vue";

export default {
  name: "moderate",
  components: {
    AuthForm,
    SubmissionList
  },
  data() {
    return {
      token: null,
      galleries: [],
      gallery: null
    };
  },
  methods: {
    setToken(token) {
      this.token = token;
      this.updateRooms();
    },
    updateRooms() {
      const url =
        window.location.origin +
        `/api/?method=get_submission_rooms&token=${this.token}`;
      const options = {
        method: "GET"
      };

      setTimeout(() => {
        fetch(url, options)
          .then(resp => {
            return resp.text();
          })
          .then(result => {
            this.galleries = JSON.parse(result) || [];
          });
      }, 250);
    }
  }
};
</script>

<style lang="scss" scoped>
.moderate,
.container {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.moderate {
  background-color: var(--background-join);
  color: var(--primary-join);

  background-image: url("/img/stripe.svg");
  background-repeat: repeat-y;
  background-position-x: 0.8rem;

  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.container {
  padding: 0 1.5rem 1rem 3.5rem;

  @media (min-width: 992px) {
    padding-left: 8rem;
  }

  font-family: monospace;
  font-size: 1.2rem;
  color: var(--primary-join);
}

hr {
  flex: 0;
  border: 0;
  width: 100%;
  height: 1px;
  border-bottom: 1px dashed var(--secondary-join);
}

p {
  margin-bottom: 1.5rem;
  line-height: 1.2;
}

.room-selector {
  background: $grey-d;
  transition: background 200ms ease-in-out;
  padding: 0.5rem;
  max-width: 675px;
}

select {
  flex: 1;

  width: 100%;
  min-width: 0;

  font-size: 2rem;
  background-color: transparent;

  padding: 0.25rem 0.5rem;
  border: 0;
  background: #fff;

  border-radius: 0;
  -webkit-appearance: none;
  -moz-appearance: none;

  &:focus {
    background-color: white;
    outline: none;
  }
}
</style>

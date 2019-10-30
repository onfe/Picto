const axs = require("axios");

const axios = axs.create({
  baseURL: window.location.protocol + window.location.host
});

export default axios

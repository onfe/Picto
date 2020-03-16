import Vue from "vue";

import { library } from "@fortawesome/fontawesome-svg-core";
import {
  faPencilAlt,
  faEraser,
  faCircle,
  faKeyboard,
  faInfoCircle,
  faUserPlus,
  faTimesCircle
} from "@fortawesome/free-solid-svg-icons";
import { faDotCircle } from "@fortawesome/free-regular-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

// Add the required icons to keep the bundle small.
library.add(
  faPencilAlt,
  faEraser,
  faCircle,
  faDotCircle,
  faKeyboard,
  faInfoCircle,
  faUserPlus,
  faTimesCircle
);

Vue.component("font-awesome-icon", FontAwesomeIcon);

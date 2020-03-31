import Vue from "vue";

import { library } from "@fortawesome/fontawesome-svg-core";
import {
  faPencilAlt,
  faEraser,
  faCircle,
  faKeyboard,
  faInfoCircle,
  faUserPlus,
  faTimesCircle,
  faDoorOpen,
  faBug,
  faHeart,
  faCodeBranch,
  faCopy
} from "@fortawesome/free-solid-svg-icons";
import { faDotCircle } from "@fortawesome/free-regular-svg-icons";
import { faTwitter } from "@fortawesome/free-brands-svg-icons";
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
  faTimesCircle,
  faDoorOpen,
  faBug,
  faHeart,
  faTwitter,
  faCodeBranch,
  faCopy
);

Vue.component("font-awesome-icon", FontAwesomeIcon);

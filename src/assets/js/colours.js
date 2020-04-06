import Color from "color";

const COLOURS = [
  "#E97777",
  "#008DAB",
  "#FFC694",
  "#89C7B6",
  "#EFC57E",
  "#AD84C7",
  "#7998C9",
  "#CF8EA4"
];

function colour(n) {
  /*
    Returns a hex colour that will always be the same for a given n.
    The first 3 colours are set for branding purposes.
    n is zero indexed.
  */
  if (n < 3) {
    return COLOURS[n];
  } else {
    var hue = Math.floor(((360 / (Math.E / 2)) * n) % 360);
    return Color.hsl([hue, 72, 69]).hex();
  }
}

export default colour;

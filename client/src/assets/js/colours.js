import Color from "color"

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
    The first 8 colours are set to ensure they look good.
    n is zero indexed.
  */
  if (n < 8) {
    return COLOURS[n];
  } else {
    var hue = sprng(7930832, n - 8) * 360
    return Color.hsl([hue, 90, 65]).hex();
  }
}

function sprng(seed, n) {
  /* calculate the nth seeded float between 0-1.
   * both seed and n must be an integer.
   * Runs in O(n)
   */

  // normalise n >= 1
  n = Math.abs(n) + 1;

  // normalise the seed
  seed = seed % 2147483647;
  seed = seed <= 0 ? seed + 2147483646 : seed;

  for (var i = 0; i <= n; i++) {
    seed = seed * 16807 % 2147483647;
  }
  return Math.abs(seed - 1) / 2147483646;
}



export default colour;

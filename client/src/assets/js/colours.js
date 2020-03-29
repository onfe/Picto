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

function GET_COLOUR(index) {
  if (index < 3) {
    return COLOURS[index];
  }
  return (
    "hsl(" + Math.floor(((255 / (Math.E / 2)) * index) % 255) + ",90%,65%)"
  );
}

export default GET_COLOUR;

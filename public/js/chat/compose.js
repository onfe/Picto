var width = 192; // Must maintain aspect.
var height = 64;
var strokeWidth;

var perX = view.size.width / width;
var perY = view.size.height / height;

var perAvg = (perX + perY) / 2

var lastPoint; // for point interpolation.

var pixels = []; // setup pixels.

for (var x = 0; x < width; x++) {
  pixels.push([]);
  for (var y = 0; y < height; y++) {
    pixels[x].push(null);
  }
}

function getPixelCoords(x, y) { // input point coords.
  var pX = Math.floor(x / perX);
  var pY = Math.floor(y / perY);
  return [pX, pY]; // return pixel coords.
}

function drawPixel(x, y) {
  if (pixels[x][y]) { // if pixel already drawn at this location, don't draw it again.
    return;
  }
  var sPointX = Math.floor(x * perX);
  var sPointY = Math.floor(y * perY);
  var sizeX = Math.ceil(perX);
  var sizeY = Math.ceil(perY);
  var pixel = Path.Rectangle(new Point(sPointX, sPointY), new Size(sizeX, sizeY))
  pixel.fillColor = 'black';
  pixels[x][y] = pixel;
}

function erasePixel(x, y) {
  if (!pixels[x][y]) {
    return;
  }
  pixels[x][y].remove(); // remove the pixel from the canvas.
  pixels[x][y] = null; // ensure array location is empty.
}

function pencil(x, y, stroke) {
  if (stroke > 1) {
    var sX = sY = Math.floor(stroke / 2);
    for (var pX = 0; pX < stroke; pX++) {
      for (var pY = 0; pY < stroke; pY++) {
        drawPixel((x - sX) + pX, (y - sY) + pY);
      }
    }
  } else {
    drawPixel(x, y);
  }
}

function eraser(x, y, stroke) {
  if (stroke > 1) {
    var sX = sY = Math.floor(stroke / 2);
    for (var pX = 0; pX < stroke; pX++) {
      for (var pY = 0; pY < stroke; pY++) {
        erasePixel((x - sX) + pX, (y - sY) + pY);
      }
    }
  } else {
    erasePixel(x, y);
  }
}

function onMouseDown(event) {

  console.log(currentSize, currentTool)

  if (currentSize == 'big') {
    strokeWidth = 3;
  } else if (currentSize == 'small') {
    strokeWidth = 1;
  }

  lastPoint = event.point;
  var pCoords = getPixelCoords(event.point.x, event.point.y);

  if (currentTool == 'pencil') {
    pencil(pCoords[0], pCoords[1], strokeWidth);
  } else if (currentTool == 'eraser') {
    eraser(pCoords[0], pCoords[1], strokeWidth);
  }
}

function onMouseDrag(event) {

  var pCoords = getPixelCoords(event.point.x, event.point.y);

  var changeVector = event.point - lastPoint;

  var inc = Math.floor((perAvg * strokeWidth) / 2);
  if (inc <= 0) {
    inc = 1;
  }

  for (var i = 0; i < changeVector.length; i += inc) {
    var pc = i / changeVector.length;
    var a = changeVector * pc
    var p = lastPoint + a
    pCoords = getPixelCoords(p.x, p.y);

    if (currentTool == 'pencil') {
      pencil(pCoords[0], pCoords[1], strokeWidth);
    } else if (currentTool == 'eraser') {
      eraser(pCoords[0], pCoords[1], strokeWidth);
    }

  }


  lastPoint = event.point;
}

// function onMouseUp(event) {
//   // Nothing here
// }
//
// function onResize(event) {
//   var perX = view.size.width / width;
//   var perY = view.size.height / height;
//   var perAvg = (perX + perY) / 2
// }

window.clearCanvas = function () {
  for (var x = 0; x < pixels.length; x++) {
    if (Array.isArray(pixels[x])) {
      for (var y = 0; y < pixels[x].length; y++) {
        if (pixels[x][y]) {
          pixels[x][y].remove();
          pixels[x][y] = null;
        }
      }
    }
  }
}

window.getContent = function () {
  var outData = [];
  // scan across then down.
  for (var y = 0; y < height; y++) {
    outData.push('');
    for (var x = 0; x < width; x++) {
      if (pixels[x][y]) {
        outData[y] += '1';
      } else {
        outData[y] += '0';
      }
    }
  }
  return outData;
}

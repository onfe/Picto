var width = 192; // Must maintain aspect.
var height = 64;

function drawMsg(data, msgid) {
  console.log(msgid);

  var msgPaper = new paper.PaperScope();

  msgPaper.setup(msgid);
  with (msgPaper) {
    var perX = view.size.width / width;
    var perY = view.size.height / height;


    for (y = 0; y < height; y++) {
      for (x = 0; x < width; x++) {
        var dataIndex = (y * width) + x;
        if (data[dataIndex] == '1') {
          draw(x, y, msgPaper, perX, perY)
          console.log(x,y)
        }
      }
    }
  }
}

function draw(x, y, pp, px, py) {
  var sPointX = Math.floor(x * px);
  var sPointY = Math.floor(y * py);
  var sizeX = Math.ceil(px);
  var sizeY = Math.ceil(py);
  var pixel = pp.Path.Rectangle(new pp.Point(sPointX, sPointY), new pp.Size(sizeX, sizeY))
  pixel.fillColor = 'black';
}

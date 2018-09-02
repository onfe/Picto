function getState(x, y) {
  var state = $(`#g-btn-${x}-${y}`).hasClass('active');
  return state;
}

$('#add').on('click', function() {
  var width = $('#width').val();
  var height = 7;
  var lines = [];
  for (let y = 0; y < height; y++) {
    var line = ''
    for (let x = 0; x < width; x++) {
      var state = getState(x, y);
      if (state) {
        line += '#';
      } else {
        line += '.'
      }
    }
    lines.push(line);
  }
  console.log(lines)
  var char = $('#char').val();
  var output = {
    char: char,
    width: width,
    height: height,
    data: lines
  }

  $('#output').append(JSON.stringify(output, null, 2) + ',\n')
});

$('#gridarea div').on('click', function(e) {
  console.log(e)
  $(e.target).toggleClass('active');
});

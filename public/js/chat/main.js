var currentTool = 'pencil';
var currentSize = 'small';
$('#' + currentTool).addClass('selected');
$('#' + currentSize).addClass('selected');

$('.tool.draw').on('click', function () {
  $('.tool.draw').removeClass('selected');
  $(this).addClass('selected');
  currentTool = $(this).attr('id')
});

$('.tool.size').on('click', function () {
  $('.tool.size').removeClass('selected');
  $(this).addClass('selected');
  currentSize = $(this).attr('id')
});

$('.control.clear').on('click', function () {
  clearCanvas(); // from compose.js
})

$('.control.send').on('click', function () {
  var msgCont = getContent(); // from compose.js

  var emptyLine = 'A'.repeat(32);

  for (var line = 0; line < msgCont.length; line++) {
    var currLine = msgCont[line];
    var encLine = base64Encode(currLine);

    msgCont[line] = encLine;

    if (msgCont[line] == emptyLine) {
      msgCont[line] = '!';
      // to remove empty space, replace 'AAA...A' (an empty line) with !.
    }
  }

  // We now have a nicely array with base64 (slimmed edition) encoded values.
  msgCont = msgCont.join('');


  console.log(msgCont);

  sendMessage(msgCont) // from socket.js
})

function base64Encode(st) {
  var out = '';
  var base64List = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-=';
  var chunks = chunkString(st, 6);
  for (var c = 0; c < chunks.length; c++) {
    var num = parseInt(chunks[c], 2) // Binary string to decimal integer.
    var char = base64List[num];
    out += char;

  }
  return out;
}

function chunkString(str, length) {
  return str.match(new RegExp('.{1,' + length + '}', 'g'));
}

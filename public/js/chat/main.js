var currentTool = 'pencil';
var currentSize = 'small';
var emptyLine = 'A'.repeat(32);

var username = Cookies.get("user");
if (!username) {
  location.href = "/";
}
Cookies.remove("user")

$(window).on('load', function () { // when everything has loaded, hide the loader and show the chat.
  $('.loader').addClass('hide');

});

$(function () {
  $('#' + currentTool).addClass('selected');
  $('#' + currentSize).addClass('selected');
  $('.msg-auth.self').html(username)
});

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

  sendMessage(msgCont) // from socket.js
})

function joined(colour) {
  $('.msg-current').addClass(colour)
}

function messageRecieved(pl) {
  var sender = pl.sender;
  var content = pl.msgCont;
  var colour = pl.colour;
  var msgID = pl.msgID;

  var expanded = ''

  for (var i = 0; i < content.length; i++) {
    var char = content[i]
    if (char === '!') {
      expanded += emptyLine;
    } else {
      expanded += char;
    }
  }

  var msgBin = base64Decode(expanded);

  var scrollBefore = $('#msg-hist').scrollTop()

  $('.msg-history').prepend('<div id="msg-' + msgID + '"class="msg ' + colour + '" data-pixel="' + expanded + '"><canvas resize id="canv-' + msgID + '"></canvas><span class="msg-auth">' + sender + '</span></div>')
  setTimeout(function () {
    drawMsg(msgBin, 'canv-' + msgID)
  }, 5); // wait for the DOM to update and size the canvas before drawing to it.

  var scrollAfter = $('#msg-hist').scrollTop()

  $('#msg-hist').scrollTop(scrollBefore) // reset the autoscroll

  $('#msg-hist').animate({scrollTop: scrollAfter}, 300)

}

function msgSent(pl) {
  isSending = false;
  clearCanvas(); // fom compose.js
}

function base64Decode(st) {
  var out = '';
  var base64List = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-=';
  for (var i = 0; i < st.length; i++) {
    var char = st[i];
    var num = base64List.indexOf(char);
    var bin = num.toString(2);
    var final = ('000000' + bin).slice(-6); // ensure is 6 chars
    out += final;
  }
  return out;
}

function base64Encode(st) {
  var out = '';
  var base64List = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-=';
  var chunks = chunkString(st, 6);
  for (var c = 0; c < chunks.length; c++) {
    var num = parseInt(chunks[c], 2) // Binary string to decimal integer.
    var char = base64List[num];
    out += char.toString(2);

  }
  return out;
}

function chunkString(str, length) {
  return str.match(new RegExp('.{1,' + length + '}', 'g'));
}

var path = location.pathname.endsWith('/') ? location.pathname.slice(0, -1) : location.pathname;
path = path.split('/')
var roomCode = path[path.length - 1]

$(function () {

  $('#room-id').html(roomCode);

  $.ajax({
    url: 'http://' + window.location.host + '/api/room/',
    data: {room: roomCode},
  }).done( room );

  $('#join').on('click', function() {
    // Get code from input box
    var name = $('#username').val();

    // Perform server side initial validation
    $.ajax({
      url: 'http://' + window.location.host + '/api/username/',
      data: {room: roomCode, name: name},
    }).done( validateName );
  });
});

function room(e) {
  if (!e.available) {
    location.href = '/';
  }
  $('#room-clients').html(e.numCli)
}

function validateName(e) {
  if (e.available) {
    Cookies.set('user', e.name)
    location.href = '/room/' + roomCode + '/';
  };
};

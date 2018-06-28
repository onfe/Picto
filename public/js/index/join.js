var path = location.pathname.split('/')
var roomCode = path[path.length - 2]
console.log(roomCode)

$(function () {

  $("#room-id").html(roomCode);

  $.ajax({
    url: "http://" + window.location.host + "/api/room/",
    data: {room: roomCode},
  }).done( room );

  $("#join").on('click', function() {
    // Get code from input box
    var name = $("#username").val();

    // Perform server side initial validation
    $.ajax({
      url: "http://" + window.location.host + "/api/username/",
      data: {room: roomCode, name: name},
    }).done( validateName );
  });
});

function room(e) {
  $('#room-clients').html(e.numCli)
}

function validateName(e) {
  if (e.available) {
    Cookies.set("user", e.name)
    location.href = "/room/" + roomCode + "/";
  };
};

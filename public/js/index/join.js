var path = location.pathname.split('/')
var roomCode = path[path.length - 2]
console.log(roomCode)

$(function () {

  $("#roomid").html(roomCode);

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

function validateName(e) {
  if (e.available) {
    Cookies.set("user", e.name)
    location.href = "/room/" + roomCode + "/";
  };
};

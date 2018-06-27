var path = location.pathname.split('/')
var roomCode = path[path.length - 1]
console.log(roomCode)


$(function () {

  $.ajax({
    url: "http://" + window.location.host + "/api/stats/",
    data: {}
  }).done( stats );

  $("#join").on('click', function() {
    // Get code from input box
    var code = $("#join-code").val();

    // Perform server side initial validation
    $.ajax({
      url: "http://" + window.location.host + "/api/room/",
      data: {room: code},
    }).done( validateJoin );
  });

  $("#create").on('click', function() {
    console.log('notbroken')
    // Perform server side initial validation
    $.ajax({
      url: "http://" + window.location.hostname + ":" + window.location.port + "/api/createroom/",
      data: {id: false},
    }).done( validateCreate );
  });
});

function validateJoin(e) {
  if (!e.available) {
    return
  }
  console.log('ok')
  location.href = "/join/" + e.room + "/";
};

function stats(e) {
  $('#current-clients').html(e.numClients);
  $('#current-rooms').html(e.numRooms)
}

function validateCreate(e) {
  if (!e.created) {
    return
  }
  console.log('ok')
  location.href = "/join/" + e.room + "/";
}

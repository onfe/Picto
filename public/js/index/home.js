$(function () {
  $("#join").on('click', function() {
    // Get code from input box
    var code = $("#joincode").val();

    // Perform server side initial validation
    $.ajax({
      url: "http://" + window.location.hostname + ":" + window.location.port + "/api/room/",
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
  location.href = "/room/" + e.room + "/";
};

function validateCreate(e) {
  if (!e.created) {
    return
  }
  console.log('ok')
  location.href = "/room/" + e.room + "/";
}

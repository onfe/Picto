$(function () {
  $("#join").on('click', function() {
    // Get code from input box
    var code = $("#joincode").val();

    // Perform server side initial validation
    $.ajax({
      url: 'http://127.0.0.1:8000/api/',
      data: {roomCode: code},
    }).done( validateSuccess );


  });
});

function validateSuccess(e) {
  console.log(e)
  if (!e.available) {
    return
  }
  console.log('ok')
  // location.href = "/room/" + e.roomCode + "/";
};

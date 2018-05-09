var ws = new WebSocket('ws://127.0.0.1:40510');

    // event emmited when connected
    ws.onopen = function () {
        console.log('websocket is connected ...')

        // sending a send event to websocket server
        ws.send('connected')
    }

    // event emmited when receiving message
    ws.onmessage = function (ev) {
        console.log(ev);
    }

    var sendButton = document.getElementById('send');
    var msgBox = document.getElementById('msg');

    sendButton.onclick = function(e) {
      ws.send(msgBox.value);
      msgBox.value = '';
    };

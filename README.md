# Picto
Picto is a web-chat app inspired by Nintendo's PictoChat, running on Node.js  
The project aims to bring the joy of Nintendo's PictoChat to any device capable
of running a modern web browser (HTML5 & JS ES6).

**Picto is optimised for mobile screens.** Desktop optimisation.. soon™

## Using Picto
Using a mouse or finger, your scribbles are translated into pixel art using JavaScript.  
By interacting directly with the canvas API (no library) drawing is kept fluid and smooth.

Keyboard and stamps are coming soon.

### Invites
Friends can be invited to your room with a single link, no sign-up required.

### Transmission
Using the WebSocket API on the client and the ws plugin for Node.js,
drawings are encoded as text and sent to the server for distribution to other members of the room.

## Building Picto
Install Node.js, npm and browserify.

Install dependencies with `npm install`.

Build the client with browserify.
```
cd public/js/chat/
browserify main.js -o pictoclient.bundle.js
```

Then, just type `npm start`.

## Credits
Code by [Onfe](https://www.onfe.co.uk)  
Icons by [Fontawesome](https://fontawesome.com/)

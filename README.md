# Videostreaming-from-web-browser-to-other-people-in-WebSocket
 Videostreaming-from-web-browser-to-other-people-in-WebSocket

Using WebRTC I receive an image from a webcam. Next, I record this video in the browser. Each piece lasts 2 seconds. I read and convert the received Blob data into Base64 and send it as JSON via WebSocket to the server. There I send them to all listeners and in the client’s browser I convert the resulting Base64 string into a Blob, then read the data from the Blob and add it to the playback buffer in the player.

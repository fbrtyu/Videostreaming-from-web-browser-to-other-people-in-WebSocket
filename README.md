# Videostreaming-from-web-browser-to-other-people-in-WebSocket
 Videostreaming-from-web-browser-to-other-people-in-WebSocket

Using Media Capture and Streams API I receive an image from a webcam. Next, I record this video in the browser. Each piece lasts 2 seconds. I read and convert the received Blob data into Base64 and send it as JSON via WebSocket to the server. There I send them to all listeners and in the client’s browser I convert the resulting Base64 string into a Blob, then read the data from the Blob and add it to the playback buffer in the player.

Используя Media Capture and Streams API, я получаю изображение с веб-камеры. Далее я записываю это видео в браузере. Каждая часть длится 2 секунды. Я читаю и конвертирую полученные данные Blob в Base64 и отправляю их в формате JSON через WebSocket на сервер. Там я рассылаю их всем слушателям и в браузере клиента конвертирую полученную строку Base64 в Blob, затем считываю данные из Blob и добавляю их в буфер воспроизведения в плеере.


let ws = new WebSocket("ws://localhost:" + global.backendPort + "/web/app/events");

ws.onmessage = (message) => {
    let obj = JSON.parse(message.data);

    // event name
    console.log(obj.event);

    // event data
    console.log(obj.AtrNameInFrontend);
}

const socket = new WebSocket('wss://localhost:8080/ws');

let connect = () => {
    console.log('Attempting connection...')

    socket.onopen = () => {
        console.log('Connection successful!')
    }

    socket.onmessage = msg => {
        console.log(msg)
    }
    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
    }

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    }
}

let sendMsg = msg => {
    console.log('sending message: ', msg)
    socket.send(msg)
}

export {connect, sendMsg}
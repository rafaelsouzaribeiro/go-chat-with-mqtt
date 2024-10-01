var hostname = "127.0.0.1";  
var port = 9090; 
var clientId = "webio4mqttexample" + new Date().getUTCMilliseconds();
var username = "root";
var password = "123mudar";
var subscription = "topic/test";

var mqttClient = new Paho.MQTT.Client(hostname, port, clientId);
mqttClient.onMessageArrived = MessageArrived;
mqttClient.onConnectionLost = ConnectionLost;


Connect();

function Connect() {
    mqttClient.connect({
        onSuccess: Connected,
        onFailure: ConnectionFailed,
        keepAliveInterval: 10,
        userName: username,
        useSSL: false,
        password: password
    });
}


function Connected() {
    console.log("Connected");
    mqttClient.subscribe(subscription);
}


function ConnectionFailed(res) {
    console.log("Connect failed: " + res.errorMessage);
}


function ConnectionLost(res) {
    if (res.errorCode !== 0) {
        console.log("Connection lost: " + res.errorMessage);
        Connect();
    }
}


function MessageArrived(message) {
    console.log("Mensagem recebida no tópico " + message.destinationName + " : " + message.payloadString);
}


function sendMessage() {
    var jsonMessage = {
        "username": "User123",
        "message": "f",
        "userId": "1"
    };

    var payload = JSON.stringify(jsonMessage);
    
    
    var message = new Paho.MQTT.Message(payload);
    message.destinationName = subscription;

    mqttClient.send(message);

    console.log("Mensagem enviada: " + payload);
}
var hostname = "127.0.0.1";  // Atualize para o IP correto
var port = 9001;  // Porta WebSocket do broker MQTT
var clientId = "webio4mqttexample" + new Date().getUTCMilliseconds();
var username = "root";
var password = "123mudar";
var subscription = "topic/test";

var mqttClient = new Paho.MQTT.Client(hostname, port, clientId);
mqttClient.onMessageArrived = MessageArrived;
mqttClient.onConnectionLost = ConnectionLost;

// Conecta ao broker MQTT
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

// Conexão bem-sucedida
function Connected() {
    console.log("Connected");
    mqttClient.subscribe(subscription);
}

// Conexão falhou
function ConnectionFailed(res) {
    console.log("Connect failed: " + res.errorMessage);
}

// Conexão perdida
function ConnectionLost(res) {
    if (res.errorCode !== 0) {
        console.log("Connection lost: " + res.errorMessage);
        Connect();
    }
}

// Mensagem recebida
function MessageArrived(message) {
    console.log("Mensagem recebida no tópico " + message.destinationName + " : " + message.payloadString);
}

// Enviar mensagem JSON
function sendMessage() {
    var jsonMessage = {
        "username": "User123",
        "message": "f",
        "userId": "1"
    };

    // Serializa o JSON em string
    var payload = JSON.stringify(jsonMessage);
    
    // Cria a mensagem MQTT
    var message = new Paho.MQTT.Message(payload);
    message.destinationName = subscription;

    // Publica a mensagem no broker MQTT
    mqttClient.send(message);

    console.log("Mensagem enviada: " + payload);
}
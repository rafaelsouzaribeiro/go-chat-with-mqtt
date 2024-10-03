var hostname = "127.0.0.1";  
var port = 9090; 
var clientId = "webio4mqttexample" + new Date().getUTCMilliseconds();
var username = "root";
var password = "123mudar";
var userId = '1';

var mqttClient = new Paho.MQTT.Client(hostname, port, clientId);
mqttClient.onMessageArrived = MessageArrived;
mqttClient.onConnectionLost = ConnectionLost;


Connect();

function Onclick() {
    document.querySelectorAll(".user-id").forEach(function(element) {
        element.addEventListener("click", function() {
            var id = this.getAttribute("id");
            FetchMessage(id);
        });
    });
}

function FetchMessage(id){
   
    fetch(`/list-message/${id}`) 
        .then(response => {
            if (!response.ok) {
                throw new Error('Erro ao carregar a página');
            }
            return response.text(); 
        })
        .then(json => {
            console.log(json);
        })
        .catch(error => {
            console.error('Erro:', error);
        }); 
}

function SelectUsers(){

    fetch(`/list-users/`) 
        .then(response => {
            if (!response.ok) {
                throw new Error('Erro ao carregar a página');
            }
            return response.text(); 
        })
        .then(json => {
            console.log(json);
            const obj = JSON.parse(json);
            var elements="";
            
            obj.forEach(element => {
                elements+=`<li id='${element.id}' class='user-id'>
                    <img src='${element.photo}' alt='${element.username}' />
                    <span>${element.username}</span>
                    <div class='clear'></div>
                </li>`;
            });

           
             document.getElementById('users').innerHTML = elements;
             Onclick();
        })
        .catch(error => {
            console.error('Erro:', error);
        });
}

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
    SelectUsers();
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
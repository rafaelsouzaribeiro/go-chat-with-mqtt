var hostname = "127.0.0.1";  
var port = 9090; 
var clientId = "webio4mqttexample" + new Date().getUTCMilliseconds();
var usernameCon = "root";
var password = "123mudar";
var userId = "";
var userName = "";

var mqttClient = new Paho.MQTT.Client(hostname, port, clientId);
mqttClient.onMessageArrived = MessageArrived;
mqttClient.onConnectionLost = ConnectionLost;


Connect();

function Onclick() {
    document.querySelectorAll(".user-id").forEach(function(element) {
        element.addEventListener("click", function() {
            var id = this.getAttribute("id");
            userId=id;
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
            if(json!=null){
                var json = JSON.parse(json);

                document.getElementById("chat-body").innerHTML="";
                json.forEach(element => {
                    document.getElementById("chat-body").innerHTML+=`<div class="message sent">
                        <p>${element.message}
                        </p>
                        <span class="time">${formatTimestamp(element.times)}</span>
                    </div>`;
                });

            }
            
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
            if (json!=null){
                var obj = JSON.parse(json);
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
            }
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
        userName: usernameCon,
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
    var json = JSON.parse(message.payloadString)

    if (json!=null && userId){
        document.getElementById("chat-body").innerHTML+=`<div class="message sent">
            <p>${json.message}
            </p>
            <span class="time">${formatTimestamp(json.times)}</span>
        </div>`;
        console.log("Mensagem recebida no tópico " + message.destinationName + " : " + message.payloadString);
    }

   
}


function sendMessage() {
    message=document.getElementById("message-input").value.trim();

    if (message!=""){
        var jsonMessage = {
            "username": loggeduser,
            "message": message,
            "userId": userId,
            "loggedId":loggedId,
            "times" :new Date().toISOString(),
        };
    
        var payload = JSON.stringify(jsonMessage);
        
        
        var message = new Paho.MQTT.Message(payload);
        message.destinationName = subscription;
    
        mqttClient.send(message);
    
        console.log("Mensagem enviada: " + payload);        
    }
   
}

function formatTimestamp(timestamp) {
    var date = new Date(timestamp);

    var day = date.getDate();
    var month = date.getMonth() + 1; 
    var year = date.getFullYear();

    var hours = date.getHours();
    var minutes = date.getMinutes();
    var seconds = date.getSeconds();

    day = day < 10 ? '0' + day : day;
    month = month < 10 ? '0' + month : month;
    hours = hours < 10 ? '0' + hours : hours;
    minutes = minutes < 10 ? '0' + minutes : minutes;
    seconds = seconds < 10 ? '0' + seconds : seconds;

    return `${day}/${month}/${year} ${hours}:${minutes}:${seconds}`;
}


window.addEventListener("load", function() {
    SelectUsers();
});

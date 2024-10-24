var clientId = "webio4mqttexample" + new Date().getUTCMilliseconds();
var userId = "";
var userName = "";
var pageTotalU = 0;
var pageTotalM = 0;
var messageObject = {};
var hasmoreusers=true;
var hasmoremessages=true;

var mqttClient = new Paho.MQTT.Client(hostname, parseInt(port), clientId);
mqttClient.onMessageArrived = MessageArrived;
mqttClient.onConnectionLost = ConnectionLost;
const emojis = ["😀", "😃", "😄", "😁", "😆", "😅", "😂", "🤣", "😊", 
    "😇","💗","💔","❤️‍🔥","❤","😍","😴","😌","😌","🤤","😱","😭","😩","🤬","🤡","👹","👺","👻","👽"
    ,"👾","🙌","🤝","🙏","👍","👎"];
const emojiContainer = document.querySelector(".icones");


emojis.forEach((emoji) => {
    const emojiDiv = document.createElement("div");
    emojiDiv.classList.add("emoji");
    emojiDiv.innerText = emoji;
    emojiContainer.appendChild(emojiDiv);
});


document.getElementById("icon").addEventListener("click",function(){
    var icones = document.getElementById("icones");
    if (icones.style.display === "none") {
        icones.style.display = "block"; 
    } else {
        icones.style.display = "none";  
    }
});

function OnclickImo(){
    emojiContainer.addEventListener("click",(e)=>{
        document.getElementById("message-input").value+=e.target.innerText;
    })
}

Connect();
OnclickImo();

function Onclick() {
    document.querySelectorAll(".user-id").forEach(function(element) {
        element.addEventListener("click", function() {
            var id = this.getAttribute("id");
            userId=id;
            document.querySelectorAll(".user-id").forEach(function(el) {
                el.style.backgroundColor = "#57606f";
            });
            this.style.backgroundColor = "#1e272e";
            FetchMessage(id);
        });
    });
}


function FetchMessage(id){
    
    fetch(`/list-message/${id}/${loggedId}/`) 
        .then(response => {
            if (!response.ok) {
                throw new Error('Error loading page');
            }
            return response.text(); 
        })
        .then(json => {
            if(json!=null){
                var json = JSON.parse(json);

                document.getElementById("chat-body").innerHTML="";
                messageObject={};

                if (json!=null){
                    json.forEach(element => {
                        pageTotalM=element.page_total;
                        console.log("total>>",pageTotalM)
                        messageObject[element.times] = element;
                      
                    });


                    var sortedTimes = Object.keys(messageObject).sort();
                    sortedTimes.forEach(time=>{
                        var element = messageObject[time];
                        document.getElementById("chat-body").innerHTML+=`<div class="message ${element.types}">
                            <p>${element.message}
                            </p>
                            <span class="time">${formatTimestamp(element.times)}</span>
                        </div>`;
                    });
                    document.getElementById("loader").click();
                }
         

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
                    pageTotalU = element.page_total;
                    if (loggedId==element.id){return;} 
                    
                    elements+=`<li id='${element.id}' class='user-id'>
                        <img src='${element.photo}' alt='${element.username}' />
                        <span>${element.username}</span>
                        <div class='clear'></div>
                    </li>`;
                });

            
                document.getElementById('users').innerHTML += elements;
                Onclick();
            }
        })
        .catch(error => {
            console.error('Erro:', error);
        });
}

function SelectUsersindex() {
    if (hasmoreusers) {
        pageTotalU--;

        fetch(`/list-users-index/${pageTotalU}`)
            .then(response => {
                if (!response.ok) {
                    throw new Error('Erro ao carregar a página');
                }
                return response.json();
            })
            .then(json => {
            
                if (!json || json.length === 0) {
                    hasmoreusers = false;
                    return; 
                }

                try {
                    var elements = "";
                    json.forEach(element => {
                    if (loggedId == element.id) { return; }

                    elements += `<li id='${element.id}' class='user-id'>
                        <img src='${element.photo}' alt='${element.username}' />
                        <span>${element.username}</span>
                        <div class='clear'></div>
                    </li>`;
                });

                document.getElementById('users').innerHTML += elements;
                Onclick();
            } catch (e) {
                hasmoreusers = false;
            }
        
            })
            .catch(error => {
                hasmoreusers = false;
                console.error('Erro:', error);
            });
    }
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
    console.log(json);

    if (json!=null){
        if (json.receive == loggedId && json.userId == userId) { 
            document.getElementById("chat-body").innerHTML+=`<div class="message received">
                <p>${json.message}
                </p>
                <span class="time">${formatTimestamp(json.times)}</span>
            </div>`;
            
        }

        if (json.receive == userId && json.userId == loggedId){
                document.getElementById("chat-body").innerHTML+=`<div class="message sent">
                <p>${json.message}
                </p>
                <span class="time">${formatTimestamp(json.times)}</span>
            </div>`;
        }

    }

   
}


function sendMessage() {
    message=document.getElementById("message-input").value.trim();

    if (message!="" && userId){ 
        var jsonMessage = {
            "username": loggeduser,
            "message": message,
            "userId": userId,
            "receive":loggedId,
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


function loadPreviousMessages() {
    if (hasmoremessages){
        pageTotalM--;
        fetch(`/list-message-index/${userId}/${loggedId}/${pageTotalM}`)
        .then(response => {
            if (!response.ok) {
                throw new Error('Error loading page');
            }
            return response.json(); 
        })
        .then(json => {
            if (!json || json.length === 0) {
                hasmoremessages = false;
                return; 
            }
                
            json.forEach(element => {
                messageObject[element.times] = element;
            });

            document.getElementById("chat-body").innerHTML="";

            var sortedTimes = Object.keys(messageObject).sort();
            sortedTimes.forEach(time=>{
                var element = messageObject[time];
                document.getElementById("chat-body").innerHTML+=`<div class="message ${element.types}">
                    <p>${element.message}
                    </p>
                    <span class="time">${formatTimestamp(element.times)}</span>
                </div>`;
            });
            
        })
        .catch(error => {
            hasmoremessages=false;
            console.error('Erro:', error);
        }); 
    }
}

document.getElementById("loader").addEventListener('click',function(){
    loadPreviousMessages();
;});

document.getElementById("btnusers").addEventListener('click',function(){
    SelectUsersindex();
;});



document.getElementById("chat-body").addEventListener('scroll', function(event) {
    if (this.scrollTop === 0) {
        loadPreviousMessages();
    }
});

window.addEventListener("load", function() {
    SelectUsers();
});


function logout() {
    fetch('/logout', {
        method: 'POST',
        credentials: 'same-origin'
    }).then(response => {
        if (response.ok) {
            window.location.href = '/'; 
        } else {
            alert('Logout failed!');
        }
    }).catch(error => {
        console.error('Error during logout:', error);
    });
}
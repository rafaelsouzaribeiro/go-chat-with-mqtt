<p>Web chat with MQTT and Cassandra, including user registration, logged-in and logged-out users, and emoji support, implemented in Go, HTML, CSS, and JavaScript </p>
<br/>
<p>To create the keyspace and tables in Cassandra, just run the following Makefile commands in the project root:</p>

 ```
Create KeysPace: make create-keyspace
Create tables: make migrateup
Drop tables and keyspace: make migratedown
 ```

<p>If you don't have make installed, install it with the following command:</p>

 ```
sudo apt install make
 ```

<p>To run Cassandra in Docker, navigate to the internal/infra/database/cassandra directory and run:</p>

 ```
sudo docker compose up
 ```

<p>Also, access the website <a href="http://rafael-developer.net/2024/10/11/go-chat-com-mqtt-implementando-cadastro-de-usuario-login-e-logout-lista-de-usuario-e-envio-de-mensagem/" title="Web chat with MQTT">here</a>.</p>
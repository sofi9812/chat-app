CHAT ROOM APPLICATION USING WEBSOCKETS IN CLI

this is a chat app, where user can enter any room and communicate with anyone
for now as this is only a simple project without a proper database; i hardcode the user to only  have 2 users (user1, user2)

people can communicate if they are in the same room, but not if they are in different room
messages will be saved on the particular room name file '.txt'

## Usage
there are 2 ways to run this program

 - Local Run without docker
    1. in the root folder, execute main.go with command "go run server.go" alternatively you can run the binary "./server.exe" if you are using windows pc and does not have go installed
    2. open any other terminal/powershell go to directory "client" run command "go run client.go <username> <password> <room>" or alternatively using binary "./client.exe <username> <password> <room>"
        - e.g: go run client.go user1 password1 general
    3. as there are only 2 users setup for now here are the username and password
        - username: user1; password: password1
        - username: user2; password: password2
    4. once you have seen the prompt "U can start chatting now!" the user between different terminal can check with each other
    5. the messages will be saves on folder "messages" in the root folder; where the txt files will be the same as what the room name is

 - Local Run with docker
    1. in the root folder run "docker build -t <any-image-name> ."
    2. then run "docker run -p 8080:8080 <image-name-you-put-when-build>
    3. open any terminal outside the container and go to the client directory of the code and run "go run client.go <username> <password> <room>"
        - e.g: go run client.go user1 password1 general
    4. as there are only 2 users setup for now here are the username and password
        - username: user1; password: password1
        - username: user2; password: password2
    5. open another terminal and repeat step 3 to simulate another new user(user2) 
    6. the messages will be saved inside the root folder in the container "/root/messages/<room_name>.txt"

## AUTHOR
- MUHAMMAD SOFIYYULLAH BIN HASHIM# chat-app

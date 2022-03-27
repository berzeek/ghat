# Ghat - Go Chat

## Goal
Build a Go chat server with the following features:
- Multiple clients can connect via telnet and can send messages to the server
- When a message is relayed to the server it should relay the following into to all connected clients:
    - timestamp
    - name of client
- All messages should be logged to a local log file
- Read the following configuration settings from a local config file:
    - listening port
    - IP
    - log file location

### Optional Features
- HTTP rest api to post messages.
- HTTP rest api to query messages.
- Segment rooms / channels.
- Add ignore option.

### Run the app
- `go build .` - build application.
- `./ghat` - start the server.
- `telnet localhost 8888` - connect the client (do the same in another terminal to see other users messages)

### Commands
- `/name <name>` - create a username for the user.
- `/join <name>` - join a room; new room will be created for non-existent room.
- `/rooms`       - list all available rooms.
- `/all <msg>`   - send message to all users in a room.
- `/exit`        - end the connection.

### Limitations
- Validation/Sanitization needs to be added to all user input.
- User can only be in one room at a time.
- Message logging only works if another client connects.
- All messages are broadcast publicly. In a later iteration I would add a direct message feature.

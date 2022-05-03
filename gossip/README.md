# gossip

## Problem Statement
Design a sample chat service and a terminal based client. A user should be
able to register from the terminal and login using their username and password.
The service should handle both chat rooms and one to one chat. Users would be 
able to create a chat room and invite others to join the room. Messages sent to 
a room is visible to everyone in the room. There should be one default chat room,
public to all users in the server. One to one messages are private to the sender
and receiver. The service should be able to handle multiple users at a time.

We care about message delivery and message persistence. We want to make sure that
messages are delivered to the user even if the server crashes or the user is
offline.



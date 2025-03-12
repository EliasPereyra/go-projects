# TCP and UDP Servers

Both TCP and UDP are protocols which allow to establish a communication between two parties: the server and a client.

TCP does a three-handshake procedure in order to establish a communication. The server is open and listening to any incoming messages, and the client proceeds to the send the message.

The message is a stream of bytes, so in order to handle the bytes one must parse it into the structure data is needed, for eg. to a string.

TCP is for a persistent connection.

UDP doesn't need to set up a connection between the two parties, and doesn't keep track of what is sent.It has no handshaking and it's unreliable. It's suitable in cases where error checking are not needed.

UDP avoids the overhead of the process. Time-sensitive apps often use UDP because dropping packets is preferable to waiting for packets delayed due to retransmission.

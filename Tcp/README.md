
a. What are packages in your implementation? What data structure do you use to transmit data and meta-data?
- Packages: In this implementation the only package imported is the fmt package (for output) and time (for delaying execution). 
Data Structure The data structure used to transmit data and meta-data is a custom Packet struct, which contains a single field Flag. This field holds integer values that represent the TCP flags (SYN, ACK, or SYN-ACK). The transmission happens using Go channels (clientToServer and serverToClient) which simulate the network communication between client and server.

b. Does your implementation use threads or processes? Why is it not realistic to use threads?
- The implementation uses goroutines, which are lightweight concurrent functions in Go that run in the background. 
Using threads directly can be problematic because of their heavy resource usage. 

c. In case the network changes the order in which messages are delivered, how would you handle message re-ordering?
- To handle message re-ordering in a realistic scenario:
Sequence Numbers: Every packet would have a sequence number in the Packet struct. This number allows both the client and server to reconstruct the correct order of packets.
Buffering: If packets arrive out of order, the receiver can buffer them temporarily until all preceding packets are received.
Acknowledgments: The receiver sends acknowledgments (ACKs) for the highest-sequenced packet that it has received in order, helping the sender know which packets have arrived.

d. In case messages can be delayed or lost, how does your implementation handle message loss?
- In this basic implementation, there is no mechanism for handling message loss. 
In a real TCP handshake, if the SYN-ACK does not arrive within a timeout period, the client would retransmit the SYN. Similarly, if the server does not receive the ACK after sending SYN-ACK, it may retransmit the SYN-ACK.

e. Why is the 3-way handshake important?
- The 3-way handshake is important forseveral reasons:

Connection Establishment: It ensures that both the client and the server are ready to establish a connection 
Synchronization: It synchronizes the sequence numbers between the client and the server. 
Confirmation of Communication: It confirms that both parties are capable of receiving and transmitting data. T
Without the 3-way handshake, there would be a risk of unreliable connections
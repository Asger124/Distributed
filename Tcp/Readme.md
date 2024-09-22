# TCP PROTOCOL

a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?

    -Packages are implemented as structs. A packet has a header and a message. The message is a string and the header,
    is in itself a struct where meta-data resides. In this implementation the only data a header contains is a SYN number
    and an ACK number. The other data has been asbtracted away for simplicty.

b) Does your implementation use threads or processes? Why is it not realistic to use threads?

    -This implementation uses threads. Threads are not realistic as they are run locally and shares resources stored on a local
    machine, which simplifies communication and data sharing. 
    In a realistic scenario the handshake would be performed on isolated machines over a network using processes.

c) In case the network changes the order in which messages are delivered, how would you handle message re-ordering?

    -In this implementation no message re-ordering is handled. A way you could handle message re-ordering would be to 
    sort sequence numbers on their sizes. After the handshake has been performed when data are received the ACK number sould be set to
    to the sequence number + the size data received. 
    When a host receives this ACK number, it is being 'told' that the next message is expected to be delivered at the ACK number(representing bytes), and a host will set its next sequence number accordingly.This way the sequence number can be used to keep track of which messages arrived in which order.

    Our implementation stops after the handshake has been performed however.

d) In case messages can be delayed or lost, how does your implementation handle message loss?

    - This implementation does not account for message loss, but as explained in the question above a receiving host will send an ACK
    number to the sending host, in a realistic scenario. A host which has sent packets can check if the receiving host has gotten all messages by checking the ACK packets it receives. If the sender does not receive the expected ACK it can retransmit the lost packet.

e) Why is the 3-way handshake important?

    -The 3-way handshake is important as it ensures that both hosts are ready and able to communicate.
    It also ensures that both hosts agrees on a initial sequence numbers which are acnkowledged in the handshake, 
    which helps to providea reliable data transmission.


    //HOW TO RUN:

    From the terminal navigate to the folder \Distributed\Tcp.   
    When in the folder Tcp, write: 'go run .'

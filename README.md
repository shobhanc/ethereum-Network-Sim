# ethereum-Network-Sim
Protocol interactions observed in my testing:

Bootstrap node is initialized.
Regular Nodes are booted with the bootstrap nodes ip and port number.
Regular nodes ping the bootstrap node and update their database with the bootstrap node as a neighbor after successful pongs and exchange of node records.
Same is reciprocated by the bootstrap node, which pings the regular nodes and updates its neighbors db with the bootstrap node after exchange of node records.
 Outside of this the nodes ping each other as keep alive pings to eliminate dead nodes from their database.
The bootstrap node and the regular nodes also exchange the node records as mentioned above  like  name of identity scheme, public key, ip, tcp/udp port numbers.
Subsequently bootstrap node and the regular nodes issue find node to each other and gets responded to by a list of the responders neighbors.
Using the recursive lookup process the bootstrap and regular nodes start pinging the neighbors received above.
 After verifying and exchanging the node records with the newly pinged nodes reported by the bootstrap node, the newly responded nodes are added to the neighbors list. Since in this experiment there is no scope for recursive lookup for XOR distances as the number of nodes are small.
Subsequently the nodes repeat the keep alives with the ping/pong and check for changes in their neighbors node records using the node records exchanged and also continue to look for new neighbors of known nodes. 

Things that I could not finish

Instead of sleeping for a definite amount of time for bootstrap node to init. I need to wait on the channel to listen to the initDone event.
Register nodeAdded callback in the table of the node callbacks, so that the callback can be invoked when the node is added.


Run Test:
go test -v -timeout 0 -run TestRun

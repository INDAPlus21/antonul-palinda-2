Problem: The code gets stuck waiting for someone to read the channel but since it is waiting it can't read the channel.
###
Solution: Put the "inputing" to the channel in a goroutine so that the program can reach the reading of the channel
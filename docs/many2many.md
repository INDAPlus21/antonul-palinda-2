* What happens if you switch the order of the statements
  `wgp.Wait()` and `close(ch)` in the end of the `main` function?

    **The channel closes before the producers have sent their strings so they try to send on a closed channel**

* What happens if you move the `close(ch)` from the `main` function
  and instead close the channel in the end of the function
  `Produce`?

    **The channel closes when the producers that finishes first finishes and then the other producers tries to send on a closed channel**
* What happens if you remove the statement `close(ch)` completely?

    **Basicaly nothing the main routine will still end and when the main routine ends the channel is "closed" (removed from existance and never thought about again R.I.P)**
* What happens if you increase the number of consumers from 2 to 4?

    **The program finishes faster: More consumers -> Less downtime for producers**
* Can you be sure that all strings are printed before the program
  stops?
  
  **No, sometimes the main routine finishes faster than the last consumer can "process" the data since it only waits for the "producers" to finish and not the "consumers"**
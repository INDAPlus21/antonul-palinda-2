Problem: The main function ends before the print routine has time to finish
###
Solution: Add a waitgroup and force the program to wait for the print routine to finish 
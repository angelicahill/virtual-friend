# Virtual Friend 

**What is this program and what is its purpose?**

This was the first program I wrote and the purpose was to practice the basic concepts I had learnt thus far: functions, loops, if/else statements, taking user input etc. The purpose of this program is for it to be a fun little personal friend that asked and answered questions for you. 
This is an MVP that could eventually turn into a more substantial personalized friend experience in future iterations. 

**How would you use this?**

This is not a tool, this is more of a mini-game. You use this when you're bored as a fun little interactive experience. 

**Why did I personally write this program?**

To practice the basic concepts I had learnt thus far: functions, loops, if/else statements, taking user input etc. in an actual application as opposed to via mini online excercises.

**How does it work?** 

The user interacts with the "virtual friend" via the terminal, answer questions and getting responses back. 

**How to run program?**

If you would like to run this program yourself in the terminal all you have to do is: 
- Clone this repository.  
- Open your terminal and navigate to this directory before running “go run virtual-friend.go” 
- The terminal should then display the following:

```
Are you in the mood to make a virtual friend today?

```
From then onwards you follow the instructions given by the "virtual friend".

**Known bugs I am un the process of debugging:**
- If the virtual friend gets your name wrong and you say “NO” to “Is that right? Your name is //USER INPUT//?” the program stalls and does not continue.
- The user has to enter the answer in one work as I have not yet added the ability for the answer to have spaces in it. 
- After "What date did Shakespeare die?" the program reverts back to the following: 
 ```
What date did Shakespeare die?
1616
Correct. A sad year!
Please answer Shakespeare or Maths. Thank you :)
Could you please choose from the two options I gave you: Maths or Shakespeare. Thank you! :)

```
Following this if you select maths it will run as desired and finish with: 

```
Well done and Thanks :) On to the next chapter of our growing friendship...see you next time!

```

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Are you in the mood to make a virtual friend today?")
	for {
		if scanner.Scan() {
			virtualFriend := scanner.Text()
			if virtualFriend == "yes" || virtualFriend == "Yes" {
				fmt.Println("Great, let's do this!")
				break
			}
			if virtualFriend == "no" || virtualFriend == "No" {
				fmt.Println("That's ok we can get to know each other another time")
				fmt.Println("See you soon!")
				return
			}
			fmt.Println("Please answer yes or no so I know if you want to be virtual buddies. Please do not add spaces when inputting your answer.")
		}
	}s

	fmt.Println("Let's start with the basics. What's your name?")
	FirstNameQu := ""
	fmt.Scanln(&FirstNameQu)
	fmt.Println("Hello " + FirstNameQu + " Great to meet you!")

	for {
		fmt.Println("Do you want to know my name?")
		knowMyNameAnswer := ""
		fmt.Scanln(&knowMyNameAnswer)
		if knowMyNameAnswer == "yes" || knowMyNameAnswer == "Yes" {
			fmt.Println("Ok. My name is Angelica. Great to meet you " + FirstNameQu)
			break
		}
		if knowMyNameAnswer == "no" || knowMyNameAnswer == "No" {
			fmt.Println("Ok then. I can be a mystery. Let's move on to the next question...")
			break
		}
		fmt.Println("Please answer yes or no so I know if you want to be know my name.")
	}
	fmt.Println("Next Question. Where are you from?")
	whereFrom := ""
	fmt.Scanln(&whereFrom)
	fmt.Println("Nice. So just to check I have this right you're name is " + FirstNameQu + " and you're from " + whereFrom)
	fmt.Println("Is that right?" + " Your name is " + FirstNameQu + "?")
	doubleCheckName := ""
	fmt.Scanln(&doubleCheckName)
	for {
		if doubleCheckName == "yes" || doubleCheckName == "Yes" {
			fmt.Println("Glad I got that right.")
			break
		}
		if doubleCheckName == "no" || doubleCheckName == "No" {
			fmt.Println("So sorry about that. I'll try again. What's your name?")
			reenterName := ""
			fmt.Scanln(&reenterName)
			fmt.Println("So you're name is not " + FirstNameQu + " it is " + reenterName + " right?")
			confirmReenterName := ""
			fmt.Scanln(&confirmReenterName)
			if confirmReenterName == "yes" || confirmReenterName == "Yes" {
				fmt.Println("Yay. Got there eventually.")
				break
			}
			if confirmReenterName == "no" || confirmReenterName == "No" {
				fmt.Println("I am so so sorry. I give up. You don't want to be friends with such a forgetful goldfish anyway. Goodbye for now. I'll work on my memory and maybe we can be friends in the future.")
				return
			}

			fmt.Println("Ok so I think I got this now. Your name is " + confirmReenterName + " and you are from " + whereFrom + ". I hope that's right or I give up. Is it?")
			finalNameCheck := ""
			fmt.Scanln(&finalNameCheck)
			if finalNameCheck == "yes" || finalNameCheck == "Yes" {
				fmt.Println("Phew. I was worried I was going to offend you there with my terrible memory. Next Question...")
				break
			}
			if finalNameCheck == "no" || finalNameCheck == "No" {
				fmt.Println("I am so so sorry. I give up. You don't want to be friends with such a forgetful goldfish anyway. Goodbye for now. I'll work on my memory and maybe we can be friends in the future.")
				return
			}

			fmt.Println("Please answer yes or no. Thank you :)")

		}
	}

	fmt.Println("Next Question...pick a topic: Maths or Shakespeare...")
	mathsOrShakespeare := ""
	fmt.Scanln(&mathsOrShakespeare)
	if mathsOrShakespeare == "Maths" || mathsOrShakespeare == "maths" {
		fmt.Println("You chose Maths, so here is little quiz just for you...")
		fmt.Println("Welcome to my Math's Game. Hope you enjoy!")

		fmt.Println("What is 2 + 2?")
		twoPlusTwoAnswer := ""
		fmt.Scanln(&twoPlusTwoAnswer)
		checkTwoPlusTwoAnswer := TwoPlusTwo(twoPlusTwoAnswer)
		fmt.Println(checkTwoPlusTwoAnswer)

		fmt.Println("What is 5 + 5?")
		fivePlusFiveAnswer := ""
		fmt.Scanln(&fivePlusFiveAnswer)
		checkFivePlusFiveAnswer := FivePlusFive(fivePlusFiveAnswer)
		fmt.Println(checkFivePlusFiveAnswer)

		fmt.Println("What is 10 x 2?")
		tenTimesTwoAnswer := ""
		fmt.Scanln(&tenTimesTwoAnswer)
		checkTenTimesTwoAnswer := TenTimesTwo(tenTimesTwoAnswer)
		fmt.Println(checkTenTimesTwoAnswer)

		fmt.Println("What is 6 + 6?")
		sixPlusSixAnswer := ""
		fmt.Scanln(&sixPlusSixAnswer)
		checkSixPlusSixAnswer := SixPlusSix(sixPlusSixAnswer)
		fmt.Println(checkSixPlusSixAnswer)

		fmt.Println("What is 4 + 4?")
		fourPlusFourAnswer := ""
		fmt.Scanln(&fourPlusFourAnswer)
		checkFourPlusFourAnswer := FourPlusFour(fourPlusFourAnswer)
		fmt.Println(checkFourPlusFourAnswer)

		fmt.Println("Well done and Thank you :) Do you want to try the other quiz? It's on Shakespeare")

		playAnotherQuizShakespeare := ""
		fmt.Scanln(&playAnotherQuizShakespeare)
		if playAnotherQuizShakespeare == "yes" || playAnotherQuizShakespeare == "Yes" {
			fmt.Println("Great here we go...")
			fmt.Println("I am going to ask you 5 questions about Shakespeare. Side note: my major in college was Shakespeare's applicability to today's political climate so they will get progressively harder...")

			fmt.Println("What is Shakespeare's most famous and performed play about love? Which also happens to be my favourite Baz Luhrmann movie, makes me cry every time! You only need to give the first word")
			lovePlayAnswer := ""
			fmt.Scanln(&lovePlayAnswer)
			checkLovePlayAnswer := ShakeQuOne(lovePlayAnswer)
			fmt.Println(checkLovePlayAnswer)

			fmt.Println("What play are the characters Ariel and Calaban in, hint: The...?? You only have to get the second word...")
			tempestAnswer := ""
			fmt.Scanln(&tempestAnswer)
			checkTempestAnswer := ShakeQuTwo(tempestAnswer)
			fmt.Println(checkTempestAnswer)

			fmt.Println("Who is the best Shakespeare actor currently alive? This is my opinion but this will establish if you are friend or best friend material, I only need his first name.")
			actorAnswer := ""
			fmt.Scanln(&actorAnswer)
			checkActorAnswer := ShakeQuThree(actorAnswer)
			fmt.Println(checkActorAnswer)

			fmt.Println("Finish the line from Hamlet: To be, or not to be, that is the...")
			toBeAnswer := ""
			fmt.Scanln(&toBeAnswer)
			checkToBeAnswer := ShakeQuFour(toBeAnswer)
			fmt.Println(checkToBeAnswer)

			fmt.Println("What date did Shakespeare die?")
			deathAnswer := ""
			fmt.Scanln(&deathAnswer)
			checkDeathAnswer := ShakeQuFive(deathAnswer)
			fmt.Println(checkDeathAnswer)

		}
		if playAnotherQuizShakespeare == "no" || playAnotherQuizShakespeare == "No" {
			fmt.Println("Ok. On to the next chapter of our friendship...")
		}

		fmt.Println("Please answer Shakespeare or Maths. Thank you :)")

	}

	if mathsOrShakespeare == "Shakespeare" || mathsOrShakespeare == "shakespeare" {
		fmt.Println("Great here we go...")
		fmt.Println("I am going to ask you 5 questions about Shakespeare. Side note: my major in college was Shakespeare's applicability to today's political climate so they will get progressively harder...")

		fmt.Println("What is Shakespeare's most famous and performed play about love? Which also happens to be my favourite Baz Luhrmann movie, makes me cry every time! You only need to give the first word")
		lovePlayAnswer := ""
		fmt.Scanln(&lovePlayAnswer)
		checkLovePlayAnswer := ShakeQuOne(lovePlayAnswer)
		fmt.Println(checkLovePlayAnswer)

		fmt.Println("What play are the characters Ariel and Calaban in, hint: The...?? You only have to get the second word...")
		tempestAnswer := ""
		fmt.Scanln(&tempestAnswer)
		checkTempestAnswer := ShakeQuTwo(tempestAnswer)
		fmt.Println(checkTempestAnswer)

		fmt.Println("Who is the best Shakespeare actor currently alive? This is my opinion but this will establish if you are friend or best friend material, I only need his first name.")
		actorAnswer := ""
		fmt.Scanln(&actorAnswer)
		checkActorAnswer := ShakeQuThree(actorAnswer)
		fmt.Println(checkActorAnswer)

		fmt.Println("Finish the line from Hamlet: To be, or not to be, that is the...")
		toBeAnswer := ""
		fmt.Scanln(&toBeAnswer)
		checkToBeAnswer := ShakeQuFour(toBeAnswer)
		fmt.Println(checkToBeAnswer)

		fmt.Println("What date did Shakespeare die?")
		deathAnswer := ""
		fmt.Scanln(&deathAnswer)
		checkDeathAnswer := ShakeQuFive(deathAnswer)
		fmt.Println(checkDeathAnswer)
	}

	if mathsOrShakespeare != "Maths" || mathsOrShakespeare != "maths" || mathsOrShakespeare != "Shakespeare" || mathsOrShakespeare != "shakespeare" {
		fmt.Println("Could you please choose from the two options I gave you: Maths or Shakespeare. Thank you! :)")
		z := ""
		fmt.Scanln(&z)
		if z == "Shakespeare" || z == "shakespeare" {
			fmt.Println("Great here we go...")
			fmt.Println("I am going to ask you 5 questions about Shakespeare. Side note: my major in college was Shakespeare's applicability to today's political climate so they will get progressively harder...")

			fmt.Println("What is Shakespeare's most famous and performed play about love? Which also happens to be my favourite Baz Luhrmann movie, makes me cry every time! You only need to give the first word")
			lovePlayAnswer := ""
			fmt.Scanln(&lovePlayAnswer)
			checkLovePlayAnswer := ShakeQuOne(lovePlayAnswer)
			fmt.Println(checkLovePlayAnswer)

			fmt.Println("What play are the characters Ariel and Calaban in, hint: The...?? You only have to get the second word...")
			tempestAnswer := ""
			fmt.Scanln(&tempestAnswer)
			checkTempestAnswer := ShakeQuTwo(tempestAnswer)
			fmt.Println(checkTempestAnswer)

			fmt.Println("Who is the best Shakespeare actor currently alive? This is my opinion but this will establish if you are friend or best friend material, I only need his first name.")
			actorAnswer := ""
			fmt.Scanln(&actorAnswer)
			checkActorAnswer := ShakeQuThree(actorAnswer)
			fmt.Println(checkActorAnswer)

			fmt.Println("Finish the line from Hamlet: To be, or not to be, that is the...")
			toBeAnswer := ""
			fmt.Scanln(&toBeAnswer)
			checkToBeAnswer := ShakeQuFour(toBeAnswer)
			fmt.Println(checkToBeAnswer)

			fmt.Println("What date did Shakespeare die?")
			deathAnswer := ""
			fmt.Scanln(&deathAnswer)
			checkDeathAnswer := ShakeQuFive(deathAnswer)
			fmt.Println(checkDeathAnswer)

			fmt.Println("Well done and Thank you :) On to the next chapter of our growing friendship...")

		}
		if z == "Maths" || z == "maths" {
			fmt.Println("What is 2 + 2?")
			twoPlusTwoAnswer := ""
			fmt.Scanln(&twoPlusTwoAnswer)
			checkTwoPlusTwoAnswer := TwoPlusTwo(twoPlusTwoAnswer)
			fmt.Println(checkTwoPlusTwoAnswer)

			fmt.Println("What is 5 + 5?")
			fivePlusFiveAnswer := ""
			fmt.Scanln(&fivePlusFiveAnswer)
			checkFivePlusFiveAnswer := FivePlusFive(fivePlusFiveAnswer)
			fmt.Println(checkFivePlusFiveAnswer)

			fmt.Println("What is 10 x 2?")
			tenTimesTwoAnswer := ""
			fmt.Scanln(&tenTimesTwoAnswer)
			checkTenTimesTwoAnswer := TenTimesTwo(tenTimesTwoAnswer)
			fmt.Println(checkTenTimesTwoAnswer)

			fmt.Println("What is 6 + 6?")
			sixPlusSixAnswer := ""
			fmt.Scanln(&sixPlusSixAnswer)
			checkSixPlusSixAnswer := SixPlusSix(sixPlusSixAnswer)
			fmt.Println(checkSixPlusSixAnswer)

			fmt.Println("What is 4 + 4?")
			fourPlusFourAnswer := ""
			fmt.Scanln(&fourPlusFourAnswer)
			checkFourPlusFourAnswer := FourPlusFour(fourPlusFourAnswer)
			fmt.Println(checkFourPlusFourAnswer)

			fmt.Println("Well done and Thanks :) On to the next chapter of our growing friendship...see you next time!")

		}
	}
}

func TwoPlusTwo(i string) string {
	if i == "4" {
		return "Correct!"
	} else {
		return "Sorry that's wrong"
	}
}

func FivePlusFive(i string) string {
	if i == "10" {
		return "Correct!"
	} else {
		return "Sorry that's wrong"
	}

}

func TenTimesTwo(i string) string {
	if i == "20" {
		return "Correct!"
	} else {
		return "Sorry that's wrong"
	}
}

func SixPlusSix(i string) string {
	if i == "12" {
		return "Correct!"
	} else {
		return "Sorry that's wrong"
	}
}

func FourPlusFour(i string) string {
	if i == "8" {
		return "Correct!"
	} else {
		return "Sorry that's wrong"
	}
}

func ShakeQuOne(i string) string {
	if i == "Romeo" || i == "romeo" {
		return "correct - well done!"
	} else {
		return "Sorry that's not right. Let's try another question."
	}
}

func ShakeQuTwo(i string) string {
	if i == "Tempest" || i == "tempest" {
		return "correct - well done!"
	} else {
		return "Sorry that's not right. Let's try another question."
	}
}

func ShakeQuThree(i string) string {
	if i == "Mark" || i == "mark" {
		return "AMAZING! We need to be best friends."
	} else {
		return "Sorry, wrong. Ok. That was a little bit too hard. I'll make the nexrt one easier. Well easier-ish..."
	}
}

func ShakeQuFour(i string) string {
	if i == "question" || i == "Question" {
		return "Perfect."
	} else {
		return "Sorry wrong. There really is no excuse for not knowing this one, unless you made a spelling mistake."
	}
}

func ShakeQuFive(i string) string {
	if i == "1616" {
		return "Correct. A sad year!"
	} else {
		return "Not right. But that was a hard one so we can still be friends"
	}
}

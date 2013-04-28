/*
* GoNemesis is a wonderful text-based RPG
* Developed by Aerilys, MSc Software Engineering at Oxford Brookes University
*/

package main

import (
"fmt" 
"time"
"math/rand"
)

/*
Race structure, to represent a race, like a human or an elf
We just define a name and a specific welcome message 
*/
type Race struct{
	Name string
	WelcomeMessage string
}

/* 
Character class struct, to represent a class (like a role). If you're not familiar with Role Playing Games, take a look at http://en.wikipedia.org/wiki/Role-playing_game
We define a name and a specific welcome message. We also define caracteristics and health points
It's used for both player's character and monsters
*/
type Class struct{
	Name string
	WelcomeMessage string
	Strength int
	Stamina int
	Intelligence int
	HealthPoints int 
	CurrentHP int
}

/*
Generic Character struct. Just a combination of the character's name, a race and a class
*/
type Character struct{
	Name string
	Race Race
	Class Class
}

//Just a boolean to know if we're in game or not
var hasWonTheGame bool = false

//Represents the current Character
var currentCharacter Character

/* *** We init all races, classes and monsters of the game *** */

var races []Race = []Race {
Race{"Human", "Just a classic human ? You are disappointing me! But let's go ahead!"},
Race{"Elf", "An elf ? Great. I love elves!"},
Race{"Dwarf", "A dwarf ? Really ? Go back to your mine!"}}

var classes []Class = []Class {
Class{"Warrior", "Another guy with a big sword. Right?", 15, 12, 10, 30, 30},
Class{"Wizard", "You look like Gandalf! You don't know who is Gandalf?", 10, 10, 18, 25, 25}}

var monsters []Class = []Class{
Class{"Goblin", "A wild goblin appears!", 5, 5, 5, 12, 12},
Class{"Blink monk", "WOLOLOOOOOOOOOOO", 5, 10, 2, 15, 15},
Class{"Orc", "The legend saids orcs come from elves. Are you kidding me?", 12, 1, 10, 16, 16},
Class{"Giant Coconut", "Sounds like there's no coconut in Mercia!", 12, 12, 12, 12, 12},
Class{"Red Dalek", "Go go Power rangers!", 16, 5, 16, 7, 7},
Class{"Draconis", "The most terrible dragon on the universe appears. Here is Draconis!", 18,18, 18, 32, 32}}


//Entry point. Basically we init random and launch the game 
func main() {
	rand.Seed(time.Now().UnixNano())
	LaunchGame()
}

//Display a welcome message and init Character creation
func LaunchGame(){
	fmt.Println("Welcome in the terrible world of Nemesis !")
	fmt.Println("I am Galaad the Archwizard, and I want you to accomplish a terrible quest: kill the terrible Dragon called Draconis !")
	fmt.Println("Are you ready ? Right ! Let's know more about you!")
	
	CreateCharacter()
}

// Launch the character creation. We ask for Name, Race and class
func CreateCharacter(){
	fmt.Println("Let's begin with the simplest. What's your name?")
	var characterName string
	fmt.Scan(&characterName)
	
	if(len(characterName)<1){
		characterName = "Bob"
		fmt.Println("You don't have a name? So okay, I'll call you... Bob!")
	}
	
	currentCharacter.Name = characterName
	
	fmt.Println("Salutations "+currentCharacter.Name+"!")
	fmt.Println("What is your race?")
	for i, race := range races{
		fmt.Println(fmt.Sprintf("%d) %s", i+1, race.Name))
	}
	
	var raceChoice int
	fmt.Scan(&raceChoice)
	
	if raceChoice >= 1 && raceChoice <= 3{
		currentCharacter.Race = races[raceChoice-1]
		fmt.Println(currentCharacter.Race.WelcomeMessage)
	} else {
		currentCharacter.Race = races[0]
		fmt.Println("If you want to be smart, that's a failure. So you will be a pity human!")
	}
	
	fmt.Println("So and what's your class?")
	for i, class := range classes{
		fmt.Println(fmt.Sprintf("%d) %s", i+1, class.Name))
	}
	
	var classChoice int
	fmt.Scan(&classChoice)
	
	if classChoice >= 1 && classChoice <= 3{
		currentCharacter.Class = classes[classChoice-1]
		fmt.Println(currentCharacter.Class.WelcomeMessage)
	} else {
		currentCharacter.Class = classes[0]
		fmt.Println("If you want to be smart, that's a failure. So you will be a dumb warrior!")
	}
	
	fmt.Println("Right. Now you are ready to entrance the terrible cavern of Draconis!")
	EnterDungeon()
}

//The character now enters the dungeon. We start the adventure after displaying a nice message
func EnterDungeon (){
	fmt.Println("Welcome in the terrible cavern of Draconis! In this dungeon, you will encounter monsters and traps. And if you are really brave, maybe you will fight and kill the terrible dragon Draconis!")
	fmt.Println("Here start the adventure...")
	PromptToContinue()
	StartAdventure()
}

//While the current character is alive, and Draconis not defeated, we generate events
func StartAdventure(){
	for currentCharacter.Class.CurrentHP > 0 && !hasWonTheGame{
		GenerateEvent()
	}
}

//We generate a random event. It can be nothing, or a monster
func GenerateEvent(){
	var random int = rand.Intn(100)
	fmt.Println("")
	if random < 90{
		//Damn, it's a monster! So we pick a random one from our list, and launch the fight!
		var monsterIndex int = rand.Intn(len(monsters))
		var monster Class = monsters[monsterIndex]
		Fight(monster)
	} else{
		NothingHappened();
	}
}

//Nothing happened, so we just display a funny message, and heal the hero!
//(If these random sentences remember you a movie, nevermind. It's probably nothing)
func NothingHappened(){
	var index int = rand.Intn(2)
	if index == 0{
		fmt.Println("Brave Sir Robin ran away...  Bravely ran away away!");
	} else if index == 1{
		fmt.Println("What's that ? Coconuts? Found them? In Mercia? The coconut's tropical! ");		
	} else{
		fmt.Println("What's your favourite color ?");		
	}
	currentCharacter.Class.CurrentHP = currentCharacter.Class.HealthPoints
	fmt.Println("Since nothing happened except silly dialogs, you heal all your wounds!")
	fmt.Println(fmt.Sprintf("You have now %d HP!", currentCharacter.Class.CurrentHP))

	PromptToContinue()
}

//The character will fight a terrible monster! Each opponent fights until one is dead. If the character survives, we generate another event. If he dies, we display a message and relaunch the game
//If the monster was Draconis, and the winner is the character, so we display the victory message! 
func Fight(monster Class){
	fmt.Println(monster.WelcomeMessage)
	
	var attackerPotential int = 0
	var defenderPotential int = 0
	
	//If it's a warrior
	if currentCharacter.Race == races[0] {
		attackerPotential = currentCharacter.Class.Strength
		defenderPotential = monster.Strength
	} else{
		attackerPotential = currentCharacter.Class.Intelligence
		defenderPotential = monster.Intelligence
	}	
	
	var attackerHP int = currentCharacter.Class.CurrentHP
	var defenderHP int = monster.HealthPoints
	
	for attackerHP > 0 && defenderHP > 0{
		var randomDmg int = rand.Intn(7)
		var dmg int = attackerPotential + randomDmg - monster.Stamina
		if dmg < 0{
			dmg = 1
		}
		defenderHP -= dmg
		if defenderHP < 0{
			defenderHP = 0
		}

		fmt.Println(fmt.Sprintf("You attack and make %d damages to %s! He still has %d HP!", dmg, monster.Name, defenderHP))

		randomDmg = rand.Intn(7)
		dmg = defenderPotential + randomDmg - currentCharacter.Class.Stamina
		if dmg < 0{
			dmg = 1
		}
		attackerHP -= dmg
		if attackerHP < 0{
			attackerHP = 0
		}

		fmt.Println(fmt.Sprintf("%s attacks and makes you %d damages! You still have %d HP!",monster.Name ,dmg, attackerHP))
	}
	
	if attackerHP > 0 && defenderHP == 0{
		fmt.Println("You won against the terrible " + monster.Name + "!");
		currentCharacter.Class.CurrentHP = attackerHP

		if monster.Name == "Draconis"{
			VictoryOverDraconis()
		}

		PromptToContinue();
	} else if attackerHP == 0 && defenderHP > 0{
		fmt.Println("You lost against the terrible "+monster.Name+"!")
		CharacterDeath()
	} else{
		fmt.Println("In a bloody battle, you kill the terrible monster. But unfortunately, your wounds are too heavy and you are about to die...")
		CharacterDeath()
	}
}

//The player is dead, so we can relaunch the game
func CharacterDeath(){
	currentCharacter.Class.CurrentHP = 0
	fmt.Println("You was brave and powerful, but not enough for the cavern of Draconis! Here dies "+currentCharacter.Name+", an epic "+currentCharacter.Class.Name+" "+currentCharacter.Race.Name+"!")
}

//The player has won. Congratulations!
func VictoryOverDraconis(){
	hasWonTheGame = true
	fmt.Println("Since you have defeated Draconis, you are now the most legendary adventurer in History! Congratulations to you "+currentCharacter.Name+". Excelsior!")
}

//Just a generic method to create pause in prompt so the player can read text before continue
func PromptToContinue(){
	fmt.Println("Enter something to continue...")
	var nothing int
	fmt.Scan(&nothing)
}
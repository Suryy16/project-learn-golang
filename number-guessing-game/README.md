# Number Guessing Game

A simple command-line number guessing game built with Go where players try to guess a randomly generated number between 1 and 100.

## Description

This interactive game challenges players to guess a secret number with a limited number of attempts. Players receive hints after each guess to help them narrow down the correct answer. The game includes three difficulty levels to accommodate different skill levels.

## Features

- **Random Number Generation**: Each game generates a random number between 1 and 100
- **Three Difficulty Levels**:
  - Easy: 10 chances to guess
  - Medium: 5 chances to guess
  - Hard: 3 chances to guess
- **Interactive Feedback**: Receive hints after each guess (higher or lower)
- **Attempt Tracking**: Keeps track of the number of attempts taken

## Requirements

- Go 1.x or higher

## How to Run

1. Make sure you have Go installed on your system
2. Navigate to the project directory:
   ```bash
   cd number-guessing-game
   ```
3. Run the game:
   ```bash
   go run main.go
   ```

## How to Play

1. Start the game and you'll be welcomed with a brief introduction
2. Select your difficulty level (1, 2, or 3)
3. Enter your guess when prompted
4. Based on your guess, you'll receive a hint:
   - If your guess is too high, you'll be told the number is less than your guess
   - If your guess is too low, you'll be told the number is greater than your guess
5. Continue guessing until you find the correct number or run out of chances
6. Win by guessing the correct number before running out of attempts!

## Game Rules

- You must guess a number between 1 and 100
- The number of chances depends on the difficulty level you choose
- Each incorrect guess will consume one chance
- The game ends when you either guess correctly or run out of chances

## Example Gameplay

```
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
You have 5 chances to guess the correct number.

Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)

Enter your choice: 2
Great! You have selected the Medium difficulty level.
Let's start the game!

chances: 5
Enter your guess: 50
Incorrect! The number is greater than 50.

chances: 4
Enter your guess: 75
Incorrect! The number is less than 75.

...
```

## Project Structure

```
number-guessing-game/
├── go.mod          # Go module file
├── main.go         # Main game logic
└── README.md       # Project documentation
```

## License

This is a learning project and is free to use and modify.

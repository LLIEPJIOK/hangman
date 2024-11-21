# Hangman

## Project Description

**Hangman** is a text-based console game where the player tries to guess a hidden word by entering letters one at a time. The game allows the player to choose a category and difficulty level; afterward, the word is randomly chosen from a predefined list. If the player struggles, a hint is available for complex words. Incorrect guesses gradually draw parts of the hangman, and the game ends either when the word is fully guessed or when the hangman drawing is complete. The game is case-insensitive.

---

## Installation and Running

1. Clone the repository:
   ```bash
   git clone git@github.com:LLIEPJIOK/hangman.git
   ```
2. Navigate to the repository folder:
   ```bash
   cd hangman
   ```
3. Run the program:
   ```bash
   go run cmd/hangman/main.go -p resources/words.json -c animals -d medium
   ```

---

## Testing

1. **Word Selection**:  
   Test that words are randomly selected based on the specified category and difficulty.

2. **Game State Updates**:  
   Verify that the game state (word and hangman figure) updates correctly after each input.

3. **Input Validation**:  
   Ensure the program correctly handles inputs of varying cases and invalid entries.

4. **End Conditions**:
   - Victory when the word is fully guessed.
   - Defeat when the hangman is fully drawn after exceeding the allowed attempts.

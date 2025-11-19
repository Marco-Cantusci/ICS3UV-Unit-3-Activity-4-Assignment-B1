/*
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-11-19
 * @fileoverview Check if the user needs a certain colour of sweaters.
 */


// input
const choice: string = prompt("Please choose a sweater colour from the available choices: blue, black, red, white. \n") || ("\n");

// if choice is black or white
if (choice.toLowerCase() == "black" || choice.toLowerCase() == "white") {
  console.log("You have enough sweaters in this colour.");
  // if choice is red
} else if (choice.toLowerCase() == "red") {
  console.log(
    "This colour will look good on you \n",
    "How about a pair of jeans to go with the sweater?",
  )
  // if choice is blue
} else if (choice.toLowerCase() == "blue") {
  console.log("This colour doesn't go well with your eyes.");
} else {
  console.log("Your colour choice is invalid.");
}

console.log("\nDone.");

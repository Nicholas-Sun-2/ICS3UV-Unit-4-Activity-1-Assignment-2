/**
 * @author Nicholas Sun
 * @version 1.0.0
 * @date 2025-12-09
 * @fileoverview Prints a right angle triangle of numbers.
 */

// Input
const rows: number = parseInt(
  prompt("How many rows do you want?") || "1",
);

// Initialize the grid.
let grid: string = "";

for (let row: number = 1; row <= rows; row += 1) {
  for (let column: number = 1; column <= row; column += 1) {
    grid = grid + `${column} `;
  }
  // For every row that is complete, add a newline to create a new row.
  grid = grid + "\n";
}

// Output the completed grid.
console.log(grid);

console.log("\nDone.");

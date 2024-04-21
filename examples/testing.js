function printFibonacci(n) {
  let a = 0,
    b = 1,
    c;
  console.log(a); // Print the first number of the sequence
  if (n >= 1) console.log(b); // Print the second number of the sequence, if needed

  for (let i = 2; i < n; i++) {
    c = a + b; // Next fibonacci number is the sum of the previous two
    console.log(c);
    a = b; // Update a and b for the next iteration
    b = c;
  }
}

// Example usage:
printFibonacci(10); // This will print first 10 Fibonacci numbers

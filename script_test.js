import { Demo } from 'k6/x/demo';

const demo = new Demo();
const greeter = demo.NewGreeter('Hello');
const addResult = demo.XAdd(5, 3);

export default function () {
  console.log(greeter.Greet('k6 User'));
  console.log(`5 + 3 = ${addResult}`);
  
  // Verify the Go function
  if (addResult !== 8) {
    throw new Error('Incorrect sum from Go function');
  }
}
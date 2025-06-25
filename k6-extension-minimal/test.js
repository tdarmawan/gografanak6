import { test } from 'k6/x/minimal';

export default function() {
    const result = test();
    if (result !== "OK") {
        throw new Error(`Expected "OK", got "${result}"`);
    }
    console.log(`Success: ${result}`);
}
// test.js
import { test } from 'k6/x/demo';
export default function() { if (test() !== "OK") throw new Error("Fail"); }
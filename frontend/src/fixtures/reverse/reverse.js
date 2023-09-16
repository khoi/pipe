#!/usr/bin/env node

const args = process.argv.slice(2);
const str = args.join(" ");

/**
 * Reverse a string
 * @param {string} str - The string to be reversed
 * @returns {string} - The reversed string
 */
const reverse = (str) => str.split("").reverse().join("");
console.log(reverse(str));

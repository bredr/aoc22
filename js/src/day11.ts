import { readFileSync } from "fs";
const file = readFileSync("./src/day11input", { encoding: "utf-8" });
const rawMonkeys = file.split("\n\n");

let monkeys = rawMonkeys.map((rawMonkey) => {
  const lines = rawMonkey.split("\n");
  const items = lines[1]
    .replace("Starting items: ", "")
    .trim()
    .split(",")
    .map((x) => parseInt(x.trim()));
  const testDivisibleBy = parseInt(lines[3].trim().match(/\d+/g)![0]);
  const ifTrue = parseInt(lines[4].trim().match(/\d+/g)![0]);
  const ifFalse = parseInt(lines[5].trim().match(/\d+/g)![0]);
  const operation =
    lines[2].trim().replace("Operation: new = ", "(old) => { return ") + ";}";
  return {
    items,
    testDivisibleBy,
    ifTrue,
    ifFalse,
    operation,
    inspected: 0,
  };
});

Array(20)
  .fill(0)
  .forEach(() => {
    monkeys.forEach(
      ({ items, testDivisibleBy, ifTrue, ifFalse, operation }, ix) => {
        while (items.length > 0) {
          const [item] = items.splice(0, 1);
          monkeys[ix].inspected++;
          const updatedItem = Math.floor(eval(operation)(item) / 3);
          if (updatedItem % testDivisibleBy === 0) {
            monkeys[ifTrue].items.push(updatedItem);
          } else {
            monkeys[ifFalse].items.push(updatedItem);
          }
        }
      }
    );
  });

const inspected = monkeys.map(({ inspected }) => inspected);
inspected.sort((a, b) => b - a);
const [top1, top2] = inspected;
console.log("part1 monkey business", top1 * top2);

monkeys = rawMonkeys.map((rawMonkey) => {
  const lines = rawMonkey.split("\n");
  const items = lines[1]
    .replace("Starting items: ", "")
    .trim()
    .split(",")
    .map((x) => parseInt(x.trim()));
  const testDivisibleBy = parseInt(lines[3].trim().match(/\d+/g)![0]);
  const ifTrue = parseInt(lines[4].trim().match(/\d+/g)![0]);
  const ifFalse = parseInt(lines[5].trim().match(/\d+/g)![0]);
  const operation =
    lines[2].trim().replace("Operation: new = ", "(old) => { return ") + ";}";
  return {
    items,
    testDivisibleBy,
    ifTrue,
    ifFalse,
    operation,
    inspected: 0,
  };
});

const gcd = (a: number, b: number): number => (b == 0 ? a : gcd(b, a % b));
const lcm = (a: number, b: number) => (a / gcd(a, b)) * b;
const lcmAll = (ns: number[]) => ns.reduce(lcm, 1);

const superModulo = lcmAll(monkeys.map(({testDivisibleBy})=>testDivisibleBy));
Array(10000)
  .fill(0)
  .forEach(() => {
    monkeys.forEach(
      ({ items, testDivisibleBy, ifTrue, ifFalse, operation }, ix) => {
        while (items.length > 0) {
          const [item] = items.splice(0, 1);
          monkeys[ix].inspected++;
          const updatedItem = eval(operation)(item) % superModulo;
          if (updatedItem % testDivisibleBy === 0) {
            monkeys[ifTrue].items.push(updatedItem);
          } else {
            monkeys[ifFalse].items.push(updatedItem);
          }
        }
      }
    );
  });

const inspected2 = monkeys.map(({ inspected }) => inspected);
inspected2.sort((a, b) => b - a);
const [top21, top22] = inspected2;
console.log("part2 monkey business", top21 * top22);

import { readFileSync } from "fs";

const input = readFileSync("src/day18input", { encoding: "utf-8" }).trim();
const cubes = input
  .split("\n")
  .map((x) => x.split(",").map((y) => parseInt(y, 10))) as Array<
  [number, number, number]
>;

const isAdjacent = (
  a: [number, number, number],
  b: [number, number, number]
) => {
  return (
    Math.abs(a[0] - b[0]) + Math.abs(a[1] - b[1]) + Math.abs(a[2] - b[2]) === 1
  );
};

const part1 = cubes.reduce(
  (yy, a, idx) =>
    yy +
    cubes.reduce(
      (xx, b, idy) => xx - (idx != idy && isAdjacent(a, b) ? 1 : 0),
      6
    ),
  0
);
console.log("part1", part1)
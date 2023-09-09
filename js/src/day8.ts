import { readFileSync } from "fs";
const file = readFileSync("./src/day8input", { encoding: "utf-8" });
const lines = file.split("\n");

const input = lines.map((x) => x.split("").map((y) => parseInt(y)));

const edgeCount = (input.length + input[0].length - 2) * 2;
let internalCount = 0;
for (let x = 1; x < input.length - 1; x++) {
  for (let y = 1; y < input[0].length - 1; y++) {
    const height = input[x][y];
    let visible = true;
    for (let xx = 0; xx < x; xx++) {
      if (input[xx][y] >= height) {
        visible = false;
        break;
      }
    }
    if (visible) {
      internalCount++;
      continue;
    }
    visible = true;
    for (let xx = x + 1; xx < input.length; xx++) {
      if (input[xx][y] >= height) {
        visible = false;
        break;
      }
    }
    if (visible) {
      internalCount++;
      continue;
    }
    visible = true;
    for (let yy = 0; yy < y; yy++) {
      if (input[x][yy] >= height) {
        visible = false;
        break;
      }
    }
    if (visible) {
      internalCount++;
      continue;
    }
    visible = true;
    for (let yy = y + 1; yy < input[0].length; yy++) {
      if (input[x][yy] >= height) {
        visible = false;
        break;
      }
    }
    if (visible) {
      internalCount++;
      continue;
    }
  }
}

console.log("part1 total visible", internalCount + edgeCount);

let maxScenicScore = 0;
for (let x = 1; x < input.length - 1; x++) {
  for (let y = 1; y < input[0].length - 1; y++) {
    const height = input[x][y];
    let left = 0;
    for (let xx = x - 1; xx >= 0; xx--) {
      left++;
      if (input[xx][y] >= height) {
        break;
      }
    }
    let right = 0;
    for (let xx = x + 1; xx < input.length; xx++) {
      right++;
      if (input[xx][y] >= height) {
        break;
      }
    }
    let up = 0;
    for (let yy = y + 1; yy < input[0].length; yy++) {
      up++;
      if (input[x][yy] >= height) {
        break;
      }
    }
    let down = 0;
    for (let yy = y - 1; yy >= 0; yy--) {
      down++;
      if (input[x][yy] >= height) {
        break;
      }
    }
    const scenicScore = left * right * down * up;
    if (scenicScore > maxScenicScore) {
      maxScenicScore = scenicScore;
    }
  }
}

console.log("part2 max scenic score", maxScenicScore);

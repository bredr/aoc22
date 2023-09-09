import { readFileSync } from "fs";
const file = readFileSync("./src/day9input", { encoding: "utf-8" });
const lines = file.split("\n");

let [Tx, Ty] = [0, 0];
let [Hx, Hy] = [0, 0];
const history: string[] = ["0,0"];
lines.forEach((line) => {
  const [dir, d] = line.split(" ");
  const distance = parseInt(d);
  Array(distance)
    .fill(0)
    .forEach(() => {
      switch (dir) {
        case "R":
          Hx += 1;
          break;
        case "L":
          Hx -= 1;
          break;
        case "U":
          Hy += 1;
          break;
        case "D":
          Hy -= 1;
          break;
        default:
          break;
      }
      if (Hx - 1 <= Tx && Tx <= Hx + 1 && Hy - 1 <= Ty && Ty <= Hy + 1) {
        return;
      }
      const dx = Hx - Tx;
      const dy = Hy - Ty;
      if(dy == 0) {
        Tx += dx > 0 ? 1 : -1;
      }
      else if (dx == 0) {
        Ty += dy > 0 ? 1 : -1;
      }
      else {
        Tx += dx > 0 ? 1 : -1
        Ty += dy > 0 ? 1 : -1 
      }
      history.push(`${Tx},${Ty}`)
    });
});

const uniquePlaces = new Set(history)
console.log("places visited",uniquePlaces.size)
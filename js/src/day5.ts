import { readFileSync } from "fs";

const file = readFileSync("./src/day5input", { encoding: "utf-8" });
const [starting, rawProcedure] = file.split("\n\n");

const procedure = rawProcedure
  .split("\n")
  .map((s) => s.match(/[0-9]+/g)?.map((x) => parseInt(x)) ?? [])
  .map(([move, from, to]) => ({ move, from: from - 1, to: to - 1 }));

const rows = starting.split("\n");
rows.reverse();
const maxColumn = Math.max(
  ...rows[0]
    .trim()
    .split("  ")
    .map((x) => parseInt(x.trim()))
);

const rowsCleaned = rows.slice(1).map((row) =>
  Array(maxColumn)
    .fill(0)
    .map((_, idx) =>
      row
        .slice(idx * 4, idx * 4 + 4)
        .trim()
        .replace(/[\[\]]/g, "")
    )
);

const startingColumns = Array(maxColumn)
  .fill([])
  .map((_, idx) =>
    rowsCleaned.reduce(
      (acc, row) => (row[idx] !== "" ? [...acc, row[idx]] : acc),
      []
    )
  );
const columns = startingColumns.map((column) => column.map((x) => x));
// Part 1
procedure.forEach(({ move, from, to }) => {
  Array(move)
    .fill(null)
    .forEach(() => {
      const box = columns[from].pop();
      if (box) {
        columns[to].push(box);
      }
    });
});

console.log(
  "part1",
  columns.map((column) => column[column.length - 1]).join("")
);

// Part 2
const columns2 = startingColumns.map((column) => column.map((x) => x));
procedure.forEach(({ move, from, to }) => {
  const boxes = columns2[from].splice(columns2[from].length - move, move);
  columns2[to].push(...boxes);
});
console.log(
  "part2",
  columns2.map((column) => column[column.length - 1]).join("")
);

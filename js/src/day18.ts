import { readFileSync } from "fs";
import { isEqual, maxBy, minBy, range } from "lodash";

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
console.log("part1", part1);

const findOutsiders = (): {
  outside: Array<[number, number, number]>;
  inside: Array<[number, number, number]>;
  airBubbles: Array<[number, number, number]>;
} => {
  const maxX = maxBy(cubes, ([x]) => x)![0] + 1;
  const maxY = maxBy(cubes, ([_x, y]) => y)![1] + 1;
  const maxZ = maxBy(cubes, ([_x, _y, z]) => z)![2] + 1;
  const minX = minBy(cubes, ([x]) => x)![0] - 1;
  const minY = minBy(cubes, ([_x, y]) => y)![1] - 1;
  const minZ = minBy(cubes, ([_x, _y, z]) => z)![1] - 1;
  const start = [minX, minY, minZ];

  let outside = new Set<string>();
  let seen = new Set<string>();
  let queue = [start];
  while (queue.length > 0) {
    const next = queue.pop();
    if (next) {
      const [x, y, z] = next;
      if (
        x < minX ||
        x > maxX ||
        y < minY ||
        y > maxY ||
        z < minZ ||
        z > maxZ
      ) {
        continue;
      }
      if (outside.has(JSON.stringify(next))) {
        continue;
      }
      if (cubes.some((x) => isEqual(x, next))) {
        outside.add(JSON.stringify(next));
        continue;
      }
      if (seen.has(JSON.stringify(next))) {
        continue;
      }
      seen.add(JSON.stringify(next));
      queue = [
        ...queue,
        [x + 1, y, z],
        [x - 1, y, z],
        [x, y + 1, z],
        [x, y - 1, z],
        [x, y, z + 1],
        [x, y, z - 1],
      ];
    }
  }
  const inside = cubes.filter((x) => !outside.has(JSON.stringify(x)));
  const s = [...seen].map((x) => JSON.parse(x)) as Array<
    [number, number, number]
  >;

  const notAir = [...s, ...cubes];
  const all = range(minX, maxX).flatMap((x) =>
    range(minY, maxY).flatMap((y) =>
      range(minZ, maxZ).map((z) => [x, y, z] as [number, number, number])
    )
  );
  const airBubbles = all.filter((x) => !notAir.some((y) => isEqual(x, y)));
  return {
    outside: [...outside].map((x) => JSON.parse(x)),
    inside,
    airBubbles,
  };
};

const { airBubbles } = findOutsiders();
console.log("air bubbles", airBubbles.length)
const airPlusCubes = [...airBubbles, ...cubes];
const part2 = airPlusCubes.reduce(
  (yy, a, idx) =>
    yy +
    airPlusCubes.reduce(
      (xx, b, idy) => xx - (idx !== idy && isAdjacent(a, b) ? 1 : 0),
      6
    ),
  0
);
console.log("part2", part2);

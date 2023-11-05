import { readFileSync } from "fs";
import { isEqual, maxBy, minBy, range } from "lodash";

const input = readFileSync("src/day18input", { encoding: "utf-8" }).trim();
const cubes = input
  .split("\n")
  .map((x) => x.split(",").map((y) => parseInt(y, 10))) as Array<
  [number, number, number]
>;

const solve1 = (cubes: Array<[number, number, number]>): number =>
  cubes.reduce(
    (xx, [x, y, z]) =>
      [
        [x + 1, y, z],
        [x - 1, y, z],
        [x, y + 1, z],
        [x, y - 1, z],
        [x, y, z + 1],
        [x, y, z - 1],
      ]
        .filter((n) => !cubes.some((c) => isEqual(c, n)))
        .reduce((a) => a + 1, xx),
    0
  );

console.log("part1", solve1(cubes));

const findWhereWaterCantGo = (): {
  lava: Array<[number, number, number]>;
} => {
  const maxX = maxBy(cubes, ([x]) => x)![0] + 1;
  const maxY = maxBy(cubes, ([_x, y]) => y)![1] + 1;
  const maxZ = maxBy(cubes, ([_x, _y, z]) => z)![2] + 1;
  const minX = minBy(cubes, ([x]) => x)![0] - 1;
  const minY = minBy(cubes, ([_x, y]) => y)![1] - 1;
  const minZ = minBy(cubes, ([_x, _y, z]) => z)![2] - 1;
  const start = [minX, minY, minZ];

  let seen = new Set<string>();
  let queue = [start];
  while (queue.length > 0) {
    const next = queue.pop();
    if (next) {
      const [x, y, z] = next;
      if (cubes.some((x) => isEqual(x, next))) {
        continue;
      }
      if (seen.has(JSON.stringify(next))) {
        continue;
      }
      seen.add(JSON.stringify(next));
      const neighbours = [
        [x + 1, y, z],
        [x - 1, y, z],
        [x, y + 1, z],
        [x, y - 1, z],
        [x, y, z + 1],
        [x, y, z - 1],
      ];
      queue = [
        ...queue,
        ...neighbours.filter(
          ([x, y, z]) =>
            minX <= x &&
            x <= maxX &&
            minY <= y &&
            y <= maxY &&
            minZ <= z &&
            z <= maxZ
        ),
      ];
    }
  }
  const all = range(minX, maxX + 1).flatMap((x) =>
    range(minY, maxY + 1).flatMap((y) =>
      range(minZ, maxZ + 1).map((z) => [x, y, z] as [number, number, number])
    )
  );
  const lava = all.filter((x) => !seen.has(JSON.stringify(x)));
  return {
    lava,
  };
};

const { lava } = findWhereWaterCantGo();
console.log("part2", solve1(lava));

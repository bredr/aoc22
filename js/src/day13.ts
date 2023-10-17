import { readFileSync } from "fs";

const input = readFileSync("src/day13input", { encoding: "utf-8" }).trim();

const compare = (left: number, right: number): number => {
    if(left < right) {
        return 1
    }
    if(left > right) {
        return -1
    }
    return 0
}

const comparePackets = (left: Array<Array<number> | number> | number, right: Array<Array<number> | number> | number): number => {
    if(right === undefined) {
        return -1;
    }
    if(!Array.isArray(left) && !Array.isArray(right)) {
        return compare(left, right)
    }
    if(!Array.isArray(left) && Array.isArray(right)) {
        return comparePackets([left], right)
    }
    if(Array.isArray(left) && !Array.isArray(right)) {
        return comparePackets(left, [right])
    }
    if(Array.isArray(left) && Array.isArray(right)) {
        const zipped = left.map((l, ix) => ([l, right[ix]]))
        for(const [l, r] of zipped) {
            const res = comparePackets(l, r)
            if(res != 0) {
                return res
            }
        }
        return comparePackets(left.length, right.length)
    }
    throw Error("shouldn't reach here")
}

const findIndex= (l: Array<Array<number>>, item: Array<Array<number>>): number => l.flat().reduce((acc, pkt) => acc + (comparePackets(pkt, item) == 1? 1 : 0), 0)

const part1 = input.split("\n\n").map((pair, idx) => {
  const index = idx + 1;
  const [left, right] = pair.split("\n").map(eval);


  if(comparePackets(left,right)== 1) {
    return index;
  } else {
    return 0;
  }
}).reduce((xx, x) => xx + x, 0);

const pairs = input.split("\n\n").map((pair, idx) => {
    const [left, right] = pair.split("\n").map(eval);
    return [left, right]
});

const part2 = (findIndex(pairs, [[2]]) +1) * (findIndex(pairs, [[6]])+2)
  

console.log("part1" ,part1)
console.log("part2" ,part2)

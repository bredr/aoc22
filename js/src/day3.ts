import { readFileSync } from "fs"

const file = readFileSync("./src/day3input", { encoding: "utf-8" })
const lines = file.split("\n")

const toNumber = (x: string) => {
    if (x.toLowerCase() === x) {
        return x.charCodeAt(0) - 96
    }
    return x.toLowerCase().charCodeAt(0) - 96 + 26
}
const sum = lines.reduce((acc, line) => {
    const comp1 = new Set(line.slice(0, line.length / 2).split("").map(toNumber));
    const comp2 = new Set(line.slice(line.length / 2).split("").map(toNumber));
    const intersect = [...comp1].filter(i => comp2.has(i));
    return acc + intersect[0]
}, 0)

console.log("part 1 total =", sum)

const sum2 = Array(lines.length / 3).fill(0).map((_, idx) => lines.slice(idx * 3, idx * 3 + 3)).reduce((acc, group) => {
    const A = new Set(group[0].split("").map(toNumber));
    const B = new Set(group[1].split("").map(toNumber));
    const C = new Set(group[2].split("").map(toNumber));
    const intersect = <T>(set1: Set<T>, set2: Set<T>) => new Set([...set1].filter(i => set2.has(i)));
    return acc + [...intersect(A, intersect(B, C))][0]
}, 0)
console.log("part 2 total =", sum2)

import {readFileSync} from "fs"

const file = readFileSync("./src/day3input", {encoding: "utf-8"})
const lines = file.split("\n")

const toNumber = (x: string) => {
    if(x.toLowerCase() === x) {
        return x.charCodeAt(0)-96
    }
    return x.toLowerCase().charCodeAt(0) - 96 + 26
}
const sum = lines.reduce((acc, line) => {
    const comp1 = new Set(line.slice(0, line.length / 2).split("").map(toNumber));
    const comp2 = new Set(line.slice(line.length/2).split("").map(toNumber));
    const intersect = [...comp1].filter(i => comp2.has(i));
    return acc + intersect[0]
}, 0) 

console.log("part 1 total =", sum)
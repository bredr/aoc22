import { readFileSync } from "fs"
const file = readFileSync("./src/day7input", { encoding: "utf-8" })
const lines = file.split("\n")

const sizes: Record<string, number> ={}
const stack: string[] = []
lines.forEach(line => {
    if(line.startsWith("$ ls") || line.startsWith("dir")) {
        return;
    }
    if(line.startsWith("$ cd")) {
        const dest = line.replace("$ cd ", "")
        if(dest == "..") {
            stack.pop()
        } else {
            const path = stack.length > 0 ?`${stack[stack.length -1]}_${dest}` : dest
            stack.push(path)
        }
    }else {
        const [size] = line.split(" ")
        stack.forEach(path => {
            if(path in sizes) {
            sizes[path] += parseInt(size)
            }  else {
                sizes[path] =  parseInt(size)
            }
        })
    }
})

console.log(sizes)
const total = Object.values(sizes).filter(x => x<=100000).reduce((xx,x)=> xx+x, 0)
console.log("part1 sum", total)
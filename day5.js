const { Console } = require("node:console");
const fs = require("node:fs/promises")

const day5_task1 = async function(){
    let depMap = new Map();
    let pages = [];
    const f = await fs.open('./inputs/5.txt');
    let dep = true;
    for await (let line of f.readLines()){
        if (line == ""){
            dep = false;
            continue;
        }
        if (dep){
            const kp = line.split('|')
            if (depMap.get(kp[0])){
                depMap.get(kp[0]).add(kp[1]);
            } else {
                depMap.set(kp[0], new Set([kp[1]]));
            }
        } else {
            pages.push(line);
        }
    }

    const sum = processPages(pages, depMap);
    console.log(`Valid pages sum: ${sum}`);
}

function processPages(pages, depMap){
    let sum = 0;
    let validPages = 0;
    for (let i = 0; i < pages.length; i++){
        const pageNumbers = pages[i].split(',');
        if (isValidPage(pageNumbers, depMap)){
            const midPageIndex = Math.floor(pageNumbers.length / 2);
            sum += Number(pageNumbers[midPageIndex]);
            validPages++;
            //console.log(`Valid Page: ${pageNumbers}`);
        }
    }
    console.log(`Total Page: ${pages.length}, Valid Pages: ${validPages}`);
    return sum;
}

function isValidPage(pages, depMap){
    //console.log(pages)
    for(let i = pages.length -1; i >= 1; i--) {
        
        const followPages = depMap.get(pages[i]);
        //console.log(pages[i], followPages)
        if (followPages){
            if (hasFollowingPages(pages, followPages, i)){
                return false;
            }
        }
    }
    return true;
}

function hasFollowingPages(pages, followPages, i){
    for(let j = 0; j < i; j++){
        if (followPages.has(pages[j])) {
            //console.log(pages, pages[j], pages[i]);
            return true;
        }
    }
    return false;
}

module.exports = {
    day5_task1
}
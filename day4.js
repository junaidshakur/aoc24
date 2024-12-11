const fs = require('node:fs/promises')

const day4_task1 = async function(){

    const lines = []
    const f = await fs.open('./inputs/4.txt')
    for await (const line of f.readLines()){
        lines.push(line)
    }
    //const wordCount = findXMAS(lines);
    const wordCount = findCrossMAS(lines);
    console.log(`Word count of XMAS: ${wordCount}`);
}

function findCrossMAS(matrix){
    let count = 0;

    for (let i = 0; i < matrix.length; i++){
        let lineLength = matrix[i].length;
        for(let k = 0; k < lineLength; k++){
            if (matrix[i][k] == 'A'){
                count += isCrossMS(matrix, i, k);
            }
        }
    }

    return count;

}

function isCrossMS(matrix, i, k){
    let topLeft, downLeft, topRight, downRight;

    if (i-1 >= 0 && k-1 >= 0){
        topLeft = matrix[i-1][k-1];
    }
    if (i-1 >= 0 && k+1 < matrix[i-1].length){
        topRight = matrix[i-1][k+1];
    }
    if (i+1 < matrix.length && k-1 >= 0){
        downLeft = matrix[i+1][k-1];
    }
    if (i+1 < matrix.length && k+1 < matrix[i+1].length){
        downRight = matrix[i+1][k+1];
    }

    if ((topLeft == 'M' && downRight == 'S') || (topLeft == 'S' && downRight == 'M')){
        if ((downLeft == 'M' && topRight == 'S') || (downLeft == 'S' && topRight == 'M')){
            return 1;
        }
    }
    return 0;
}

function findXMAS(matrix) {
    let count = 0;

    for (let i = 0; i < matrix.length; i++){
        let lineLength = matrix[i].length;
        for(let k = 0; k < lineLength; k++){
            if (matrix[i][k] == 'X'){
                count += findUp(matrix, i, k);
                count += findDown(matrix, i, k);
                count += findLeft(matrix, i, k);
                count += findRight(matrix, i, k);
                count += findUpLeft(matrix, i, k);
                count += findUpRight(matrix, i, k);
                count += findDownLeft(matrix, i, k);
                count += findDownRight(matrix, i, k);
            }
        }
    }

    return count;
}

function findUp(matrix, i, k){
    if (i -1 >= 0){
        if (matrix[i-1][k] == 'M'){
            if (i - 2 >= 0){
                if (matrix[i-2][k] == 'A'){
                    if (i -3 >= 0){
                        if (matrix[i-3][k] == 'S'){
                            return 1;
                        }
                    }
                }
            }
        }
    }
    return 0;
}

function findDown(matrix, i, k){
    if (i +1 < matrix.length){
        if (matrix[i+1][k] == 'M'){
            if (i + 2 < matrix.length){
                if (matrix[i+2][k] == 'A'){
                    if (i +3 < matrix.length){
                        if (matrix[i+3][k] == 'S'){
                            return 1;
                        }
                    }
                }
            }
        }
    }
    return 0;
}

function findLeft(matrix, i, k){
    if (k -1 >= 0){
        if (matrix[i][k-1] == 'M'){
            if (k - 2 >= 0){
                if (matrix[i][k-2] == 'A'){
                    if (k -3 >= 0){
                        if (matrix[i][k-3] == 'S'){
                            return 1;
                        }
                    }
                }
            }
        }
    }
    return 0;
}

function findRight(matrix, i, k){
    if (k+1 < matrix[i].length){
        if (matrix[i][k+1] == 'M'){
            if (k+2 < matrix[i].length){
                if (matrix[i][k+2] == 'A'){
                    if (k+3 < matrix[i].length){
                        if (matrix[i][k+3] == 'S'){
                            return 1;
                        }
                    }
                }
            }
        }
    }
    return 0;
}

function findUpRight(matrix, i, k){
    if (i -1 >= 0 && k+1 < matrix[i-1].length){
        if (matrix[i-1][k+1] == 'M'){
            if (i - 2 >= 0 && k+2 < matrix[i-2].length){
                if (matrix[i-2][k+2] == 'A'){
                    if (i -3 >= 0 && k+3 < matrix[i-3].length){
                        if (matrix[i-3][k+3] == 'S'){
                            return 1;
                        }
                    }
                }
            }
        }
    }
    return 0;
}

function findUpLeft(matrix, i, k){
    if (i -1 >= 0 && k-1 >= 0){
        if (matrix[i-1][k-1] == 'M'){
            if (i - 2 >= 0 && k-2 >= 0){
                if (matrix[i-2][k-2] == 'A'){
                    if (i -3 >= 0 && k-3 >= 0){
                        if (matrix[i-3][k-3] == 'S'){
                            return 1;
                        }
                    }
                }
            }
        }
    }
    return 0;

}

function findDownRight(matrix, i, k){
    if (i+1 < matrix.length && k+1 < matrix[i+1].length){
        if (matrix[i+1][k+1] == 'M'){
            if (i+2 < matrix.length && k+2 < matrix[i+2].length){
                if (matrix[i+2][k+2] == 'A'){
                    if (i+3 < matrix.length && k+3 < matrix[i+3].length){
                        if (matrix[i+3][k+3] == 'S'){
                            return 1;
                        }
                    }
                }
            }
        }
    }
    return 0;

}

function findDownLeft(matrix, i, k){
    if (i+1 < matrix.length && k-1 >= 0){
        if (matrix[i+1][k-1] == 'M'){
            if (i+2 < matrix.length && k-2 >= 0){
                if (matrix[i+2][k-2] == 'A'){
                    if (i+3 < matrix.length && k-3 >= 0){
                        if (matrix[i+3][k-3] == 'S'){
                            return 1;
                        }
                    }
                }
            }
        }
    }
    return 0;

}

module.exports = {
    day4_task1
}
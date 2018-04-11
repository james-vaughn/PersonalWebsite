//following https://jonnoftw.github.io/2017/01/18/markov-chain-image-generation
class MarkovChain {
    constructor() {
        // pixel => neighbor pixels => frequency
        this.mapping = [];
        this.state = null; //state for iterating with
    }

    put(currColor, neighborColor) {
        for(let pair of this.mapping) {
            if(pair.left.equals(currColor)) {
                for(let neighbor of pair.right) {
                    if(neighbor.left.equals(neighborColor)) {
                        neighbor.right += 1;
                        return;
                    }
                }

                pair.right.push(new Pair(neighborColor, 1));
                return;

            }
        }

        this.mapping.push(new Pair(currColor, [new Pair(neighborColor, 1)]));
    }

    next(currColor) {
        //for
        return currColor;
    }
}

class Pair {
    constructor(left, right) {
        this.left = left;
        this.right = right;
    }
}
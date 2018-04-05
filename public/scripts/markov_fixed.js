//following https://jonnoftw.github.io/2017/01/18/markov-chain-image-generation
class MarkovChain {
    constructor() {
        // pixel => neighbor pixels => frequency
        this.mapping = {};
    }

    add(currColor, neighborColor) {
        if(this.mapping[currColor][neighborColor] === undefined) {
            this.mapping[currColor][neighborColor] = 1;
        } else {
            this.mapping[currColor][neighborColor] += 1;
        }
    }
}
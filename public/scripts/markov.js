class MarkovNode {
    constructor(value) {
        this.value = value;
        this.freq = 0;
        this.neighbors = [];
    }

    addNeighbor(neighbor) {
        neighbor.freq = 1;
        this.neighbors.push(neighbor);
        return neighbor;
    }

    adjustNeighbor(neighborVal, amt = 1) {
        let neighborNode = this.neighborWithValue(neighborVal);

        if (neighborNode === undefined) {
            neighborNode = new MarkovNode(neighborVal);
            this.addNeighbor(neighborNode);
        } else {
            const currAmt = neighborNode.freq;
            neighborNode.freq = Math.max(0, currAmt + amt);
        }

        return neighborNode;
    }

    neighborWithValue(value) {
        for (let node of this.neighbors) {
            if (node.value.equals(value)) {
                return node;
            }
        }

        return undefined;
    }

    next() {
        let total = 0;
        for(let node of this.neighbors) {
            //total += node.freq;
            total += 1; //random walk
        }

        let chosen = Math.ceil(total * Math.random());

        for(let node of this.neighbors) {
            //chosen -= node.freq;
            chosen -= 1; //random walk

            if(chosen <= 0) {
                // console.log(node.value);
                return node;
            }
        }

        return undefined;
    }

    //kinda doesnt work
    print() {
        let visited = [];
        let queue = [this];

        while (queue.length > 0) {
            let neighbor = queue.shift();
            visited.push(neighbor);

            console.log(neighbor);

            outer:
            for(let node of neighbor.neighbors) {
                for(let vNode of visited) {
                    if(vNode.value.equals(node.value)) {
                    // if(vNode === node) {
                        continue outer;
                    }
                }

                queue.push(node);
            }
        }
    }
}
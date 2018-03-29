class MarkovNode {
    constructor(value) {
        this.value = value;
        this.freq = 0;
        this.neighbors = [this];
    }

    addNeighbor(neighbor, amt = 1) {
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
            total += node.freq;
        }

        let chosen = Math.floor(total * Math.random());

        for(let node of this.neighbors) {
            chosen -= node.freq;

            if(chosen <= 0) {
                return node;
            }
        }

        return undefined;
    }

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
                        continue outer;
                    }
                }

                queue.push(node);
            }
        }
    }
}
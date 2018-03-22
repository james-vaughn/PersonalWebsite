class MarkovNode {
    constructor(value) {
        this.value = value
        this.neighbors = {};
    }

    addNeighbor(neighbor, amt = 1) {
        this.neighbors[neighbor] = amt;
        return neighbor;
    }

    adjustNeighbor(neighborVal, amt = 1) {
        let neighborNode = this.neighborWithValue(neighborVal);

        if (neighborNode === undefined) {
            neighborNode = new MarkovNode(neighborVal);
            this.addNeighbor(neighborNode);
        } else {
            const currAmt = this.neighbors[neighborNode];
            this.neighbors[neighborNode] = Math.max(0, currAmt + amt);
        }

        return neighborNode;
    }

    neighborWithValue(value) {
        for (let node in this.neighbors) {
            if (node.value === value) {
                return node;
            }
        }

        return undefined;
    }

    next() {

    }

    print() {
        let visited = [];

        for (let neighbor in Object.keys(this.neighbors)) {
            console.log(neighbor);
            if(visited.indexOf(neighbor) !== -1) continue;

            visited.push(neighbor);
            console.log(neighbor.value);

            neighbor.print();
        }
    }
}
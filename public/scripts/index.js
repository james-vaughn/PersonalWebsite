$(document).ready(function () {
    const submitFileInput = $("#submitfile");
    const originalId = $("#original");
    const newId = $("#new");
    const width = originalId.prop("width");
    const height = originalId.prop("height");

    //see https://www.nczonline.net/blog/2012/05/15/working-with-files-in-javascript-part-2/
    submitFileInput.on("change", function() {
        if (this.files && this.files[0]) {
            const file = this.files[0];

            const reader = new FileReader();
            reader.onload = function(event) {
                const dataUri = event.target.result;
                const originalContext = originalId[0].getContext("2d");

                const img = new Image();
                img.onload = function() {
                    originalContext.drawImage(img, 0, 0, width, height);
                    // colors(originalContext);
                    createMarkovImage(originalContext, newId[0].getContext("2d"));
                };

                img.src = dataUri;
            };

            reader.onerror = function(event) {
                console.error("File could not be read! Code " + event.target.error.code);
            };
            reader.readAsDataURL(file);
        }
    });

    function createMarkovImage(oldContext, newContext) {
        // let iterNode = createMarkovChain(oldContext, newContext);
        let chain = createMarkovChain(oldContext, newContext);


        // for (let x = 0; x < width; x++) {
        //     for (let y = 0; y < height; y++) {
        //         iterNode = iterNode.next();
        //         drawPixel(newContext, x, y, iterNode.value);
        //     }
        // }
    }

    function colors(ctx) {
        let pixelData = ctx.getImageData(0, 0, width, height).data;
        let colors = [];

        outer:
        for (let i = 0, n = pixelData.length; i < n; i += 4) {
            let color = new Color(pixelData[i], pixelData[i + 1], pixelData[i + 2], pixelData[i + 3]);
            for(let c of colors) {
                if(c.equals(color)) {
                    continue outer;
                }
            }
            colors.push(color);
        }
        console.log(colors);
    }

    function createMarkovChain(oldContext, newContext) {
        // let headNode = new MarkovNode(new Color(0,0,0,1)); //todo
        // let iterNode = headNode;
        let chain = new MarkovChain();

        let pixelData = oldContext.getImageData(0, 0, width, height).data;

        for (let h = 0; h < height; h++) {
            for (let w = 0; w < width; w++) {

                let color = colorFromPos(pixelData, w-1, h);
                adjustNeighbor(iterNode, color);

                color = colorFromPos(pixelData, w-1, h-1);
                adjustNeighbor(iterNode, color);

                color = colorFromPos(pixelData, w-1, h+1);
                adjustNeighbor(iterNode, color);

                color = colorFromPos(pixelData, w, h-1);
                adjustNeighbor(iterNode, color);

                color = colorFromPos(pixelData, w, h+1);
                adjustNeighbor(iterNode, color);

                color = colorFromPos(pixelData, w-1, h);
                adjustNeighbor(iterNode, color);

                color = colorFromPos(pixelData, w+1, h);
                adjustNeighbor(iterNode, color);

                color = colorFromPos(pixelData, w+1, h);
                iterNode = adjustNeighbor(iterNode, color);
            }
        }

        return headNode;
    }

    function drawPixel(context, x, y, color) {
        context.fillStyle = `rgba(${color.r}, ${color.g}, ${color.b}, ${color.a})`;
        context.fillRect(x, y, 1, 1);
    }

    function colorFromPos(pixelData, w, h) {
        //clamp w and h to prevent bad values
        w = clamp(w, 0, width);
        h = clamp(h, 0, height);

        const pos = (w, h) => 4 * (h * width + w);
        const pixelIndex = pos(w, h);
        return new Color(pixelData[pixelIndex], pixelData[pixelIndex+1], pixelData[pixelIndex+2], 1);
    }

    function adjustNeighbor(markovNode, color) {
        //if the color will be visible
        if(color.a > 0) {
            return markovNode.adjustNeighbor(color);
        }
    }

    function clamp(x, min, max) {
        return Math.min(max, Math.max(min, x));
    }
});
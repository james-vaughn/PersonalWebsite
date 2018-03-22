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
                    drawImageScaled(img, originalContext);

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

    function drawImageScaled(img, ctx) {
        var canvas = ctx.canvas ;
        var hRatio = canvas.width  / img.width    ;
        var vRatio =  canvas.height / img.height  ;
        var ratio  = Math.min ( hRatio, vRatio );
        var centerShift_x = ( canvas.width - img.width*ratio ) / 2;
        var centerShift_y = ( canvas.height - img.height*ratio ) / 2;
        ctx.clearRect(0,0,canvas.width, canvas.height);
        ctx.drawImage(img, 0,0, img.width, img.height,
            centerShift_x,centerShift_y,img.width*ratio, img.height*ratio);
    }

    function createMarkovImage(oldContext, newContext) {
        let iterNode = createMarkovChain(oldContext, newContext);
        iterNode.print();
        newContext.putImageData(oldContext.getImageData(0, 0 , width, height),0, 0);
    }

    function createMarkovChain(oldContext, newContext) {
        let headNode = new MarkovNode(new Color(0,0,0,1)); //todo
        let iterNode = headNode;

        let pixelData = oldContext.getImageData(0, 0, width, height).data;

        for (let i = 0, n = pixelData.length; i < n; i += 4) {
            let color = new Color(pixelData[i], pixelData[i+1], pixelData[i+2], pixelData[i+3]);
            iterNode = iterNode.adjustNeighbor(color);
        }

        return headNode;
    }

    function drawPixel(context, x, y, color) {
        context.fillStyle = `rgba(${color.r}, ${color.g}, ${color.b}, ${color.a})`;
        context.fillRect(x, y, 1, 1);
    }
});



$(document).ready(function () {
    const submitFileInput = $("#submitfile");
    const originalId = $("#original");
    const newId = $("#new");

    //see https://www.nczonline.net/blog/2012/05/15/working-with-files-in-javascript-part-2/
    submitFileInput.on("change", function() {
        if (this.files && this.files[0]) {
            const file = this.files[0];

            const reader = new FileReader();
            reader.onload = function(event) {
                const dataUri = event.target.result;
                const originalContext = original.getContext("2d");

                const img = new Image();
                img.onload = function() {
                    drawImageScaled(img, originalContext);

                    createMarkovImage(originalContext, newId.getContext("2d"));
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

    }
});



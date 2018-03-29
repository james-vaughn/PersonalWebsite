class Color {
    //r,g,b range 0..255
    //a range 0..1
    constructor(r, g, b, a) {
        this.r = r;
        this.g = g;
        this.b = b;
        this.a = a;
    }

    equals(color) {
        return  this.r === color. r &&
                this.g === color.g &&
                this.b === color.b &&
                this.a === color.a;
    }
}
class Sketchpad {
    constructor(width,height,canvasID,rainbowMode) {
        rainbowMode = rainbowMode || false
        this.rainbowMode = rainbowMode

        this.colourIndex = rainbowMode? 2 : 1
        
        this.width = width
        this.height = height

        this.canvas = document.getElementById(canvasID)
        this.ctx = this.canvas.getContext("2d")

        this.canvas.width = width
        this.canvas.height = height

        this.mousedown = false
        this.canvas.addEventListener("mousedown",function(){this.mousedown=true}.bind(this))
        this.canvas.addEventListener("mouseup",function(){this.mousedown=false}.bind(this))

        this.lastX = -1
        this.lastY = -1

        this.canvas.addEventListener("mousemove",this.drawTo.bind(this))
        this.canvas.addEventListener("mousedown",this.drawPix.bind(this))
        this.canvas.addEventListener("pointerleave",this.resetMousePos.bind(this))

        /*Saturation should be in range 0-255 inclusive*/
        this.saturation = 255

        /*imageData is in a format that's ready to be WHOOSHED to the server*/
        this.imageData = {"span":width,"data":new Array(width*height).fill(0)}
    }

    resetMousePos(e) {
        this.lastX = -1
        this.lastY = -1
    }

    loadImageData(imageData) {
        for (var i = 0; i < imageData["data"].length; i++){
            this.setPixel(i%imageData["span"],Math.floor(i/imageData["span"]),imageData["data"][i])
        }
    }

    setPixel(x,y,i) {
        [x, y] = [Math.round(x),Math.round(y)]
        this.imageData["data"][y*this.width+x] = i
        this.ctx.fillStyle=this.getColour(i)
        this.ctx.fillRect(x,y,1,1)
    }

    getColour(i) {
        switch (i) {
            case 0:
                return "#FFFFFF"
            case 1:
                return "#000000"
            default:
                var toHex = (r,g,b) => {
                    return "#"+Math.round(r).toString(16).padStart(2,"0")
                              +Math.round(g).toString(16).padStart(2,"0")
                              +Math.round(b).toString(16).padStart(2,"0")
                }
        
                var up = i => (255-this.saturation)+(this.saturation*(i%1))
                var down = i => (255-this.saturation)+(this.saturation*(1-(i%1)))
                var flat = () => 255
        
                var nui = ((i-2)/254)*3
        
                switch (Math.floor(nui)){
                    case 0:
                        return toHex(flat(nui),up(nui),down(nui))
                    case 1:
                        return toHex(down(nui),flat(nui),up(nui))
                    case 2:
                        return toHex(up(nui),down(nui),flat(nui))
                }            
        }
    }

    drawTo(event) {
        var [x, y] = [event.clientX/this.canvas.clientWidth*this.canvas.width, 
                      event.clientY/this.canvas.clientHeight*this.canvas.height]

        if (this.mousedown && !(this.lastX == -1 && this.lastY == -1)){
            var deltas = [x-this.lastX, y-this.lastY]
            var dist = Math.sqrt(deltas[0]*deltas[0]+deltas[1]*deltas[1])

            for (var i = 0; i < dist; i += 0.1) {
                this.setPixel(x-(deltas[0]*(i/dist)),y-(deltas[1]*(i/dist)),this.colourIndex)
            }    

            if (this.rainbowMode) {
                this.colourIndex = ((this.colourIndex+1)%254)+2
            }
        }

        [this.lastX, this.lastY] = [x, y]
    }

    drawPix(event) {
        var [x, y] = [event.clientX/this.canvas.clientWidth*this.canvas.width, 
            event.clientY/this.canvas.clientHeight*this.canvas.height]

        this.setPixel(x,y,this.colourIndex)

        if (this.rainbowMode) {
            this.colourIndex = ((this.colourIndex+1)%254)+2
        }
    }
}
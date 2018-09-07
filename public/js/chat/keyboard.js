module.exports = class Keyboard {
  constructor(dcb) {
    this.loadLayer(0);
    this.keyCallback = dcb;

    $('.key').on('click', function(e) {
      var keyidx = $(e.target).attr('data-keyidx')
      if (!keyidx) {
        if ($(e.target).hasClass('space')) {
          var keytype = "space"
        } else if ($(e.target).hasClass('bkspace')) {
          var keytype = "bkspace"
        }
        var data = false;
      } else {
        var keytype = "char";
        var data = keydata[this.currentLayer].chars[keyidx]
      }

      this.keyCallback(keytype, data)

    }.bind(this));
  }

  loadLayer(l) {
    this.currentLayer = l;
    var layer = keydata[l];
    console.log('KBD - Loading Layer ' + l + ': ' + layer.name);
    for (let c = 0; c < layer.chars.length; c++) {
      let char = layer.chars[c];
      $('.key.idx-' + c).html(char.char)
    }
  }
}

const keydata = [
  {
    name: 'qwerty',
    chars: [
      {
        "char": "q",
        "width": "5",
        "height": 7,
        "data": [
          ".....",
          ".##.#",
          "#..##",
          "#...#",
          ".####",
          "....#",
          "....#"
        ]
      },
      {
        "char": "w",
        "width": "5",
        "height": 7,
        "data": [
          ".....",
          ".....",
          "#...#",
          "#...#",
          "#.#.#",
          "#.#.#",
          ".####"
        ]
      },
      {
        "char": "e",
        "width": "5",
        "height": 7,
        "data": [
          ".....",
          ".....",
          ".###.",
          "#...#",
          "#####",
          "#....",
          ".####"
        ]
      },
      {
        "char": "r",
        "width": "5",
        "height": 7,
        "data": [
          ".....",
          ".....",
          "#.##.",
          "##..#",
          "#....",
          "#....",
          "#...."
        ]
      },
      {
        "char": "t",
        "width": "3",
        "height": 7,
        "data": [
          ".#.",
          ".#.",
          "###",
          ".#.",
          ".#.",
          ".#.",
          "..#"
        ]
      },
      {
        "char": "y",
        "width": "5",
        "height": 7,
        "data": [
          ".....",
          "#...#",
          "#...#",
          "#...#",
          ".####",
          "....#",
          "####."
        ]
      },
      {
        "char": "u",
        "width": "5",
        "height": 7,
        "data": [
          ".....",
          ".....",
          "#...#",
          "#...#",
          "#...#",
          "#...#",
          ".####"
        ]
      },
      {
        "char": "i",
        "width": "1",
        "height": 7,
        "data": [
          "#",
          ".",
          "#",
          "#",
          "#",
          "#",
          "#"
        ]
      },
      {
        "char": "o",
        "width": "5",
        "height": 7,
        "data": [
          ".....",
          ".....",
          ".###.",
          "#...#",
          "#...#",
          "#...#",
          ".###."
        ]
      },
      {
        "char": "p",
        "width": "5",
        "height": 7,
        "data": [
          ".....",
          "#.##.",
          "##..#",
          "#...#",
          "####.",
          "#....",
          "#...."
        ]
      },
      {
        "char": "a",
        "width": "5",
        "height": 7,
        "data": [
          ".....",
          ".....",
          ".###.",
          "....#",
          ".####",
          "#...#",
          ".####"
        ]
      },
      {
        "char": "s",
        "width": "5",
        "height": 7,
        "data": [
          ".....",
          ".....",
          ".####",
          "#....",
          ".###.",
          "....#",
          "####."
        ]
      },
      {
        "char": "d",
        "width": "5",
        "height": 7,
        "data": [
          "....#",
          "....#",
          ".##.#",
          "#..##",
          "#...#",
          "#...#",
          ".####"
        ]
      },
      {
        "char": "f",
        "width": "4",
        "height": 7,
        "data": [
          "..##",
          ".#..",
          "####",
          ".#..",
          ".#..",
          ".#..",
          ".#.."
        ]
      },
      {
        "char": "g",
        "width": "5",
        "height": 7,
        "data": [
          ".....",
          ".####",
          "#...#",
          "#...#",
          ".####",
          "....#",
          "####."
        ]
      },
      {
        "char": "h",
        "width": "5",
        "height": 7,
        "data": [
          "#....",
          "#....",
          "#.##.",
          "##..#",
          "#...#",
          "#...#",
          "#...#"
        ]
      },
      {
        "char": "j",
        "width": "5",
        "height": 7,
        "data": [
          "....#",
          ".....",
          "....#",
          "....#",
          "#...#",
          "#...#",
          ".###."
        ]
      },
      {
        "char": "k",
        "width": "4",
        "height": 7,
        "data": [
          "#...",
          "#...",
          "#..#",
          "#.#.",
          "##..",
          "#.#.",
          "#..#"
        ]
      },
      {
        "char": "l",
        "width": "2",
        "height": 7,
        "data": [
          "#.",
          "#.",
          "#.",
          "#.",
          "#.",
          "#.",
          ".#"
        ]
      },
      {
        "char": "z",
        "width": "5",
        "height": 7,
        "data": [
          ".....",
          ".....",
          "#####",
          "...#.",
          "..#..",
          ".#...",
          "#####"
        ]
      },
      {
        "char": "x",
        "width": "5",
        "height": 7,
        "data": [
          ".....",
          ".....",
          "#...#",
          ".#.#.",
          "..#..",
          ".#.#.",
          "#...#"
        ]
      },
      {
        "char": "c",
        "width": "5",
        "height": 7,
        "data": [
          ".....",
          ".....",
          ".###.",
          "#...#",
          "#....",
          "#...#",
          ".###."
        ]
      },
      {
        "char": "v",
        "width": "5",
        "height": 7,
        "data": [
          ".....",
          ".....",
          "#...#",
          "#...#",
          "#...#",
          ".#.#.",
          "..#.."
        ]
      },
      {
        "char": "b",
        "width": "5",
        "height": 7,
        "data": [
          "#....",
          "#....",
          "#.##.",
          "##..#",
          "#...#",
          "#...#",
          ".###."
        ]
      },
      {
        "char": "n",
        "width": "5",
        "height": 7,
        "data": [
          ".....",
          ".....",
          "####.",
          "#...#",
          "#...#",
          "#...#",
          "#...#"
        ]
      },
      {
        "char": "m",
        "width": "5",
        "height": 7,
        "data": [
          ".....",
          ".....",
          "##.#.",
          "#.#.#",
          "#.#.#",
          "#...#",
          "#...#"
        ]
      },
      {
        "char": ",",
        "width": "1",
        "height": 7,
        "data": [
          ".",
          ".",
          ".",
          ".",
          ".",
          "#",
          "#"
        ]
      },
      {
        "char": ".",
        "width": "1",
        "height": 7,
        "data": [
          ".",
          ".",
          ".",
          ".",
          ".",
          ".",
          "#"
        ]
      },
      {
        "char": "?",
        "width": "5",
        "height": 7,
        "data": [
          ".###.",
          "#...#",
          "....#",
          "...#.",
          "..#..",
          ".....",
          "..#.."
        ]
      },
    ]
  }
]

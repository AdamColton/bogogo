var Game = {
  positions: [],
  turn: 1,
  current: (function(){
    var x,y;
    var s = "";
    for(x=0; x<10; x+=1){
      for(y=0; y<10; y+=1){
        s += "+";
      }
      s += "\n";
    }
    return s
  })(),
  move: function(x,y){
    $.post("/move",{
      id: Game.current,
      x: x,
      y: y
    }, function(data){
      Game.positions.push(Game.current);
      Game.current = data;
      Game.turn = 3 - Game.turn;
      Game.draw();
      Game.joseki();
    });
  },
  draw: function(){
    var i,x,y,p;
    Game.ctx.lineWidth = 0;
    Game.ctx.strokeStyle = "rgb(255,255,255)";
    Game.ctx.fillStyle = "rgb(255,255,255)";
    Game.ctx.rect(0,0,600,600);
    Game.ctx.fill();

    Game.ctx.lineWidth = 1;
    Game.ctx.strokeStyle = "rgb(0,0,0)";
    for(i=0; i<10; i+=1){  
      Game.ctx.beginPath();
      Game.ctx.moveTo(i*42 + 22, 22);
      Game.ctx.lineTo(i*42 + 22, 420);
      Game.ctx.stroke();

      Game.ctx.beginPath();
      Game.ctx.moveTo(22, i*42 + 22);
      Game.ctx.lineTo(420, i*42 + 22);
      Game.ctx.stroke();
    }
    for(x=0; x<10; x+=1){
      for(y=0; y<10; y+=1){
        p = Game.current[11*x+y];
        if (p != "+"){
          if ( (p == "@" && Game.turn == 1) || (p == "O" && Game.turn == 2) ){
            Game.ctx.drawImage(Game.black, x*42+1, y*42+1);
          } else {
            Game.ctx.drawImage(Game.white, x*42+1, y*42+1);
          }
        }
      }
    }
  },
  calculateMove: function(e){
    var x = e.layerX - e.currentTarget.offsetLeft;
    var y = e.layerY - e.currentTarget.offsetTop;
    x = Math.round( (x-22)/42 )
    y = Math.round( (y-22)/42 )
    Game.move(x,y);
  },
  joseki: function(){
    $.post("/joseki", {id: Game.current, turn: Game.turn}, function(data){
      var x,y,p;
      data = data.split(",");
      for(x=0; x<10; x+=1){
        for(y=0; y<10; y+=1){
          p = parseInt(data[10*x+y])/1000;
          Game.ctx.beginPath();
          Game.ctx.fillStyle = "rgba(0,255,0,"+p+")";
          Game.ctx.rect(x*42,y*42,42,42);
          Game.ctx.fill();
        }
      }
    });
  }
}

/*
pieces are 42x42
*/

$(function(){
  var c = 0;
  var f = function(){
    c += 1;
    if (c === 2){
      Game.draw();
      Game.joseki();
    }
  };

  var canvas = document.getElementById("game");
  canvas.onclick = Game.calculateMove;

  Game.ctx = canvas.getContext("2d");
  Game.white = new Image();
  Game.white.onload = f;
  Game.white.src = "/whitePiece.gif";

  Game.black = new Image();
  Game.black.onload = f;
  Game.black.src = "/blackPiece.gif";
});
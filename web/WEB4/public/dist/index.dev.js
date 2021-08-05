"use strict";

$(function () {
  $("#wrapper").click(function (e) {
    $(this).addClass("dark");
    var x = e.pageX + "px";
    var y = e.pageY + "px";
    var img = $('<div class="blip"></div>');
    var div = $('<div class="blip">').css({
      position: "absolute",
      left: x,
      top: y
    });
    div.append(img);
    $("#wrapper").append(div);
    setTimeout(function () {
      $("#wrapper").removeClass("dark");
    }, 1250);
  });
});
//# sourceMappingURL=index.dev.js.map

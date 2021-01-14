$(document).ready(function () {
    var Idx = 0;

    $("#draw-btn").click(function () {
        $.get("/api/v1/DrawCardUser", function (data) {
            $.get("/api/v1/UserDrawResult", function (data) {
                var res = jQuery.parseJSON(data);
                $(".user").append(`<div id="R${++Idx}" value="${res.value}">
                ${res.card}</div>`);
                if (res.status == "burnout") {
                    alert(res.status);
                    $("#draw-btn").remove();
                    $(".stand-btn").remove();
                    $(".input-console").append(`<button class="generate-btn">Play</button>`);
                }
            });
        });
    });

    $(".generate-btn").click(function () {
        $(".generate-btn").remove();
        $(".user").remove();
        $(".column1").append(`<div class="user"></div>`);
        $.get("/api/v1/GenerateDeck", function (data) {
            var res2 = jQuery.parseJSON(data);
            console.log(res2)
            $.each(res2.user_hand, function (i, item) {
                $(".user").append(`<div id="R${++Idx}" value="${item.value}">
            ${item.name}</div>`);
            });
        });
        $(".input-console").append(`<button id="draw-btn">Draw</button>`);
        $(".input-console").append(`<button id="stand-btn">Stand</button>`);
    });
});
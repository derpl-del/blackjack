$(document).ready(function () {
    var Idx = 0;
    var l = 0;
    $('.input-console').on('click', "#generate-btn", function () {
        $("#generate-btn").remove();
        $("#next-btn").remove();
        $("#draw-btn").remove();
        $("#stand-btn").remove();
        $(".user").remove();
        $(".column1").append(`<div class="user"></div>`);
        $(".dealer").remove();
        $(".column2").append(`<div class="dealer"></div>`);
        $.get("/api/v1/GenerateDeck", function (data) {
            var res2 = jQuery.parseJSON(data);
            //console.log(res2)
            $.each(res2.user_hand, function (i, item) {
                $(".user").append(`<div id="R${++Idx}" value="${item.value}">
            ${item.name}</div>`);
            });
            $.each(res2.dealer_hand, function (i, item) {
                $(".dealer").append(`<div id="R${++Idx}" value="${item.value}">
            ${item.name}</div>`);
            });
        });
        $.get("/api/v1/ViewResult", function (data) {
            var res = jQuery.parseJSON(data);
            $(".dealer_score").text(res.dealer_value);
            $(".user_score").text(res.user_value);
            user_value = res.user_value;
            if (user_value == 21) {
                alert("Player Win")
                $(".input-console").append(`<button id="generate-btn">Play</button>`);
            } else if (user_value > 21) {
                alert("Com Win")
                $(".input-console").append(`<button id="generate-btn">Play</button>`);
            }
            else {
                $(".input-console").append(`<button id="draw-btn">Draw</button>`);
                $(".input-console").append(`<button id="stand-btn">Stand</button>`);
            }
        });
    });

    $('.input-console').on('click', "#draw-btn", function () {
        $.get("/api/v1/DrawCardUser", function (data) {
            $.get("/api/v1/UserDrawResult", function (data) {
                var res = jQuery.parseJSON(data);
                $(".user").append(`<div id="R${++Idx}" value="${res.value}">
                ${res.card}</div>`);
                $(".user_score").text(res.score);
                if (res.status == "burnout") {
                    alert(res.status);
                    $("#draw-btn").remove();
                    $("#stand-btn").remove();
                    $(".input-console").append(`<button id="generate-btn">Play</button>`);
                }
            });
        });
    });

    $('.input-console').on('click', "#stand-btn", function () {
        $("#draw-btn").remove();
        $("#stand-btn").remove();
        $(".input-console").append(`<button id="next-btn">Next</button>`);
    });

    $('.input-console').on('click', "#next-btn", function () {
        values_dealer = parseInt($(".dealer_score").text())
        values_player = parseInt($(".user_score").text())
        if (values_dealer > values_player) {
            alert("COM win");
            $("#next-btn").remove();
            $(".input-console").append(`<button id="generate-btn">Play</button>`);
        } else if (values_dealer >= 17) {
            if (values_dealer == values_player) {
                alert("Draw");
                $("#next-btn").remove();
                $(".input-console").append(`<button id="generate-btn">Play</button>`);
            } else if (values_dealer > values_player) {
                alert("COM win");
                $("#next-btn").remove();
                $(".input-console").append(`<button id="generate-btn">Play</button>`);
            } else {
                alert("Player win");
                $("#next-btn").remove();
                $(".input-console").append(`<button id="generate-btn">Play</button>`);
            }
        }
        else {
            $.get("/api/v1/DrawCardDealer", function (data) {
                var res = $.get("/api/v1/DealerDrawResult", function (data) {
                    var res = jQuery.parseJSON(data);
                    $(".dealer").append(`<div id="R${++Idx}" value="${res.value}">
                 ${res.card}</div>`);
                    $(".dealer_score").text(res.score);
                    if (res.status == "burnout") {
                        alert("player win");
                        $("#next-btn").remove();
                        $(".input-console").append(`<button id="generate-btn">Play</button>`);
                    }
                });
            });
        }
    });
});
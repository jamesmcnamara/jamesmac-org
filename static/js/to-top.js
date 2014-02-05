$(document).ready(function() {
    var offset = 250;
    var duration = 500;

    $(window).scroll(function() {
        if ($(this).scrollTop() > offset) {
            $('#back-to-top').fadeIn(duration);
        } 
        else {
            $('#back-to-top').fadeOut(duration);
        }
    });

    $('#back-to-top').click(function(event) {
        event.preventDefault();
        $('html').animate({scrollTop:0}, duration);
    });
});

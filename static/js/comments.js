

	$("a").filter(".comment").click(function() { 
		var divID = $(this).attr("id");
		$(divID + "comments").animate({height:500});
		});

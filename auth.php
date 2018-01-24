<?php

	if ( $_GET["channel"] == "some/channel" && $_SERVER["HTTP_AUTHORIZATION"] != "Bearer 123" ) {
		http_response_code(403);
	} else {
		http_response_code(200);
	}

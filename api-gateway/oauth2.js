function introspectAccessToken(r) {
    r.subrequest("/_oauth2_send_request",
        function (reply) {
            switch (reply.status) {
                case 200:
                    var response = JSON.parse(reply.responseText);

                    if (response.hasOwnProperty("id")) {
                        r.headersOut['X-User-Id'] = response.id;
                    }

                    if (response.hasOwnProperty("role")) {
                        r.headersOut['X-Role'] = response.role;
                    }

                    r.return(200); // Token is valid, return success code
                    break;
                default:
                    r.return(reply.status)
                    break;
            }
        }
    );
}

export default { introspectAccessToken };
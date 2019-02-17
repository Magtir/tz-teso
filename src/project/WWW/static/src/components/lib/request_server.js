function RequestServer(href, method, body, callback) {
    let out = {
        rpl: {},
        error: null
    };
    fetch(href, {
        method: method,
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        },
        body: body
    })
        .then(res => res.json())
        .then(
            (rpl) => {
                out.rpl = rpl;
                callback(out);
            })
        .catch(
            (error) => {
                out.error = error;
                callback(out);
            });
}

export default RequestServer
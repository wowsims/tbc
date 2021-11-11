var workerID = "";

addEventListener('message', async (e) => {
	const msg = e.data.msg;
	const id = e.data.id;

    if (msg == "setID") {
		workerID = id;
		postMessage({ msg: "idconfirm" })
        return;
	}

    var url = "/"+msg;
    let response = await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-protobuf'
        },
        body: e.data.inputData
    });

    const content = await response.arrayBuffer();
    var uint8View = new Uint8Array(content);
    postMessage({
        msg: msg,
        outputData: uint8View,
        id: id,
    });

}, false);

// Let UI know worker is ready.
postMessage({
    msg: "ready"
});
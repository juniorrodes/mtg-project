var dialog;
window.onload = function() {
    dialog = document.querySelector("dialog.advance-search");
}

htmx.defineExtension('submitjson', {
    onEvent: function (name, evt) {
        if (name === "htmx:configRequest") {
            evt.detail.headers['Content-Type'] = "application/json"
            evt.detail.headers['X-API-Key'] = 'sjk_xxx'
        }
    },
    encodeParameters: function(xhr, parameters, elt) {
        xhr.overrideMimeType('text/json') // override default mime type
        const formData = new FormData(event.target);

        var requestForm = {};
        for (const key of formData.keys()) {
            const splitedName = key.split('_');
            const values = formData.getAll(key);
            if (splitedName.length == 1) {
                requestForm[splitedName[0]] = { value: values[0] };
                continue;
            }
            
            var finalValue = values;
            if (splitedName[1] === "operand") {
                finalValue = values[0];
            }

            if (requestForm[splitedName[0]] == undefined) {
                requestForm[splitedName[0]] = { [splitedName[1]]:  finalValue };
                continue;
            }
            requestForm[splitedName[0]][splitedName[1]] = finalValue;
        }
        if (dialog.open) {
            dialog.close();
        }
        const body = { // set your request body
            ...requestForm,
        }
        return (JSON.stringify(body))
    }
})

function handleCardSearch(event) {
    event.preventDefault();
    const formData = new FormData(event.target);

    var requestForm = {};
    for (const key of formData.keys()) {
        const splitedName = key.split('_');
        const values = formData.getAll(key);
        if (splitedName.length == 1) {
            requestForm[splitedName[0]] = { value: values[0] };
            continue;
        }
        
        var finalValue = values;
        if (splitedName[1] === "operand") {
            finalValue = values[0];
        }

        if (requestForm[splitedName[0]] == undefined) {
            requestForm[splitedName[0]] = { [splitedName[1]]:  finalValue };
            continue;
        }
        requestForm[splitedName[0]][splitedName[1]] = finalValue;
    }
    if (dialog.open) {
        dialog.close();
    }

    htmx.ajax('POST', '/api/search', {source: event.target, values: requestForm, target: '#output', swap: 'innerHTML'})
    return false;
}

function openModal() {
    dialog.showModal();
}

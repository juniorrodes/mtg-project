function handleCardSearch(event) {
    event.preventDefault();
    const formData = new FormData(event.target);

    var requestForm = {};
    for (const key of formData.keys()) {
        const splitedName = key.split('_');
        const value = formData.getAll(key);
        if (splitedName.length == 1) {
            requestForm[splitedName[0]] = { value: value[0] };
            continue;
        }
        
        var finalValue;
        if (value.length > 1) {
            finalValue = value;
        } else {
            finalValue = value[0];
        }

        if (requestForm[splitedName[0]] == undefined) {
            requestForm[splitedName[0]] = { [splitedName[1]]:  finalValue };
            continue;
        }
        requestForm[splitedName[0]][splitedName[1]] = finalValue;
    }

    htmx.ajax('POST', '/api/search', {values: requestForm, target: '#output', swap: 'innerHTML'})
    return false;
}


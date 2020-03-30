// It obtain the current date.
exports.GetCurrentDate = function () {
    var date = new Date();
    var year = date.getFullYear();
    var month = date.getMonth() + 1;
    var day = date.getDate();

    var currDate = {};
    currDate.year = year;
    currDate.month = month;
    currDate.day = day;

    return currDate;
}

// It checks the equivalence of two objects. 
exports.IsEquivalent = function (a, b) {
    // Create arrays of property names.
    var aProps = Object.getOwnPropertyNames(a);
    var bProps = Object.getOwnPropertyNames(b);

    // If the number of properties is different, objects are not equivalent.
    if (aProps.length != bProps.length) {
        return false;
    }

    for (var i = 0; i < aProps.length; i++) {
        var propName = aProps[i];

        // If the values of same property are not equal, objects are not equivalent.
        if (a[propName] !== b[propName]) {
            return false;
        }
    }

    return true;
}
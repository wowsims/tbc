// Returns if the two items are equal, or if both are null / undefined.
export function equalsOrBothNull(a, b, comparator) {
    if (a == null && b == null)
        return true;
    if (a == null || b == null)
        return false;
    return (comparator || ((_a, _b) => a == b))(a, b);
}
export function sum(arr) {
    return arr.reduce((total, cur) => total + cur, 0);
}
// Returns the index of maximum value, or null if empty.
export function maxIndex(arr) {
    return arr.reduce((cur, v, i, arr) => v > arr[cur] ? i : cur, 0);
}
// Returns a new array containing only elements present in both a and b.
export function intersection(a, b) {
    return a.filter(value => b.includes(value));
}
// Splits an array into buckets, where elements are placed in the same bucket if the
// toString function returns the same value.
export function bucket(arr, toString) {
    const buckets = {};
    arr.forEach(val => {
        const valString = toString(val);
        if (buckets[valString]) {
            buckets[valString].push(val);
        }
        else {
            buckets[valString] = [val];
        }
    });
    return buckets;
}
export function stDevToConf90(stDev, N) {
    return 1.645 * stDev / Math.sqrt(N);
}
export async function wait(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}
// Only works for numeric enums
export function getEnumValues(enumType) {
    return Object.keys(enumType)
        .filter(key => !isNaN(Number(enumType[key])))
        .map(key => parseInt(enumType[key]));
}
// Whether a click event was a right click.
export function isRightClick(event) {
    return event.button == 2;
}
// Converts from '#ffffff' --> 'rgba(255, 255, 255, alpha)'
export function hexToRgba(hex, alpha) {
    if (/^#([A-Fa-f0-9]{3}){1,2}$/.test(hex)) {
        let parts = hex.substring(1).split('');
        if (parts.length == 3) {
            parts = [parts[0], parts[0], parts[1], parts[1], parts[2], parts[2]];
        }
        const c = '0x' + parts.join('');
        return 'rgba(' + [(c >> 16) & 255, (c >> 8) & 255, c & 255].join(',') + ',' + alpha + ')';
    }
    throw new Error('Invalid hex color: ' + hex);
}
export function camelToSnakeCase(str) {
    let result = str.replace(/[A-Z]/g, letter => `_${letter.toLowerCase()}`);
    if (result.startsWith('_')) {
        result = result.substring(1);
    }
    return result;
}
export function downloadJson(json, fileName) {
    const dataStr = "data:text/json;charset=utf-8," + encodeURIComponent(JSON.stringify(json));
    const downloadAnchorNode = document.createElement('a');
    downloadAnchorNode.setAttribute("href", dataStr);
    downloadAnchorNode.setAttribute("download", fileName);
    document.body.appendChild(downloadAnchorNode); // required for firefox
    downloadAnchorNode.click();
    downloadAnchorNode.remove();
}

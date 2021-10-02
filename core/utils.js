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
// Returns a new array containing only elements present in both a and b.
export function intersection(a, b) {
    return a.filter(value => b.includes(value));
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

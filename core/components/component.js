export class Component {
    constructor(parentElem, rootCssClass, rootElem) {
        this.rootElem = rootElem || document.createElement('div');
        this.rootElem.classList.add(rootCssClass);
        if (parentElem) {
            parentElem.appendChild(this.rootElem);
        }
    }
}

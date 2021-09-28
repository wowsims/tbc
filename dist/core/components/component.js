export class Component {
    constructor(parentElem, rootCssClass, rootElem) {
        this.rootElem = rootElem || document.createElement('div');
        this.rootElem.classList.add(rootCssClass);
        parentElem.appendChild(this.rootElem);
    }
}

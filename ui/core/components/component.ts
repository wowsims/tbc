export abstract class Component {
  readonly rootElem: HTMLElement;

  constructor(parentElem: HTMLElement, rootCssClass: string, rootElem?: HTMLElement) {
    this.rootElem = rootElem || document.createElement('div');
    this.rootElem.classList.add(rootCssClass);
    parentElem.appendChild(this.rootElem);
  }
}

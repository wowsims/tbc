export abstract class Component {
  readonly rootElem: HTMLDivElement;

  constructor(parentElem: HTMLElement, rootCssClass: string) {
    this.rootElem = document.createElement('div');
    this.rootElem.classList.add(rootCssClass);
    parentElem.appendChild(this.rootElem);
  }
}

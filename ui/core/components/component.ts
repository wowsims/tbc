export abstract class Component {
  abstract getRootElement(): HTMLElement;

  appendTo(newParent: Element) {
    newParent.appendChild(this.getRootElement());
  }
}

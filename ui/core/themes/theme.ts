export abstract class Theme {
  readonly parentElem: HTMLElement;

  constructor(parentElem: HTMLElement) {
    this.parentElem = parentElem;
  }
}

import { Component } from './component.js';
export class CloseButton extends Component {
    constructor(parent, onClick) {
        super(parent, 'close-button');
        this.rootElem.innerHTML = `
			<span aria-hidden="true">&times;</span>
		`;
        this.rootElem.addEventListener('click', event => onClick());
    }
}

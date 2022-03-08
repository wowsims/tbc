import { IndividualSimSettings } from '/tbc/core/proto/ui.js';
import { downloadString } from '/tbc/core/utils.js';
import { Popup } from './popup.js';
export function newIndividualExporters(simUI) {
    const exportSettings = document.createElement('div');
    exportSettings.classList.add('export-settings', 'sim-dropdown-menu');
    exportSettings.innerHTML = `
		<span id="exportMenuLink" class="dropdown-toggle fas fa-file-export" role="button" data-toggle="dropdown" aria-haspopup="true" arai-expanded="false"></span>
		<div class="dropdown-menu dropdown-menu-right" aria-labelledby="exportMenuLink">
		</div>
	`;
    const linkElem = exportSettings.getElementsByClassName('dropdown-toggle')[0];
    tippy(linkElem, {
        'content': 'Export',
        'allowHTML': true,
    });
    const menuElem = exportSettings.getElementsByClassName('dropdown-menu')[0];
    const addMenuItem = (label, onClick) => {
        const itemElem = document.createElement('span');
        itemElem.classList.add('dropdown-item');
        itemElem.textContent = label;
        itemElem.addEventListener('click', onClick);
        menuElem.appendChild(itemElem);
    };
    addMenuItem('Link', () => {
        new IndividualLinkExporter(menuElem, simUI);
    });
    addMenuItem('Json', () => {
        new IndividualJsonExporter(menuElem, simUI);
    });
    return exportSettings;
}
class Exporter extends Popup {
    constructor(parent, title, allowDownload) {
        super(parent);
        this.rootElem.classList.add('exporter');
        this.rootElem.innerHTML = `
			<span class="exporter-title">${title}</span>
			<div class="export-content">
				<textarea class="exporter-textarea" readonly></textarea>
			</div>
			<div class="actions-row">
				<button class="exporter-button sim-button clipboard-button">COPY TO CLIPBOARD</button>
				<button class="exporter-button sim-button download-button">DOWNLOAD</button>
			</div>
		`;
        this.addCloseButton();
        this.textElem = this.rootElem.getElementsByClassName('exporter-textarea')[0];
        const clipboardButton = this.rootElem.getElementsByClassName('clipboard-button')[0];
        clipboardButton.addEventListener('click', event => {
            const data = this.textElem.textContent;
            if (navigator.clipboard == undefined) {
                alert(data);
            }
            else {
                navigator.clipboard.writeText(data);
            }
        });
        const downloadButton = this.rootElem.getElementsByClassName('download-button')[0];
        if (allowDownload) {
            downloadButton.addEventListener('click', event => {
                const data = this.textElem.textContent;
                downloadString(data, 'wowsims.json');
            });
        }
        else {
            downloadButton.remove();
        }
    }
    init() {
        this.textElem.textContent = this.getData();
    }
}
class IndividualLinkExporter extends Exporter {
    constructor(parent, simUI) {
        super(parent, 'Sharable Link', false);
        this.simUI = simUI;
        this.init();
    }
    getData() {
        const proto = this.simUI.toProto();
        // When sharing links, people generally don't intend to share settings/ep weights.
        proto.settings = undefined;
        proto.epWeights = [];
        const protoBytes = IndividualSimSettings.toBinary(proto);
        const deflated = pako.deflate(protoBytes, { to: 'string' });
        const encoded = btoa(String.fromCharCode(...deflated));
        const linkUrl = new URL(window.location.href);
        linkUrl.hash = encoded;
        return linkUrl.toString();
    }
}
class IndividualJsonExporter extends Exporter {
    constructor(parent, simUI) {
        super(parent, 'JSON Export', true);
        this.simUI = simUI;
        this.init();
    }
    getData() {
        return JSON.stringify(IndividualSimSettings.toJson(this.simUI.toProto()), null, 2);
    }
}

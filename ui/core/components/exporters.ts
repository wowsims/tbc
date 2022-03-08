import { Class } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { IndividualSimSettings } from '/tbc/core/proto/ui.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { Player } from '/tbc/core/player.js';
import { classNames, nameToClass, nameToRace } from '/tbc/core/proto_utils/names.js';
import { talentSpellIdsToTalentString } from '/tbc/core/talents/factory.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { downloadString } from '/tbc/core/utils.js';

import { Popup } from './popup.js';

declare var $: any;
declare var tippy: any;
declare var pako: any;

export function newIndividualExporters<SpecType extends Spec>(simUI: IndividualSimUI<SpecType>): HTMLElement {
	const exportSettings = document.createElement('div');
	exportSettings.classList.add('export-settings', 'sim-dropdown-menu');
	exportSettings.innerHTML = `
		<span id="exportMenuLink" class="dropdown-toggle fas fa-file-export" role="button" data-toggle="dropdown" aria-haspopup="true" arai-expanded="false"></span>
		<div class="dropdown-menu dropdown-menu-right" aria-labelledby="exportMenuLink">
		</div>
	`;
	const linkElem = exportSettings.getElementsByClassName('dropdown-toggle')[0] as HTMLElement;
	tippy(linkElem, {
		'content': 'Export',
		'allowHTML': true,
	});

	const menuElem = exportSettings.getElementsByClassName('dropdown-menu')[0] as HTMLElement;
	const addMenuItem = (label: string, onClick: () => void) => {
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

abstract class Exporter extends Popup {
  private readonly textElem: HTMLElement;

  constructor(parent: HTMLElement, title: string, allowDownload: boolean) {
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

    this.textElem = this.rootElem.getElementsByClassName('exporter-textarea')[0] as HTMLElement;

    const clipboardButton = this.rootElem.getElementsByClassName('clipboard-button')[0] as HTMLElement;
		clipboardButton.addEventListener('click', event => {
			const data = this.textElem.textContent!;
			if (navigator.clipboard == undefined) {
				alert(data);
			} else {
				navigator.clipboard.writeText(data);
			}
		});

    const downloadButton = this.rootElem.getElementsByClassName('download-button')[0] as HTMLElement;
		if (allowDownload) {
			downloadButton.addEventListener('click', event => {
				const data = this.textElem.textContent!;
				downloadString(data, 'wowsims.json');
			});
		} else {
			downloadButton.remove();
		}
	}

	protected init() {
		this.textElem.textContent = this.getData();
	}

	abstract getData(): string;
}

class IndividualLinkExporter<SpecType extends Spec> extends Exporter {
	private readonly simUI: IndividualSimUI<SpecType>;

  constructor(parent: HTMLElement, simUI: IndividualSimUI<SpecType>) {
    super(parent, 'Sharable Link', false);
		this.simUI = simUI;
		this.init();
	}

	getData(): string {
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

class IndividualJsonExporter<SpecType extends Spec> extends Exporter {
	private readonly simUI: IndividualSimUI<SpecType>;

  constructor(parent: HTMLElement, simUI: IndividualSimUI<SpecType>) {
    super(parent, 'JSON Export', true);
		this.simUI = simUI;
		this.init();
	}

	getData(): string {
		return JSON.stringify(IndividualSimSettings.toJson(this.simUI.toProto()), null, 2);
	}
}

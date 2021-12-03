import { Component } from '/tbc/core/components/component.js';
import { EnumPicker } from '/tbc/core/components/enum_picker.js';
import { makePhaseSelector } from '/tbc/core/components/other_inputs.js';
import { Raid } from '/tbc/core/raid.js';
import { MAX_PARTY_SIZE } from '/tbc/core/party.js';
import { Party } from '/tbc/core/party.js';
import { Player } from '/tbc/core/player.js';
import { Class } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Faction } from '/tbc/core/proto_utils/utils.js';
import { SpecOptions } from '/tbc/core/proto_utils/utils.js';
import { SpecRotation } from '/tbc/core/proto_utils/utils.js';
import { classColors } from '/tbc/core/proto_utils/utils.js';
import { specToClass } from '/tbc/core/proto_utils/utils.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { getEnumValues } from '/tbc/core/utils.js';
import { hexToRgba } from '/tbc/core/utils.js';

declare var tippy: any;

export class RaidPicker extends Component {
	private readonly raid: Raid;
  private readonly partyPickers: Array<PartyPicker>;

  constructor(parent: HTMLElement, raid: Raid, presets: Array<PresetSpecSettings<any>>) {
    super(parent, 'raid-picker-root');
		this.raid = raid;

    const raidViewer = document.createElement('div');
    raidViewer.classList.add('current-raid-viewer');
    this.rootElem.appendChild(raidViewer);
		raidViewer.innerHTML = `
			<div class="parties-container">
			</div>
		`;

    const partiesContainer = this.rootElem.getElementsByClassName('parties-container')[0] as HTMLDivElement;
		this.partyPickers = this.raid.getParties().map((party, i) => new PartyPicker(partiesContainer, party, i));

    const newPlayerPickerRoot = document.createElement('div');
    newPlayerPickerRoot.classList.add('new-player-picker');
    this.rootElem.appendChild(newPlayerPickerRoot);

		const newPlayerPicker = new NewPlayerPicker(newPlayerPickerRoot, this.raid, presets);
	}
}

export class PartyPicker extends Component {
	private readonly party: Party;
	private readonly index: number;
  private readonly playerPickers: Array<PlayerPicker>;

  constructor(parent: HTMLElement, party: Party, index: number) {
    super(parent, 'party-picker-root');
		this.party = party;
		this.index = index;

		this.rootElem.innerHTML = `
			<div class="party-header">
				<span>Group ${index + 1}</span>
			</div>
			<div class="players-container">
			</div>
		`;

    const playersContainer = this.rootElem.getElementsByClassName('players-container')[0] as HTMLDivElement;
		this.playerPickers = [...Array(MAX_PARTY_SIZE).keys()].map(i => new PlayerPicker(playersContainer, this.party, i));
	}
}

export class PlayerPicker extends Component {
	private readonly party: Party;
	private readonly playerIndex: number;
	private player: Player<any> | null;

	private readonly nameElem: HTMLSpanElement;

  constructor(parent: HTMLElement, party: Party, playerIndex: number) {
    super(parent, 'player-picker-root');
		this.party = party;
		this.playerIndex = playerIndex;
		this.player = null;

		this.rootElem.innerHTML = `
			<div class="player-label">
				<span class="player-name"></span>
			</div>
			<div class="player-options">
			</div>
		`;

		this.nameElem = this.rootElem.getElementsByClassName('player-name')[0] as HTMLSpanElement;

		this.party.changeEmitter.on(() => {
			const newPlayer = this.party.getPlayer(this.playerIndex);

			if (((newPlayer == null) != (this.player == null)) || newPlayer != this.player) {
				this.setPlayer(newPlayer);
				return;
			}

			this.update();
		});

		this.update();
	}

	setPlayer(newPlayer: Player<any> | null) {
		this.player = newPlayer;
		this.party.setPlayer(this.playerIndex, this.player);

		this.update();
	}

	private update() {
		if (this.player == null) {
			this.rootElem.classList.add('empty');
			this.nameElem.textContent = '';
		} else {
			this.rootElem.classList.remove('empty');
			this.nameElem.textContent = this.player.getName();
		}
	}
}

export class NewPlayerPicker extends Component {
	private readonly raid: Raid;
	private currentFaction: Faction;

  constructor(parent: HTMLElement, raid: Raid, presets: Array<PresetSpecSettings<any>>) {
    super(parent, 'new-player-picker-root');
		this.raid = raid;
		this.currentFaction = Faction.Alliance;

		this.rootElem.innerHTML = `
			<div class="new-player-picker-controls">
				<div class="faction-selector"></div>
				<div class="phase-selector"></div>
			</div>
			<div class="presets-container"></div>
		`;

		const factionSelector = new EnumPicker<NewPlayerPicker>(this.rootElem.getElementsByClassName('faction-selector')[0] as HTMLElement, this, {
			label: 'Faction',
			labelTooltip: 'Default faction for newly-created players.',
			values: [
				{ name: 'Alliance', value: Faction.Alliance },
				{ name: 'Horde', value: Faction.Horde },
			],
			changedEvent: (picker: NewPlayerPicker) => new TypedEvent<void>(),
			getValue: (picker: NewPlayerPicker) => picker.currentFaction,
			setValue: (picker: NewPlayerPicker, newValue: number) => {
				picker.currentFaction = newValue;
			},
		});

		const phaseSelector = makePhaseSelector(this.rootElem.getElementsByClassName('phase-selector')[0] as HTMLElement, this.raid.sim);

		const presetsContainer = this.rootElem.getElementsByClassName('presets-container')[0] as HTMLElement;
		getEnumValues(Class).forEach(wowClass => {
			const matchingPresets = presets.filter(preset => specToClass[preset.spec] == wowClass);
			if (matchingPresets.length == 0 || wowClass == Class.ClassUnknown) {
				return;
			}

			const classPresetsContainer = document.createElement('div');
			classPresetsContainer.classList.add('class-presets-container');
			presetsContainer.appendChild(classPresetsContainer);
			classPresetsContainer.style.backgroundColor = hexToRgba(classColors[wowClass as Class], 0.5);

			matchingPresets.forEach(matchingPreset => {
				const presetIndex = presets.findIndex(matchingPreset);

				const presetElem = document.createElement('div');
				presetElem.classList.add('preset-picker');
				classPresetsContainer.appendChild(presetElem);

        const presetIconElem = document.createElement('img');
        presetIconElem.classList.add('preset-picker-icon');
				presetElem.appendChild(presetIconElem);
				presetIconElem.src = matchingPreset.iconUrl;
				tippy(presetIconElem, {
					'content': matchingPreset.tooltip,
					'allowHTML': true,
				});

				presetElem.setAttribute('draggable', true);
				presetElem.addEventListener("dragstart", event => {
					event.dataTransfer.setData('text/plain', presetIndex);
					event.dataTransfer.setDragImage(matchingPreset.iconUrl, 30, 30);
					event.dataTransfer.dropEffect = 'copy';
				});
			});
		});
	}
}

export interface PresetSpecSettings<SpecType extends Spec> {
	spec: Spec,
	rotation: SpecRotation<SpecType>,
	talents: string,
	specOptions: SpecOptions<SpecType>,
	consumes: Consumes,

	defaultFactionRaces: Record<Faction, Race>,
	defaultGear: Record<Faction, Record<number, EquipmentSpec>>,

	tooltip: string,
	iconUrl: string,
}

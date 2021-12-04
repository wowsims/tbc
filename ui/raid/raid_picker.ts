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

const NEW_PLAYER: number = -1;

export class RaidPicker extends Component {
	readonly raid: Raid;
	readonly presets: Array<PresetSpecSettings<any>>;
  readonly partyPickers: Array<PartyPicker>;
	readonly newPlayerPicker: NewPlayerPicker;

	// Hold data about the player being dragged while the drag is happening.
	currentDragPlayer: Player<any> | null = null;
	currentDragPlayerFromIndex: number = NEW_PLAYER;

  constructor(parent: HTMLElement, raid: Raid, presets: Array<PresetSpecSettings<any>>) {
    super(parent, 'raid-picker-root');
		this.raid = raid;
		this.presets = presets;

    const raidViewer = document.createElement('div');
    raidViewer.classList.add('current-raid-viewer');
    this.rootElem.appendChild(raidViewer);
		raidViewer.innerHTML = `
			<div class="parties-container">
			</div>
		`;

    const partiesContainer = this.rootElem.getElementsByClassName('parties-container')[0] as HTMLDivElement;
		this.partyPickers = this.raid.getParties().map((party, i) => new PartyPicker(partiesContainer, party, i, this));

    const newPlayerPickerRoot = document.createElement('div');
    newPlayerPickerRoot.classList.add('new-player-picker');
    this.rootElem.appendChild(newPlayerPickerRoot);

		this.newPlayerPicker = new NewPlayerPicker(newPlayerPickerRoot, this);

		this.rootElem.ondragend = event => {
			if (this.currentDragPlayerFromIndex != NEW_PLAYER) {
				const playerPicker = this.getPlayerPicker(this.currentDragPlayerFromIndex);
				playerPicker.setPlayer(null);
			}

			this.clearDragPlayer();
		};
	}

	getCurrentFaction(): Faction {
		return this.newPlayerPicker.currentFaction;
	}

	getCurrentPhase(): number {
		return this.raid.sim.getPhase();
	}

	getPlayerPicker(raidIndex: number): PlayerPicker {
		return this.partyPickers[Math.floor(raidIndex / MAX_PARTY_SIZE)].playerPickers[raidIndex % MAX_PARTY_SIZE];
	}

	setDragPlayer(player: Player<any>, fromIndex: number) {
		this.clearDragPlayer();

		this.currentDragPlayer = player;
		this.currentDragPlayerFromIndex = fromIndex;

		if (fromIndex != NEW_PLAYER) {
			const playerPicker = this.getPlayerPicker(fromIndex);
			playerPicker.rootElem.classList.add('dragFrom');
		}
	}

	clearDragPlayer() {
		if (this.currentDragPlayerFromIndex != NEW_PLAYER) {
			const playerPicker = this.getPlayerPicker(this.currentDragPlayerFromIndex);
			playerPicker.rootElem.classList.remove('dragFrom');
		}

		this.currentDragPlayer = null;
		this.currentDragPlayerFromIndex = NEW_PLAYER;
	}
}

export class PartyPicker extends Component {
	readonly party: Party;
	readonly index: number;
	readonly raidPicker: RaidPicker;
  readonly playerPickers: Array<PlayerPicker>;

  constructor(parent: HTMLElement, party: Party, index: number, raidPicker: RaidPicker) {
    super(parent, 'party-picker-root');
		this.party = party;
		this.index = index;
		this.raidPicker = raidPicker;

		this.rootElem.innerHTML = `
			<div class="party-header">
				<span>Group ${index + 1}</span>
			</div>
			<div class="players-container">
			</div>
		`;

    const playersContainer = this.rootElem.getElementsByClassName('players-container')[0] as HTMLDivElement;
		this.playerPickers = [...Array(MAX_PARTY_SIZE).keys()].map(i => new PlayerPicker(playersContainer, this, i));
	}
}

export class PlayerPicker extends Component {
	// Index of this player within its party (0-4).
	readonly index: number;

	// Index of this player within the whole raid (0-24).
	readonly raidIndex: number;

	player: Player<any> | null;

	readonly partyPicker: PartyPicker;
	readonly raidPicker: RaidPicker;

	private readonly labelElem: HTMLElement;
	private readonly iconElem: HTMLImageElement;
	private readonly nameElem: HTMLSpanElement;

  constructor(parent: HTMLElement, partyPicker: PartyPicker, index: number) {
    super(parent, 'player-picker-root');
		this.index = index;
		this.raidIndex = partyPicker.index * MAX_PARTY_SIZE + index;
		this.player = null;
		this.partyPicker = partyPicker;
		this.raidPicker = partyPicker.raidPicker;

		this.rootElem.innerHTML = `
			<div class="player-label">
				<img class="player-icon"></img>
				<span class="player-name"></span>
			</div>
			<div class="player-options">
			</div>
		`;

		this.labelElem = this.rootElem.getElementsByClassName('player-label')[0] as HTMLElement;
		this.iconElem = this.rootElem.getElementsByClassName('player-icon')[0] as HTMLImageElement;
		this.nameElem = this.rootElem.getElementsByClassName('player-name')[0] as HTMLSpanElement;

		this.partyPicker.party.changeEmitter.on(() => {
			const newPlayer = this.partyPicker.party.getPlayer(this.index);

			if (((newPlayer == null) != (this.player == null)) || newPlayer != this.player) {
				this.setPlayer(newPlayer);
				return;
			}

			this.update();
		});

		this.labelElem.ondragstart = event => {
			if (this.player == null) {
				event.preventDefault();
				return;
			}

			const iconSrc = this.iconElem.src;
			console.log('icon: ' + iconSrc);
			const dragImage = new Image();
			dragImage.src = iconSrc;
			event.dataTransfer!.setDragImage(dragImage, 30, 30);
			event.dataTransfer!.setData("text/plain", iconSrc);

			event.dataTransfer!.dropEffect = 'move';

			this.raidPicker.setDragPlayer(this.player, this.raidIndex);
		};

		let dragEnterCounter = 0;
		this.rootElem.ondragenter = event => {
			event.preventDefault();
			dragEnterCounter++;
			this.rootElem.classList.add('dragto');
		};
		this.rootElem.ondragleave = event => {
			event.preventDefault();
			dragEnterCounter--;
			if (dragEnterCounter <= 0) {
				this.rootElem.classList.remove('dragto');
			}
		};
		this.rootElem.ondragover = event => {
			event.preventDefault();
		};
		this.rootElem.ondrop = event => {
			event.preventDefault();
			dragEnterCounter = 0;
			this.rootElem.classList.remove('dragto');

			if (this.raidPicker.currentDragPlayer == null) {
				return;
			}

			if (this.raidPicker.currentDragPlayerFromIndex == this.raidIndex) {
				this.raidPicker.clearDragPlayer();
				return;
			}

			this.setPlayer(this.raidPicker.currentDragPlayer);
			this.iconElem.src = event.dataTransfer!.getData('text/plain');

			if (this.raidPicker.currentDragPlayerFromIndex != NEW_PLAYER) {
				const fromPlayerPicker = this.raidPicker.getPlayerPicker(this.raidPicker.currentDragPlayerFromIndex);
				fromPlayerPicker.setPlayer(null);
			}

			this.raidPicker.clearDragPlayer();
		};

		this.update();
	}

	setPlayer(newPlayer: Player<any> | null) {
		this.player = newPlayer;
		this.partyPicker.party.setPlayer(this.index, this.player);

		this.update();
	}

	private update() {
		if (this.player == null) {
			this.rootElem.classList.add('empty');
			this.rootElem.style.backgroundColor = 'black';
			this.labelElem.setAttribute('draggable', 'false');
			this.nameElem.textContent = '';
		} else {
			this.rootElem.classList.remove('empty');
			this.rootElem.style.backgroundColor = classColors[specToClass[this.player.spec]];
			this.labelElem.setAttribute('draggable', 'true');
			this.nameElem.textContent = this.player.getName();
		}
	}
}

export class NewPlayerPicker extends Component {
	readonly raidPicker: RaidPicker;
	currentFaction: Faction;

  constructor(parent: HTMLElement, raidPicker: RaidPicker) {
    super(parent, 'new-player-picker-root');
		this.raidPicker = raidPicker;
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

		const phaseSelector = makePhaseSelector(this.rootElem.getElementsByClassName('phase-selector')[0] as HTMLElement, this.raidPicker.raid.sim);

		const presetsContainer = this.rootElem.getElementsByClassName('presets-container')[0] as HTMLElement;
		getEnumValues(Class).forEach(wowClass => {
			const matchingPresets = this.raidPicker.presets.filter(preset => specToClass[preset.spec] == wowClass);
			if (matchingPresets.length == 0 || wowClass == Class.ClassUnknown) {
				return;
			}

			const classPresetsContainer = document.createElement('div');
			classPresetsContainer.classList.add('class-presets-container');
			presetsContainer.appendChild(classPresetsContainer);
			classPresetsContainer.style.backgroundColor = hexToRgba(classColors[wowClass as Class], 0.5);

			matchingPresets.forEach(matchingPreset => {
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

				presetElem.setAttribute('draggable', 'true');
				presetElem.ondragstart = event => {
					const dragImage = new Image();
					dragImage.src = matchingPreset.iconUrl;
					event.dataTransfer!.setDragImage(dragImage, 30, 30);
					event.dataTransfer!.setData("text/plain", matchingPreset.iconUrl);

					event.dataTransfer!.dropEffect = 'copy';

					const newPlayer = new Player(matchingPreset.spec, this.raidPicker.raid.sim);
					newPlayer.setRace(matchingPreset.defaultFactionRaces[this.raidPicker.getCurrentFaction()]);
					newPlayer.setRotation(matchingPreset.rotation);
					newPlayer.setTalentsString(matchingPreset.talents);
					newPlayer.setSpecOptions(matchingPreset.specOptions);
					newPlayer.setGear(
							this.raidPicker.raid.sim.lookupEquipmentSpec(
									matchingPreset.defaultGear[this.raidPicker.getCurrentFaction()][this.raidPicker.getCurrentPhase()]));
					newPlayer.setConsumes(matchingPreset.consumes);
					newPlayer.setName(matchingPreset.defaultName);

					this.raidPicker.setDragPlayer(newPlayer, NEW_PLAYER);
				};
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

	defaultName: string,
	defaultFactionRaces: Record<Faction, Race>,
	defaultGear: Record<Faction, Record<number, EquipmentSpec>>,

	tooltip: string,
	iconUrl: string,
}
